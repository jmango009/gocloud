package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gocloud/model"
	"html/template"
	"fmt"
)

func ArticleGet(c *gin.Context) {
	c.HTML(http.StatusOK, "article.html", nil)
}

func ArticleShow(c *gin.Context) {
	id := c.Param("id")

	article := &model.Article{}
	db.First(article, id)
	params := map[string]template.HTML{"articleContent": template.HTML(article.Content)}
	c.HTML(http.StatusOK, "article_show.html", params)
}

func ArticlePost(c *gin.Context) {
	content := c.PostForm("article")

	var article model.Article
	article.Content = content

	db.Create(&article)

	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/article/%v", article.ID))
}
