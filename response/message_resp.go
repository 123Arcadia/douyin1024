package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// MessageActionResponse 发送消息响应
type MessageActionResponse struct {
	StatusCode string `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

// MessageActionResponseHandler 发送消息
func MessageActionResponseHandler(c *gin.Context, code string, msg string) {
	var httpStat int
	if code == "0" {
		httpStat = http.StatusOK
	} else {
		httpStat = http.StatusInternalServerError
	}
	c.JSON(httpStat, MessageActionResponse{
		StatusCode: code,
		StatusMsg:  msg,
	})
}

type Message struct {
	Content    string `json:"content"`      // 消息内容
	CreateTime int64  `json:"create_time"`  // 消息发送时间 yyyy-MM-dd HH:MM:ss
	FromUserID int64  `json:"from_user_id"` // 消息发送者id
	ID         int64  `json:"id"`           // 消息id
	ToUserID   int64  `json:"to_user_id"`   // 消息接收者id
}

// 聊天记录响应
type MessageChatResponse struct {
	StatusCode  string    `json:"status_code"`            // 状态码，0-成功，其他值-失败
	StatusMsg   string    `json:"status_msg"`             // 返回状态描述
	MessageList []Message `json:"message_list,omitempty"` // 用户列表

}

// MessageChatResponseHandler 聊天记录获取
func MessageChatResponseHandler(c *gin.Context, code string, msg string, msgs []Message) {
	var httpStat int
	if code == "0" {
		httpStat = http.StatusOK
	} else {
		httpStat = http.StatusInternalServerError
	}
	c.JSON(httpStat, MessageChatResponse{
		StatusCode:  code,
		StatusMsg:   msg,
		MessageList: msgs,
	})
}
