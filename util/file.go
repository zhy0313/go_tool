/*
 * Created by 一只尼玛 on 2016/8/12.
 * 功能： 文件帮助功能
 *
 */
package util

import (
	"io/ioutil"
	"os"
	"strings"
	"errors"
	"path/filepath"
	"runtime"
)

// 获取调用者的当前文件DIR
func CurDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return filepath.Dir(filename)
}

// 将字节数组保存到文件中去
func SaveTofile(filepath string, content []byte) error {
	//全部权限写入文件
	err := ioutil.WriteFile(filepath, content, 0777)
	return err
}

//根据传入文件夹名字递归新建文件夹
func MakeDir(filedir string) error {
	return os.MkdirAll(filedir, 777)
}

//根据传入文件名，递归创建文件夹
func MakeDirByFile(filepath string) error {
	temp := strings.Split(filepath, "/")
	if len(temp) < 2 {
		return errors.New("请传入类似filedir/filename形式的文件名")
	}
	dirpath := strings.Join(temp[0:len(temp) - 1], "/")
	return MakeDir(dirpath)
}

