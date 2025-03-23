package ctype

import "encoding/json"

type Role int

const (
	Admin        Role = 1 //管理员
	NormalUser   Role = 2 //普通用户
	Visitor      Role = 3 //游客
	DisabledUser Role = 4 //禁言用户
)

func (r Role) MarshJSON() ([]byte, error) {
	return json.Marshal(r.String())
}
func (r Role) String() string {
	var str string
	switch r {
	case Admin:
		str = "管理员"
	case NormalUser:
		str = "普通用户"
	case Visitor:
		str = "游客"
	case DisabledUser:
		str = "禁言用户"
	default:
		str = "其他"
	}
	return str
}
