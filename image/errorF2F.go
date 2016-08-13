package image
import "errors"
var(
	ExtNotSupportError=errors.New("载图像后不能识别的类型")
	FileNameError=errors.New("文件名错误")
	FileExistError=errors.New("文件已经存在")
)