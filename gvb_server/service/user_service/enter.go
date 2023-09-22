package user_service

import (
	"github.com/golang-jwt/jwt/v5"
	"gvb_server/service/redis_service"
	"time"
)

type UserService struct {
}

// Logout 注销
func (UserService) Logout(expRaw any, token string) {
	exp, _ := expRaw.(*jwt.NumericDate)

	diff := exp.Sub(time.Now())
	_ = redis_service.RedisService{}.Logout(token, diff)
}
