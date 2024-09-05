package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_blog/global"
	"go_blog/model/res"
	"go_blog/utils"
	"io/fs"
	"os"
	"path"
	"strings"
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`  //文件名
	IsSuccess bool   `json:"is_success"` //是否上传成功
	Msg       string `json:"msg"`        //消息
}

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
	//pathList := strings.Split(global.Config.Upload.Path, "/")
	_, err = os.ReadDir(global.Config.Upload.Path)

	if err != nil {
		//递归创建
		err := os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}
	// 不存在就创建
	var resList []FileUploadResponse

	//fileHeader, err := c.FormFile("image")
	for _, file := range fileList {

		fileName := file.Filename
		nameList := strings.Split(fileName, ".")
		suffix := strings.ToLower(nameList[len(nameList)-1])
		/*for _, v := range global.WhiteImageList {
			if strings.EqualFold(suffix, v) {
				res.FailWithMessage("文件类型不允许", c)
				return
			}
		}*/
		if !utils.InList(suffix, global.WhiteImageList) {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "文件类型不允许",
			})
			continue
		}
		filePath := path.Join(basePath, file.Filename)
		// 判断大小
		size := float64(file.Size) / float64(1024*1024)
		if size >= float64(global.Config.Upload.Size) {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("文件大小超出限制,当前大小为:%.2fMB设置的大小为:%dMB", size, global.Config.Upload.Size),
			})
			continue
		}
		resList = append(resList, FileUploadResponse{
			FileName:  filePath,
			IsSuccess: true,
			Msg:       "上传成功",
		})
		err := c.SaveUploadedFile(file, filePath)
		if err != nil {
			global.Log.Error("文件保存失败", err)
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "上传失败",
			})
			continue
		}
		resList = append(resList, FileUploadResponse{
			FileName:  filePath,
			IsSuccess: true,
			Msg:       "上传成功",
		})

	}
	res.OkWithData(resList, c)
}
