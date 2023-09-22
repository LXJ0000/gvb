package redis_service

import (
	"gvb_server/global"
	"gvb_server/utils"
	"time"
)

type RedisService struct {
}

const prefix = "logout_"

// Logout 针对注销操作
func (RedisService) Logout(token string, diff time.Duration) error {
	return global.Redis.Set(prefix+token, "", diff).Err()
}

// CheckLogout 判断是否注销了
func (RedisService) CheckLogout(token string) bool {
	//判断是否在redis中 若是则注销了
	keys := global.Redis.Keys(prefix + "*").Val()
	if utils.InList("logout_"+token, keys) {
		return true
	}
	return false
}
