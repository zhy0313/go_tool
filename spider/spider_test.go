package spider

import (
	"testing"
)

func TestSpider(t *testing.T) {
	spiders := Spider{}
	spiders.Method = "get"
	spiders.Wait = 2
	spiders.Url = "http://www.lenggirl.com"

	spiders.SetLogLevel("DEBUg")

	spiders.NewHeader("", "www.baidu.com", "")
	body, err := spiders.Go()
	if err != nil {
		spiders.Log().Error(err)
	} else {
		spiders.Log().Infof("%s", string(body))
	}
	err = TooSortSizes(body, 500)
	spiders.Log().Error(err.Error())
}
