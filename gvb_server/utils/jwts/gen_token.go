package jwts

import (
	"github.com/golang-jwt/jwt/v5"
	"gvb_server/global"
	"time"
)

// GenToken 创建 Token
func GenToken(user JwtPayLoad) (string, error) {
	MySecret = []byte(global.Config.Jwt.Secret)
	claim := CustomClaims{
		user,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.Config.Jwt.Expires) * time.Hour)), // 默认2小时过期
			Issuer:    global.Config.Jwt.Issuer,                                                                 // 签发人
		},
	}

	//生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	//加密
	return token.SignedString(MySecret)
}
