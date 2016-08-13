/*
 * Created by 一只尼玛 on 2016/8/12.
 * 功能： 网络COOKIE功能
 *
 */
package spider

import (
	"log"
	"net/http"
	"net/http/cookiejar"
)

func NewJar() *cookiejar.Jar {
	cookieJar, _ := cookiejar.New(nil)
	return cookieJar
}

var (
	Client = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			log.Printf("------------------自动跳转跳转%v-------------------\n", req.URL)
			return nil
		},
		Jar: NewJar(),
	}
	//每次访问携带的cookie
	Cookieb = []*http.Cookie{} //map[string][]string
)

//合并Cookie，后来的覆盖前来的
func MergeCookie(before []*http.Cookie, after []*http.Cookie) []*http.Cookie {
	cs := make(map[string]*http.Cookie)

	for _, b := range before {
		cs[b.Name] = b
	}

	for _, a := range after {
		if a.Value != "" {
			cs[a.Name] = a
		}
	}

	res := make([]*http.Cookie, 0, len(cs))

	for _, q := range cs {
		res = append(res, q)

	}

	return res

}

// 克隆头部
func CloneHeader(h map[string][]string) map[string][]string {
	return CopyM(h)
}
