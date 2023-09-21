package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/gin-gonic/gin"
)

// SHA1 sha1对字符串加密
func SHA1(s string) string {
	o := sha1.New()
	o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

// SHAMiddleWare 对密码进行加密
func SHAMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		password := context.Query("password")
		if password == "" {
			password = context.PostForm("password")
		}
		//map
		context.Set("password", SHA1(password))
		context.Next()
	}
}

// CheckPassword 判断密码是否一直 前者为需要判断的明文密码 后者为密文密码
func CheckPassword(pwd, password string) bool {
	if hashPassword := SHA1(pwd); hashPassword != password {
		return false
	}
	return true
}
