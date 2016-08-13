package main

import (
	. "github.com/hunterhug/go_tool/util"
	"log"
	"github.com/hunterhug/go_tool/example/testspider"
	"path/filepath"
)

func main() {
	//创建文件夹
	year := 2016
	filedir := filepath.Join(CurDir(),"../","data",IS(year))
	log.Println(filedir)
	err := MakeDir(filedir)
	if err != nil {
		log.Println(err.Error())
	}
	testspider.LoopCatUrl(filedir, year)
}
