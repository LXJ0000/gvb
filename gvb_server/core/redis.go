package core

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"time"
)

func ConnectRedis() *redis.Client {
	return ConnectRedisDB(0)
}

func ConnectRedisDB(db int) *redis.Client {
	redisConf := global.Config.Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr(),
		Password: redisConf.Password,
		DB:       db,
		PoolSize: redisConf.PoolSize,
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
		logrus.Errorf("redis连接失败：%s", redisConf.Addr())
		return nil
	}
	return rdb
}

//T(n)     = T(n - 1) + n - 1
//T(n)     = T(n - 2) + (n - 2) + n - 1
//T(n)     = T(n - 3) + (n - 3) + (n - 2) + (n - 1)
//T(n)     = T(n - k) + (n - k) + ... +     (n - 1)
//k = n - 1 T(n) = T(1) + 1 + 2 + ... + (n - 1)
//T(1) = 0
