package routes

import (
	"github.com/gin-gonic/gin"
	"gocloud/common/utils"
	"gocloud/service"
	"net/http"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"gocloud/common/config"
	"path/filepath"
)

func Router() *http.ServeMux {
	router := gin.Default()
	gin.SetMode(gin.DebugMode)

	//router.LoadHTMLGlob(filepath.Join(config.Config.TemplateDir + "/*"))

	router.Use(authentication())

	router.GET("/", service.GotoHome)

	router.POST("/login", service.LoginPost)

	router.POST("/image", service.UploadImagePost)
	router.POST("/video", service.UploadVideoPost)
	router.GET("/image/all", service.GetImageList)
	router.GET("/video/all", service.GetVideoList)

	//router.GET("/article", service.ArticleGet)
	//router.GET("/article/:id", service.ArticleShow)
	//router.POST("/article", service.ArticlePost)

	router.StaticFile("/login.html", filepath.Join(config.ViewDir, "login.html"))
	router.StaticFile("/image_list.html", filepath.Join(config.ViewDir, "image_list.html"))
	router.StaticFile("/video_list.html", filepath.Join(config.ViewDir, "video_list.html"))
	router.StaticFile("/upload_image.html", filepath.Join(config.ViewDir, "upload_image.html"))
	router.StaticFile("/upload_video.html", filepath.Join(config.ViewDir, "upload_video.html"))

	rootMux := http.NewServeMux()
	rootMux.Handle("/", router)

	rootMux.Handle("/js/",
		http.StripPrefix("/js/", http.FileServer(http.Dir(filepath.Join(config.Config.RootDir, "static/js")))))
	rootMux.Handle("/css/",
		http.StripPrefix("/css/", http.FileServer(http.Dir(filepath.Join(config.Config.RootDir, "static/css")))))
	//rootMux.Handle("/ckeditor/",
	//	http.StripPrefix("/ckeditor/", http.FileServer(http.Dir("static/ckeditor"))))
	rootMux.Handle("/i/",
		http.StripPrefix("/i/", http.FileServer(http.Dir(filepath.Join(config.Config.RootDir, "static/i")))))
	rootMux.Handle("/fonts/",
		http.StripPrefix("/fonts/", http.FileServer(http.Dir(filepath.Join(config.Config.RootDir, "static/fonts")))))

	rootMux.Handle("/img/",
		http.StripPrefix("/img/", http.FileServer(http.Dir(config.ImageDir))))
	rootMux.Handle("/tn/",
		http.StripPrefix("/tn/", http.FileServer(http.Dir(config.ThumbnailDir))))
	rootMux.Handle("/v/",
		http.StripPrefix("/v/", http.FileServer(http.Dir(config.EncodedVideoDir))))

	return rootMux
}

func authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		urlPath := c.Request.URL.Path
		if !strings.HasPrefix(urlPath, "/upload") {
			c.Next()
			return
		}

		// If no Auth cookie is set then return a 404 not found
		cookieValue, err := c.Cookie("Auth")
		if err != nil {
			c.Redirect(http.StatusSeeOther, "/login.html")
			return
		}

		// Parse, validate and return a token.
		token, err := jwt.ParseWithClaims(cookieValue, &utils.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error){
			// Prevents a known exploit
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
				return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
			}
			return []byte("sheiyebiexiangpojiewomima"), nil
		})

		// Validate the token and save the token's claims to a context
		if claims, ok := token.Claims.(*utils.MyCustomClaims); ok && token.Valid {
			c.Set("Claims", claims)
		} else {
			c.Redirect(http.StatusSeeOther, "/login.html")
			return
		}

		c.Next()
	}
}