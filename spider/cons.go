//常量包
package spider

const (
	//暂停时间
	WaitTime = 5
)

var (
	//浏览器头部
	FoxfireLinux = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:46.0) Gecko/20100101 Firefox/46.0"
	SpiderHeader = map[string][]string{
		"User-Agent": {
			FoxfireLinux,
		},

	}
)

func NewHeader(ua string, host string, refer string) map[string][]string {
	if ua == "" {
		ua = FoxfireLinux
	}
	h := map[string][]string{
		"User-Agent": {
			ua,
		},
		"Host": {
			host,
		},
		"Referer":{
			refer,
		},
	}
	return h
}