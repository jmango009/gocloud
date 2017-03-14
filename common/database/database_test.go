package database

import (
	"fmt"
	"gocloud/model"
	"testing"
)

func Test_1(t *testing.T) {
	image1 := model.Image{Description: "Test", Filename: "test1.jpg"}
	comment1 := model.ImageComment{Content: "test comment1"}
	image1.Comments = []model.ImageComment{comment1}

	image2 := model.Image{Description: "Test", Filename: "test2.jpg"}
	comment2 := model.ImageComment{Content: "test comment2"}
	comment3 := model.ImageComment{Content: "test comment3"}
	image2.Comments = []model.ImageComment{comment2, comment3}

	//DB.Create(&image1)
	//DB.Create(&image2)

	images := []model.Image{}
	DB.Preload("Comments").Find(&images)

	for _, image := range images {
		fmt.Println(image.Filename)
		for _, comment := range image.Comments {
			fmt.Println(comment.Content)
		}
	}

	//for _, image := range images {
	//	DeleteImage(&image)
	//}
}
