package RedisService

import (
	"fmt"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/utils"
	"time"
)

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
