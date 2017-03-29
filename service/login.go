package service

import (
	"github.com/gin-gonic/gin"
	"gocloud/common/config"
	"gocloud/common/utils"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"time"
	"crypto/md5"
	"encoding/hex"
)

func LoginPost(c *gin.Context) {
	passwd_md5_bytes := md5.Sum([]byte(config.Config.Password))
	passwd_md5_hex := hex.EncodeToString(passwd_md5_bytes[:])
	if c.PostForm("password") != passwd_md5_hex {
		c.Redirect(http.StatusSeeOther, "/login.html")
		return
	}

	// Expires the token and cookie in 24 hours
	expireToken := time.Now().Add(time.Hour * 24).Unix()
	expireCookie := time.Now().Add(time.Hour * 24)


	claims := utils.MyCustomClaims {
		"admin",
		jwt.StandardClaims {
			ExpiresAt: expireToken,
			Issuer: "kunmingli.com",
		},
	}

	// Create the token using your claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Signs the token with a secret.
	signedToken, _ := token.SignedString([]byte("sheiyebiexiangpojiewomima"))

	// This cookie will store the token on the client side
	cookie := http.Cookie{Name: "Auth", Value: signedToken, Expires: expireCookie, HttpOnly: true}
	http.SetCookie(c.Writer, &cookie)
}
