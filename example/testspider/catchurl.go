package testspider

import (
	. "github.com/hunterhug/go_tool/spider"
	"log"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strings"
	. "github.com/hunterhug/go_tool/util"
)

var (
	urlpath = "http://www.hrssgz.gov.cn/vsgzpiapp01/GZPI/Gateway/QueryGraduateApply.aspx"
	header = Requestheader
)

//第一次抓取不需要post数据，只需Get
func FirstCatUrl() ([]byte, error) {
	//头部
	content, err := Get(urlpath, header)
	if err != nil {
		log.Println(err.Error())
	}
	return content, err
}

//循环抓取
func LoopCatUrl(filedir string, year int) {
	//第一次抓取
	content, err := FirstCatUrl()

	if err != nil {
		panic("抓取失败")
	}
	//保存到文件中
	err = SaveTofile(filedir + "/1.html", content)
	if err != nil {
		log.Print("%s/n", err.Error())
	}

	//接着，开始循环
	//第二次开始需要post数据，这种反爬虫手段特别牛逼
	// __EVENTTARGET
	//2 1
	//3 2
	//4 3
	//5 4
	//6 5
	//7 6
	//8 7
	//9 8
	//10 9
	//11 10     ------------ -1

	//-----------
	//12 2    ----------- 取个位数
	//13 3
	//14 4
	//19 9
	//20 10
	//21 1
	//22 2
	//23 3
	//30 10
	//规律出来了
	Hiddendata := url.Values{}
	for i := 2; i < 20; i++ {
		Hiddendata = FindSearchHidden(content)
		Hiddendata.Set("DDApplyYear", IS(year))
		Hiddendata.Set("TxtCollege", "")
		Hiddendata.Set("TxtCorpName", "")
		Hiddendata.Set("TxtIDCard", "")
		Hiddendata.Set("TxtName", "")
		if i <= 11 {
			Hiddendata.Set("__EVENTTARGET", "DDDeclareInfo$_ctl14$_ctl" + IS(i - 1))
			log.Println("DDDeclareInfo$_ctl14$_ctl" + IS(i - 1))
		}else {
			temp := IS(i)
			lastone := temp[len(temp) - 1]
			log.Println("DDDeclareInfo$_ctl14$_ctl" + string(lastone))
			Hiddendata.Set("__EVENTTARGET", "DDDeclareInfo$_ctl14$_ctl" + string(lastone))
		}
		content, err = Post(urlpath, Hiddendata, header,true)
		if err != nil {
			panic(err.Error())
		}
		log.Printf("已经抓到第%d页\n", i)
		filetemp := filedir + "/" + IS(i) + ".html"

		sizes := 5.0
		TooSortSizes(content, sizes)
		err := SaveTofile(filetemp, content)
		if err != nil {
			log.Println(err.Error())
		}else {
			log.Printf("已经创建文件%s\n", filetemp)
		}
	}
}

// 查找隐藏字段，构造post
func FindSearchHidden(data []byte) url.Values {
	searchdata := url.Values{}
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(data)))
	doc.Find("input[type=hidden]").Each(func(i int, node *goquery.Selection) {
		ss, _ := node.Attr("name")
		vv, _ := node.Attr("value")
		if ss != "" && vv != "" {
			searchdata.Set(ss, vv)
		}
	})
	return searchdata
}