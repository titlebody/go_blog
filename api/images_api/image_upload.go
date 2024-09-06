package images_api

import (
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model/res"
	"go_blog/service"
	"go_blog/service/image_ser"
	"io/fs"
	"os"
)

// ImageUploadView 上传多个文件
func (ImagesApi) ImageUploadView(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage("不存在的文件", c)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("不存在的文件", c)
		return
	}
	//判断路径是否存在
	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		//递归创建
		err := os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}
	// 不存在就创建
	var resList []image_ser.FileUploadResponse

	//fileHeader, err := c.FormFile("image")
	for _, file := range fileList {
		// 上传文件
		serviceRes := service.ServiceApp.ImageServiceGroup.ImageUploadService(file)
		if !serviceRes.IsSuccess {
			resList = append(resList, serviceRes)
			continue
		}
		// 成功
		// 本地上传
		if !global.Config.QiNiu.IsEnable {
			err = c.SaveUploadedFile(file, serviceRes.FileName)
			if err != nil {
				global.Log.Error("文件保存失败", err)
				serviceRes.Msg = "文件保存失败"
				serviceRes.IsSuccess = false
				resList = append(resList, serviceRes)
				continue
			}
		}

		resList = append(resList, serviceRes)

	}
	res.OkWithData(resList, c)
}
