package models

type LoginData struct {
	UserID    uint   `json:"user_id"`
	UserModel User   `gorm:"foreignkey:UserID" json:"-"`
	IP        string `gorm:"size:255" json:"ip"`
	NickName  string `gorm:"size:36" json:"nick_name"`
	Token     string `gorm:"size:255" json:"token"`
	Device    string `gorm:"size:255" json:"device"`
	Address   string `gorm:"size:255" json:"address"`
}
