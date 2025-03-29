package messageApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
	"go_code/gin-vue-blog/utils/jwts"
)

type MessageRequest struct {
	RevUserID uint   `json:"rev_user_id" binding:"required"` // 接收人 ID
	Content   string `json:"content" binding:"required"`     // 消息内容
}

// MessageCreateView 发布消息
func (messageApi *MessageApi) MessageCreateView(c *gin.Context) {
	//获取发送人 ID
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var cr MessageRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//获取发送人、接收人
	var sendUser, revUser models.User
	err = global.DB.Take(&sendUser, claims.UserID).Error
	if err != nil {
		res.FailWithMsg("发送人不存在", c)
		return
	}
	err = global.DB.Take(&revUser, cr.RevUserID).Error
	if err != nil {
		res.FailWithMsg("接收人不存在", c)
		return
	}
	err = global.DB.Create(&models.Message{
		SendUserID:       sendUser.ID,
		SendUserNickName: sendUser.Nickname,
		SendUserAvatar:   sendUser.Avatar,
		RevUserID:        revUser.ID,
		RevUserNickName:  revUser.Nickname,
		RevUserAvatar:    revUser.Avatar,
		IsRead:           false,
		Content:          cr.Content,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("消息发送失败", c)
		return
	}
	res.OKWithMsg("消息发送成功", c)
}
