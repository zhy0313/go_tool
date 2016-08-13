/*
 * Created by 一只尼玛 on 2016/8/12.
 * 功能： 网站爬取功能
 *
 */
package spider

import (
	"log"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//可以允许添加自定义头部
func Get(url string, resheader map[string][]string) (body []byte, e error) {
	WaitM()

	log.Println("GET链接:" + url)

	//新建请求
	request, _ := http.NewRequest("GET", url, nil)

	//带头部，并发不影响，所以克隆
	request.Header = CloneHeader(resheader)

	//OutputMaps("---------request携带头部--------", request.Header)

	//开始请求
	response, err := Client.Do(request)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer response.Body.Close()

	//OutputMaps("----------response携带头部-----------", response.Header)
	log.Printf("\n状态：%v:%v\n", response.Status, response.Proto)

	//设置新Cookie
	Cookieb = MergeCookie(Cookieb, response.Cookies())

	//返回内容
	body, e = ioutil.ReadAll(response.Body)

	return
}

// Post附带信息
func Post(url string, postValues url.Values, header map[string][]string,printpostdata bool) (body []byte, e error) {
	WaitM()
	log.Println("POST链接:" + url)
	if printpostdata {
		OutputMaps("POST 参数:", postValues)
	}
	var request = &http.Request{}
	if postValues != nil {
		pr := ioutil.NopCloser(strings.NewReader(postValues.Encode()))
		request, _ = http.NewRequest("POST", url, pr)
	} else {
		request, _ = http.NewRequest("POST", url, nil)
	}
	request.Header = CloneHeader(header)

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	//OutputMaps("-----------request携带头部-----------", request.Header)

	response, err := Client.Do(request)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	defer response.Body.Close()

	body, e = ioutil.ReadAll(response.Body)

	//OutputMaps("---------response携带头部-------------", response.Header)
	log.Printf("\n状态：%v:%v\n", response.Status, response.Proto)

	//设置新Cookie
	MergeCookie(Cookieb, response.Cookies())

	return
}
