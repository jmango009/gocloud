package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GotoHome(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, "/image_list.html")
}
