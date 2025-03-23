package imagesApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models/res"
	"go_code/gin-vue-blog/service"
	"os"
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`  //文件名
	Msg       string `json:"msg"`        //消息
	IsSuccess bool   `json:"is_success"` //是否上传成功
}

// ImageUploadView 上传图片，返回图片url
func (imagesApi *ImagesApi) ImageUploadView(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMsg("不存在的文件", c)
		return
	}

	//判断路径是否存在
	_, err = os.ReadDir(global.Config.Upload.Path)
	if err != nil {
		//递归创建
		err = os.MkdirAll(global.Config.Upload.Path, os.ModePerm)
		if err != nil {
			global.Log.Error(err.Error())
		}
	}

	var resList []FileUploadResponse

	for _, file := range fileList {
		response := service.ServiceGroupApp.ImageService.ImageUploadService(file, c)
		resList = append(resList, FileUploadResponse{
			FileName:  response.FileName,
			IsSuccess: response.IsSuccess,
			Msg:       response.Msg,
		})

	}
	res.OKWithData(resList, c)
}
