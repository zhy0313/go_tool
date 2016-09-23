/*
 * Created by 一只尼玛 on 2016/8/12.
 * 功能： 网站爬取功能
 *
 */
package spider

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"github.com/op/go-logging"
	"os"
	"time"
)

var Log = logging.MustGetLogger("go_tool_spider")
var format = logging.MustStringFormatter(
	"%{color}%{time:2006-01-02 15:04:05.000} %{longpkg}:%{longfunc} [%{level:.5s}]:%{color:reset} %{message}",
)

func init() {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)
	logging.SetLevel(logging.INFO, "go_tool_spider")
}

type  Spider struct {
	Url    string
	Method string //Get Post
	Header http.Header
	Data   url.Values
	Wait   time.Duration
}

func (this *Spider) SetLogLevel(level string) {
	lvl, _ := logging.LogLevel(level)
	logging.SetLevel(lvl, "go_tool_spider")
}

func (this *Spider) Go() (body []byte, e error) {
	if strings.ToLower(this.Method) == "post" {
		return this.Post()
	} else {
		return this.Get()
	}

}
//可以允许添加自定义头部
func (this *Spider)  Get() (body []byte, e error) {
	Wait(this.Wait)

	Log.Debug("GET url:" + this.Url)

	//新建请求
	request, _ := http.NewRequest("GET", this.Url, nil)

	//带头部，并发不影响，所以克隆
	request.Header = CloneHeader(this.Header)

	OutputMaps("---------request header--------", request.Header)

	//开始请求
	response, err := Client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	OutputMaps("----------response header-----------", response.Header)
	Log.Debugf("Status：%v:%v", response.Status, response.Proto)

	//设置新Cookie
	//Cookieb = MergeCookie(Cookieb, response.Cookies())

	//返回内容
	body, e = ioutil.ReadAll(response.Body)

	return
}

// Post附带信息
func (this *Spider)  Post() (body []byte, e error) {
	Wait(this.Wait)

	Log.Debug("POST url:" + this.Url)

	var request = &http.Request{}
	if this.Data != nil {
		pr := ioutil.NopCloser(strings.NewReader(this.Data.Encode()))
		request, _ = http.NewRequest("POST", this.Url, pr)
	} else {
		request, _ = http.NewRequest("POST", this.Url, nil)
	}
	request.Header = CloneHeader(this.Header)

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	OutputMaps("---------request header--------", request.Header)

	response, err := Client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	OutputMaps("----------response header-----------", response.Header)
	Log.Debugf("Status：%v:%v", response.Status, response.Proto)

	body, e = ioutil.ReadAll(response.Body)

	//设置新Cookie
	//MergeCookie(Cookieb, response.Cookies())

	return
}

func (this *Spider) NewHeader(ua string, host string, refer string) {
	this.Header = NewHeader(ua, host, refer)
}

func (this *Spider) Log() *logging.Logger {
	return Log
}