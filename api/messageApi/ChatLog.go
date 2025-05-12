package messageApi

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-vue-blog/global"
	"go_code/gin-vue-blog/models"
	"go_code/gin-vue-blog/models/res"
	"go_code/gin-vue-blog/utils/jwts"
)

func (messageApi *MessageApi) ChatLogView(c *gin.Context) {
	//获取用户ID
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	id := c.Query("id") //聊天人id

	var msgList []models.Message
	err := global.DB.Raw(`
        SELECT t.*
        FROM messages t
        JOIN (
            SELECT 
                LEAST(send_user_id, rev_user_id) AS user1,
                GREATEST(send_user_id, rev_user_id) AS user2
            FROM messages
            WHERE (send_user_id = ? OR rev_user_id = ?) AND (send_user_id = ? OR rev_user_id = ?)  -- 当前用户参与的聊天
            GROUP BY user1, user2
        ) grouped
        ON LEAST(t.send_user_id, t.rev_user_id) = grouped.user1
           AND GREATEST(t.send_user_id, t.rev_user_id) = grouped.user2
        ORDER BY t.created_at DESC
    `, claims.UserID, claims.UserID, id, id).Scan(&msgList).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("消息列表出错", c)
		return
	}
	res.OKWithData(msgList, c)
}
