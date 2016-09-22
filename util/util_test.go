package util

import (
	"testing"
	"fmt"
)

func TestUtil(t *testing.T) {
	i := 2
	if ("2" == IS(i)) {
		fmt.Println("int to string")
	}

	v, err := SI("e2")
	if (err == nil&&v == i) {
		fmt.Println("string to int")
	} else {
		fmt.Println(err.Error())
	}

	fmt.Println(CurDir())

	err = MakeDir("../data")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("dir already exist")
	}

	err = MakeDirByFile("../data/testutil.txt")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("dir already exist")
	}

	err = SaveTofile("../data/testutil.txt", []byte("testutil"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
