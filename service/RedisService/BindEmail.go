package RedisService

import (
	"go_code/gin-vue-blog/global"
	"time"
)

// BindEmail 将邮箱和验证码存入redis，以便校验，30分钟过期
func (redisService *RedisService) BindEmail(code, email string) error {
	key := email + ":" + code
	err := global.Redis.Set(key, "", 30*time.Minute).Err()
	return err
}

// CheckBind 检测校验数据是否在redis中
func (redisService *RedisService) CheckBind(code, email string) bool {
	key := email + ":" + code
	err := global.Redis.Get(key).Err()
	if err != nil {
		return false
	}
	return true
}
