//常量包
package spider

const (
	//暂停时间
	DeadTime = 5
)

var (
	//浏览器头部
	Ua = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:46.0) Gecko/20100101 Firefox/46.0"
	Host = "www.hrssgz.gov.cn"
	Referer = "http://www.hrssgz.gov.cn/gzbys/sjcx/"
	Requestheader = map[string][]string{
		"User-Agent": {
			Ua,
		},
		"Host": {
			Host,
		},
		"Accept": {
			"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
		},
		"Connection": {
			"keep-alive",
		},
		"Referer":{
			Referer,
		},
	}
)
