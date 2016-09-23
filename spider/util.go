/*
 * Created by 一只尼玛 on 2016/8/13.
 * 功能： 
 *
 */
package spider

import (
	"time"
	"net/http"
	"fmt"
	"errors"
)

func Wait(waittime time.Duration) {
	if waittime <= 0 {
		return
	} else {
		Log.Debugf("Stop %d Second～～", waittime)
		time.Sleep(waittime * time.Second)
	}
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


//文件太小，访问太过频繁，死！！！kb
func TooSortSizes(data []byte, sizes float64) error {
	if float64(len(data)) / 1000 < sizes {
		return errors.New(fmt.Sprintf("FileSize:%d bytes,%d kb < %f kb dead too sort", len(data), len(data) / 1000, sizes))
	}
	return nil
}

//打印映射
func OutputMaps(info string, args map[string][]string) {
	Log.Debugf("%s:%v", info, args)
}

