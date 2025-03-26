package jwts

import (
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"go_code/gin-vue-blog/global"
	"time"
)

// GenToken 创建Token
func GenToken(user JwtPayLoad) (string, error) {
	MySecret = []byte(global.Config.Jwt.Secret)
	fmt.Println(MySecret)
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(global.Config.Jwt.Expires))), //默认2小时过期
			Issuer:    global.Config.Jwt.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(MySecret)
}
