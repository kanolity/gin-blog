package random

import (
	"fmt"
	"math/rand"
	"time"
)

// Code 生成四位验证码
func Code() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%4v", rand.Intn(9000)+1000)
}
