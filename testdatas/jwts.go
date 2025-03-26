package main

import (
	"fmt"
	"go_code/gin-vue-blog/core"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/utils/jwts"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()
	str, err := jwts.GenToken(jwts.JwtPayLoad{
		Username: "test",
		Nickname: "test",
		Role:     1,
		UserID:   1,
	})
	fmt.Println()
	fmt.Println(str, err)

	claims, err := jwts.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QiLCJuaWNrbmFtZSI6InRlc3QiLCJyb2xlIjoxLCJ1c2VyX2lkIjoxLCJleHAiOjE3NDMwNzk2OTQuNDQyNDg5MSwiaXNzIjoieHgifQ.MeE5S3gK6cqVJEgnWyNzS2O0QqeC-tCjA3pU7hDu5m4")

	fmt.Println(claims, err)
}
