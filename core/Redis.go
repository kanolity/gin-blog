package core

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"go_code/gin-vue-blog/global"
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
		logrus.Errorf("redis 连接失败:%s", redisConf.Addr())
		logrus.Errorf("redis 连接失败:%v", err)
	}
	return rdb
}
