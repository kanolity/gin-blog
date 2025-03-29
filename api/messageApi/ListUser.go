package messageApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
	"go_code/gin-vue-blog/utils/jwts"
	"time"
)

type Message struct {
	ID               uint      `gorm:"primary_key" json:"id"`
	CreatedAt        time.Time `json:"created_at"`
	SendUserID       uint      `json:"send_user_id"`
	SendUserNickName string    `json:"send_user_nick_name"`
	SendUserAvatar   string    `json:"send_user_avatar"`
	RevUserID        uint      `json:"rev_user_id"`
	RevUserNickName  string    `json:"rev_user_nick_name"`
	RevUserAvatar    string    `json:"rev_user_avatar"`
	IsRead           bool      `json:"is_read"`
	Content          string    `json:"content"`
}

func (messageApi *MessageApi) MessageListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var msgList []models.Message
	err := global.DB.Raw(`
        SELECT t.*
        FROM messages t
        JOIN (
            SELECT 
                LEAST(send_user_id, rev_user_id) AS user1,
                GREATEST(send_user_id, rev_user_id) AS user2,
                MAX(created_at) AS latest_time
            FROM messages
            WHERE send_user_id = ? OR rev_user_id = ?  -- 当前用户参与的聊天
            GROUP BY user1, user2
        ) grouped
        ON LEAST(t.send_user_id, t.rev_user_id) = grouped.user1
           AND GREATEST(t.send_user_id, t.rev_user_id) = grouped.user2
           AND t.created_at = grouped.latest_time
        ORDER BY t.created_at DESC
    `, claims.UserID, claims.UserID).Scan(&msgList).Error
	if err != nil {
		res.FailWithMsg("消息列表出错", c)
		return
	}
	res.OKWithData(msgList, c)
}
