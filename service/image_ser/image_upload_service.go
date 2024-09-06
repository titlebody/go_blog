package image_ser

import (
	"fmt"
	"go_blog/global"
	"go_blog/model"
	"go_blog/model/c_type"
	"go_blog/plugins/qiniu"
	"go_blog/utils"
	"io"
	"mime/multipart"
	"path"
	"strings"
)

var (
	// WhiteImageList 图片上传的白名单
	WhiteImageList = []string{
		"png",
		"jpg",
		"jpeg",
		"gif",
		"webp",
		"bmp",
		"svg",
		"ico",
		"tiff",
		"tif",
		"jfif",
		"jpe",
		"jif",
	}
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`  //文件名
	IsSuccess bool   `json:"is_success"` //是否上传成功
	Msg       string `json:"msg"`        //消息
}

func (ImageService) ImageUploadService(file *multipart.FileHeader) (res FileUploadResponse) {
	fileName := file.Filename
	basePath := global.Config.Upload.Path
	filePath := path.Join(basePath, file.Filename)
	res.FileName = filePath

	// 文件白名单判断
	nameList := strings.Split(fileName, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1])
	if !utils.InList(suffix, WhiteImageList) {
		res.Msg = "文件类型不允许"
		return
	}
	//判断文件大小
	size := float64(file.Size) / float64(1024*1024)
	if size >= float64(global.Config.Upload.Size) {
		res.Msg = fmt.Sprintf("文件大小超出限制,当前大小为:%.2fMB设置的大小为:%dMB", size, global.Config.Upload.Size)
		return
	}

	// 读取文件内容 hash
	fileObj, err := file.Open()
	if err != nil {
		global.Log.Error(err)
	}
	byteDate, err := io.ReadAll(fileObj)
	imageHash := utils.Md5(byteDate)

	// 去数据库中查判断是否存在
	var bannerModel model.BannerModel
	err = global.DB.Take(&bannerModel, "hash = ?", imageHash).Error
	if err == nil {
		// 找到了
		res.Msg = "图片已存在"
		res.FileName = bannerModel.Path
		return
	}

	// 上传到七牛
	fileType := c_type.Local
	res.Msg = "图片上传成功"
	res.IsSuccess = true

	if global.Config.QiNiu.IsEnable {
		filePath, err := qiniu.UploadImages(byteDate, fileName, global.Config.QiNiu.Prefix)
		if err != nil {
			global.Log.Error(err)
			res.Msg = err.Error()
			return
		}
		res.FileName = filePath
		res.Msg = "上传七牛成功"
		fileType = c_type.QiNiu
	}

	// 图片入库
	global.DB.Create(&model.BannerModel{
		Path: res.FileName,
		Hash: imageHash,
		Name: fileName,
		Type: fileType,
	})
	return

}
