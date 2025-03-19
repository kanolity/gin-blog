package models

type Message struct {
	SendUserID       uint   `gorm:"primaryKey" json:"send_user_id"`     // 发送人 ID
	SendUserModel    User   `gorm:"foreignKey:SendUserID" json:"-"`     // 发送人模型
	SendUserNickName string `gorm:"size:42" json:"send_user_nick_name"` // 发送人昵称
	SendUserAvatar   string `json:"send_user_avatar"`                   // 发送人头像

	RevUserID       uint   `gorm:"primaryKey" json:"rev_user_id"`     // 接收人 ID
	RevUserModel    User   `gorm:"foreignKey:RevUserID" json:"-"`     // 接收人模型
	RevUserNickName string `gorm:"size:42" json:"rev_user_nick_name"` // 接收人昵称
	RevUserAvatar   string `json:"rev_user_avatar"`                   // 接收人头像

	IsRead  bool   `gorm:"default:false" json:"is_read"` // 是否已读
	Content string `json:"content"`                      // 消息内容
}
