package spider

import (
	"testing"
	"fmt"
)

func TestSpider(t *testing.T) {
	url := "http://www.baidu.com"
	body, err := Get(url, nil)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(body))
	}
}
