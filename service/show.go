package service

import (
	"github.com/gin-gonic/gin"
	"gocloud/model"
)

type MediaList struct {
	MediaList []MediasByDate
}

type MediasByDate struct {
	Date string
	MediaInfos []MediaInfo
}

type MediaInfo struct {
	ThumbnailUrl string
	MediaUrl string
}

func GetImageList(c *gin.Context) {
	images := []model.Image{}
	db.Find(&images)

	imageByDateMap := map[string][]MediaInfo{}
	for _, image := range images {
		createdDate := image.CreatedAt.Format("2006-01-02")
		imageInfo := MediaInfo{
			ThumbnailUrl: "/t/tn/" + image.Filename,
			MediaUrl: "/t/img/" +  image.Filename,
		}
		if imageInfos, ok := imageByDateMap[createdDate]; ok {
			imageInfos = append(imageInfos, imageInfo)
			imageByDateMap[createdDate] = imageInfos
		} else {
			imageInfos = []MediaInfo{imageInfo}
			imageByDateMap[createdDate] = imageInfos
		}
	}
	imageList := []MediasByDate{}
	for key, images := range imageByDateMap {
		imageByDate := MediasByDate{Date:key, MediaInfos:images}
		imageList = append(imageList, imageByDate)
	}

	c.JSON(200, gin.H{
		"imageList": imageList,
	})

}

func GetVideoList(c *gin.Context) {
	videos := []model.Video{}
	db.Find(&videos)

	videoByDateMap := map[string][]MediaInfo{}
	for _, video := range videos {
		createdDate := video.CreatedAt.Format("2006-01-02")
		imageInfo := MediaInfo{
			MediaUrl: "/t/v/" + video.Filename,
		}
		if videoInfos, ok := videoByDateMap[createdDate]; ok {
			videoInfos = append(videoInfos, imageInfo)
			videoByDateMap[createdDate] = videoInfos
		} else {
			videoInfos = []MediaInfo{imageInfo}
			videoByDateMap[createdDate] = videoInfos
		}
	}
	videoList := []MediasByDate{}
	for key, videos := range videoByDateMap {
		videoByDate := MediasByDate{Date:key, MediaInfos:videos}
		videoList = append(videoList, videoByDate)
	}

	c.JSON(200, gin.H{
		"videoList": videoList,
	})

}
