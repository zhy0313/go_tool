/*
 * Created by 一只尼玛 on 2016/8/12.
 * 功能： 杂类
 *
 */
package util

import (
	"strconv"
)

//string to int
func SI(s string) (i int, e error) {
	i, e = strconv.Atoi(s)
	return
}

//int to string
func IS(i int) string {
	return strconv.Itoa(i)
}
