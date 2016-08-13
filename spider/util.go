/*
 * Created by 一只尼玛 on 2016/8/13.
 * 功能： 
 *
 */
package spider

import (
	"time"
	"log"
	"net/http"
)

func Wait(waittime time.Duration) {
	if waittime < 0 {
		return
	} else {
		log.Printf("暂停%d秒～～", waittime)
		time.Sleep(waittime * time.Second)
	}
}

//默认休眠时间
func WaitM() {
	log.Printf("暂停%d秒～～", DeadTime)
	time.Sleep(DeadTime * time.Second)
}

//Header map[string][]string
func CopyM(h http.Header) http.Header {
	h2 := make(http.Header, len(h))
	for k, vv := range h {
		vv2 := make([]string, len(vv))
		copy(vv2, vv)
		h2[k] = vv2
	}
	return h2
}


//文件太小，访问太过频繁，死！！！
func TooSortSizes(data []byte, sizes float64) {
	if float64(len(data)) / 1000 < sizes {
		log.Printf("文件大小:%d字节,所以死掉了", len(data))
		panic("抓完了")
	}
}

//打印映射
func OutputMaps(info string, args map[string][]string) {
	log.Print(info)
	for i, v := range args {
		log.Printf("%s:%v", i, v)
	}
}

