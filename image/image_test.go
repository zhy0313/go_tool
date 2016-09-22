package image

import (
	"testing"
	"fmt"
)

func TestImage(t *testing.T) {
	err := ThumbnailF2F("../data/image.png", "../data/image100*100.png", 100, 100)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = ScaleF2F("../data/image.png", "../data/image200.png", 200)
	if err != nil {
		fmt.Println(err.Error())
	}

	filename, err := RealImageName("../data/image.png")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("real filename"+filename)
	}
}
