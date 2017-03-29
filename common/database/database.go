package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gocloud/model"
	"os"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	DB, err = gorm.Open("mysql", "root:admin@/gocloud?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("DB can not be initialized!")
		os.Exit(1)
	}

	DB.Debug()
	DB.AutoMigrate(&model.Image{}, &model.Video{}, &model.Article{})
	DB.AutoMigrate(&model.ImageComment{}, &model.VideoComment{}, &model.ArticleComment{})
}

func DeleteImage(image *model.Image) {
	DB = DB.Debug()

	for _, comment := range image.Comments {
		DB.Delete(&comment)
	}
	DB.Delete(&image)
}

func DeleteVideo(video *model.Video) {
	DB = DB.Debug()

	for _, comment := range video.Comments {
		DB.Delete(&comment)
	}
	DB.Delete(&video)
}

func DeleteArticle(article *model.Article) {
	DB = DB.Debug()

	for _, comment := range article.Comments {
		DB.Delete(&comment)
	}
	DB.Delete(&article)
}
