package ctype

import "encoding/json"

type SignStatus int

const (
	QQ    SignStatus = 1 //QQ
	Gitee SignStatus = 2 //gitee
	Email SignStatus = 3 //邮箱
)

func (s SignStatus) MarshJSON() ([]byte, error) {
	return json.Marshal(s.String())
}
func (s SignStatus) String() string {
	var str string
	switch s {
	case QQ:
		str = "QQ"
	case Gitee:
		str = "gitee"
	case Email:
		str = "邮箱"
	default:
		str = "其他"
	}
	return str
}
