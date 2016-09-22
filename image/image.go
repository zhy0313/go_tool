package image

import (
	"errors"
	"github.com/hunterhug/go-image/graphics"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

var (
	ExtNotSupportError = errors.New("ext of filename not support")
	FileNameError = errors.New("filename error")
	FileExistError = errors.New("File already exist error")
)

//按宽度和高度进行比例缩放
func ThumbnailF2F(filename string, savepath string, width int, height int) (err error) {
	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	src, filetype, err := LoadImage(filename)
	if err != nil {
		return
	}
	err = graphics.Thumbnail(dst, src)
	if err != nil {
		return
	}
	err = SaveImage(savepath, dst, filetype)
	return
}

//按宽度进行比例缩放
func ScaleF2F(filename string, savepath string, width int) (err error) {
	img, filetype, err := Scale(filename, width)
	if err != nil {
		return
	}
	err = SaveImage(savepath, img, filetype)
	if err != nil {
		return
	}
	return
}

//图像文件的真正名字
func RealImageName(filename string) (filenewname string, err error) {
	_, ext, err := LoadImage(filename)
	if err != nil {
		return
	}
	temp := strings.Split(filename, ".")
	if len(temp) < 2 {
		err = FileNameError
	}
	temp[len(temp) - 1] = ext
	filenewname = strings.Join(temp, ".")
	return
}

//文件改名,如果force为假,且新的文件名已经存在,那么抛出错误
func ChangeImageName(oldname string, newname string, force bool) (err error) {
	if !force {
		_, err = os.Open(newname)
		if err == nil {
			err = FileExistError
			return
		}
	}
	err = os.Rename(oldname, newname)
	return

}


// 根据文件名打开图片,并编码,返回编码对象和文件类型
func LoadImage(path string) (img image.Image, filetype string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	img, filetype, err = image.Decode(file)
	return
}

// 将编码对象存入文件中
func SaveImage(path string, img *image.RGBA, filetype string) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return
	}
	if filetype == "png" {
		err = png.Encode(file, img)
	} else if filetype == "jpeg" {
		err = jpeg.Encode(file, img, nil)
	} else {
		err = ExtNotSupportError
	}
	defer file.Close()
	return
}

//对文件中的图片进行等比例变化,宽度为newdx,返回图像编码和文件类型
func Scale(filename string, newdx int) (dst *image.RGBA, filetype string, err error) {
	src, filetype, err := LoadImage(filename)
	if err != nil {
		return
	}
	bound := src.Bounds()
	dx := bound.Dx()
	dy := bound.Dy()
	dst = image.NewRGBA(image.Rect(0, 0, newdx, newdx * dy / dx))
	// 产生缩略图,等比例缩放
	err = graphics.Scale(dst, src)
	return
}
