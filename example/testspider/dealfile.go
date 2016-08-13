/*
 * Created by 一只尼玛 on 2016/8/12.
 * 功能： 提取网页信息
 *
 */
package testspider

import (
	//"log"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func DealFile(data []byte) [][]interface{}{
	returndatas:=[][]interface{}{}
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(data)))
	doc.Find("tr[class=ListItem]").Each(func(i int, node *goquery.Selection) {
		returndata:=[]interface{}{}
		node.Find("td").Each(func(i int, node1 *goquery.Selection) {
			//log.Printf("%v\t", node1.Text())
			returndata=append(returndata,node1.Text())
		})
		if len(returndata)==6 {
			returndatas = append(returndatas, returndata)
		}
		//log.Println()

	})
	doc.Find("tr[class=ListAltern]").Each(func(i int, node *goquery.Selection) {
		returndata:=[]interface{}{}
		node.Find("td").Each(func(i int, node1 *goquery.Selection) {
			//log.Printf("%v\t", node1.Text())
			returndata=append(returndata,node1.Text())
		})
		if len(returndata)==6 {
			returndatas = append(returndatas, returndata)
		}
		//log.Println()

	})
	return returndatas
}