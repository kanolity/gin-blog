package ImageService

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/ctype"
	"go_code/gin-vue-blog/utils"
	"io"
	"mime/multipart"
	"path"
	"strings"
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`  //文件名
	Msg       string `json:"msg"`        //消息
	IsSuccess bool   `json:"is_success"` //是否上传成功
}

var (
	// WhiteImageList 图片上传白名单
	WhiteImageList = []string{
		"jpg",
		"jpeg",
		"png",
		"gif",
		"ico",
		"webp",
		"tiff",
		"svg",
	}
)

// ImageUploadService 文件上传方法
func (ImageService *ImageService) ImageUploadService(file *multipart.FileHeader, c *gin.Context) (res FileUploadResponse) {
	res.FileName = file.Filename
	res.IsSuccess = false

	//判断图片是否合法
	nameList := strings.Split(file.Filename, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1])
	if !utils.InList(suffix, WhiteImageList) {
		return FileUploadResponse{
			Msg: "非法文件",
		}
	}

	filepath := path.Join(global.Config.Upload.Path, file.Filename)

	//判断大小
	size := float64(file.Size) / float64(1024*1024)
	if size >= float64(global.Config.Upload.Size) {
		msg := fmt.Sprintf("图片大小超过限制，图片大小为%.2fMB,限制为%dMB", size, global.Config.Upload.Size)
		return FileUploadResponse{
			Msg: msg,
		}
	}

	fileObj, err := file.Open()
	if err != nil {
		global.Log.Error(err)
	}
	byteData, err := io.ReadAll(fileObj)
	imageHash := utils.Md5(byteData)
	//去数据库中查图片是否存在
	var bannerModel models.Banner
	err = global.DB.Take(&bannerModel, "hash = ?", imageHash).Error
	if err == nil {
		//已存在图片
		return FileUploadResponse{
			FileName: bannerModel.Path,
			Msg:      "图片已存在",
		}
	}

	//存储文件
	err = c.SaveUploadedFile(file, filepath) //若路径不存在SaveUploadedFile函数自动创建
	if err != nil {
		global.Log.Error(err)
		return FileUploadResponse{
			Msg: err.Error(),
		}
	}

	//写入数据库
	global.DB.Create(&models.Banner{
		Path:      filepath,
		Hash:      imageHash,
		Name:      file.Filename,
		ImageType: ctype.Local,
	})
	return FileUploadResponse{
		FileName:  filepath,
		IsSuccess: true,
		Msg:       "上传成功",
	}
}
