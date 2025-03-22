package imagesApi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
	"go_code/gin-vue-blog/utils"
	"io"
	"os"
	"path"
	"strings"
)

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

type FileUploadResponse struct {
	FileName  string `json:"file_name"`  //文件名
	Msg       string `json:"msg"`        //消息
	IsSuccess bool   `json:"is_success"` //是否上传成功
}

// appendList 汇总每张图片上传的响应
func appendList(resList []FileUploadResponse, filename string, isSuccess bool, msg string) []FileUploadResponse {
	resList = append(resList, FileUploadResponse{
		FileName:  filename,
		IsSuccess: isSuccess,
		Msg:       msg,
	})
	return resList
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
		//判断图片是否合法
		nameList := strings.Split(file.Filename, ".")
		suffix := strings.ToLower(nameList[len(nameList)-1])
		if !utils.InList(suffix, WhiteImageList) {
			resList = appendList(resList, file.Filename, false, "非法文件")
			continue
		}

		filepath := path.Join(global.Config.Upload.Path, file.Filename)

		//判断大小
		size := float64(file.Size) / float64(1024*1024)
		if size >= float64(global.Config.Upload.Size) {
			msg := fmt.Sprintf("图片大小超过限制，图片大小为%.2fMB,限制为%dMB", size, global.Config.Upload.Size)
			resList = appendList(resList, file.Filename, false, msg)
			continue
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
			resList = appendList(resList, bannerModel.Path, false, "图片已存在")
			continue
		}

		//存储文件
		err = c.SaveUploadedFile(file, filepath) //若路径不存在SaveUploadedFile函数自动创建
		if err != nil {
			global.Log.Error(err)
			resList = appendList(resList, file.Filename, false, err.Error())
			continue
		}

		resList = appendList(resList, filepath, true, "上传成功")

		//写入数据库
		global.DB.Create(&models.Banner{
			Path: filepath,
			Hash: imageHash,
			Name: file.Filename,
		})
	}
	res.OKWithData(resList, c)
}
