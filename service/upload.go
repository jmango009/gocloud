package service

import (
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"gocloud/common/config"
	"gocloud/common/logger"
	"gocloud/common/utils"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"gocloud/model"
	"gocloud/common/database"
)

var db = database.DB

func UploadImagePost(c *gin.Context) {
	maxSize := int64(config.Config.MaxSizeByMb_Image * 1024 * 1024)
	if c.Request.ContentLength > maxSize {
		logger.Error("file too large! size: ", c.Request.ContentLength)
		c.Writer.WriteHeader(http.StatusRequestEntityTooLarge)
		return
	}

	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxSize)
	c.Request.ParseMultipartForm(maxSize)

	file, handle, err := c.Request.FormFile("file")
	checkErr(err)
	defer file.Close()

	fileHeader := make([]byte, 512)
	_, err = file.Read(fileHeader)
	checkErr(err)
	if !utils.IsImage(fileHeader) {
		logger.Error("File is not an image: " + handle.Filename)
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = file.Seek(0, 0)
	checkErr(err)

	filePath := filepath.Join(config.ImageDir, handle.Filename)
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	checkErr(err)
	defer f.Close()
	io.Copy(f, file)

	go makeThumbnail(handle.Filename)

	var image model.Image
	c.Bind(&image)
	image.Filename = handle.Filename
	db.Create(&image)

	logger.Info("uploaded file: " + handle.Filename)
}

func UploadVideoPost(c *gin.Context) {
	maxSize := int64(config.Config.MaxSizeByMb_Video * 1024 * 1024)
	if c.Request.ContentLength > maxSize {
		logger.Error("file too large! size: ", c.Request.ContentLength)
		c.Writer.WriteHeader(http.StatusRequestEntityTooLarge)
		return
	}

	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxSize)
	c.Request.ParseMultipartForm(maxSize)

	file, handle, err := c.Request.FormFile("file")
	checkErr(err)
	defer file.Close()

	fileHeader := make([]byte, 512)
	_, err = file.Read(fileHeader)
	checkErr(err)
	if !utils.IsVideo(fileHeader) {
		logger.Error("File is not a video: " + handle.Filename)
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = file.Seek(0, 0)
	checkErr(err)

	filePath := filepath.Join(config.VideoDir, handle.Filename)
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	checkErr(err)
	defer f.Close()
	io.Copy(f, file)

	var video model.Video
	c.Bind(&video)
	video.Filename = handle.Filename
	db.Create(&video)

	logger.Info("uploaded file: " + handle.Filename)
}

func checkErr(err error) {
	if err != nil {
		logger.Panic(err.Error())
	}
}

func makeThumbnail(filename string) {
	img, err := imaging.Open(filepath.Join(config.ImageDir, filename))
	if err != nil {
		logger.Error("%v", err)
		return
	}
	thumb := imaging.Resize(img, 200, 0, imaging.CatmullRom)

	imaging.Save(thumb, filepath.Join(config.ThumbnailDir, filename))
}