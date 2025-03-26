package desens

import "strings"

// DesensitizationPhone 电话数据脱敏
func DesensitizationPhone(phone string) string {
	if len(phone) != 11 {
		return ""
	}
	return phone[:3] + "xxxx" + phone[7:]
}

// DesensitizationEmail 邮箱数据脱敏
func DesensitizationEmail(email string) string {
	eList := strings.Split(email, "@")
	if len(eList) != 2 {
		return ""
	}
	return eList[0][:1] + "xxxx@" + eList[1]
}
