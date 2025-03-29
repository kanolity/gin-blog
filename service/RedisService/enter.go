package RedisService

import (
	"fmt"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/utils"
	"time"
)

type RedisService struct{}

// Logout 将token存入redis
func (redisService *RedisService) Logout(token string, diff time.Duration) error {
	err := global.Redis.Set(fmt.Sprintf("logout_%s", token), "", diff).Err()
	return err
}

// CheckLogout 检测token是否在redis中（即注销）
func (redisService *RedisService) CheckLogout(token string) bool {
	keys := global.Redis.Keys("logout_*").Val()
	if utils.InList("logout_"+token, keys) {
		return true
	}
	return false
}

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
