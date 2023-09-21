package controller

import (
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/response"
	"github.com/123Arcadia/douyin1024CodeSpaceDemo.git/service"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type ChatResponse struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

var MessageService service.MessageService

// MessageAction 发送消息Post
func MessageAction(c *gin.Context) {
	toUserId, err := strconv.ParseUint(c.Query("to_user_id"), 10, 64)
	if err != nil {
		response.ToUserIdConversionErr(c, "参数类型转换错误")
	}
	actionType := c.Query("action_type") // 1-发送消息
	userId := c.GetUint("user_id")
	//fmt.Printf("toUserId = %+v,action_type = %+v, ctx =  %+v\n", toUserId, actionType, content)

	if actionType == "1" {
		content := c.Query("content")
		err := MessageService.AddMsg(userId, uint(toUserId), content)
		if err != nil {
			response.MessageActionResponseHandler(c, "1", "消息发送失败")
			return
		}
		// 信息发送成功
		response.MessageActionResponseHandler(c, "0", "消息发送成功")
		return
	}
	// 操作异常
	response.MessageActionResponseHandler(c, "1", "消息发送操作异常")
}

// MessageChat 聊天记录
// url: ./**&to_user_id=5&pre_msg_time=1695307902785
func MessageChat(c *gin.Context) {
	toUserId, err := strconv.ParseUint(c.Query("to_user_id"), 10, 64)
	if err != nil {
		response.ToUserIdConversionErr(c, "参数类型转换错误")
	}
	// 当前用户
	userId := c.GetUint("user_id")
	preMsgTime := c.Query("pre_msg_time") // 发送的时间戳
	msgTime := time.Now()
	//fmt.Printf("userId = %+v, toUserId = %+v, preMsgTime = %+v\n", userId, toUserId, preMsgTime)
	if preMsgTime != "" {
		preMsgTimeUnix, err := strconv.ParseInt(preMsgTime, 10, 64)
		if err != nil {
			response.ToUserIdConversionErr(c, "时间戳转换失败")
			return
		}
		msgTime = time.Unix(0, preMsgTimeUnix*int64(time.Millisecond))
	}
	// 获取信息列表
	messages, err := MessageService.GetMsgListWithTime(userId, uint(toUserId), msgTime)
	if err != nil {
		response.MessageChatResponseHandler(c, "1", "获取聊天记录失败", nil)
	}

	var msgsResp []response.Message
	for _, msg := range messages {
		msgsResp = append(msgsResp, response.Message{
			Content:    msg.Content,
			CreateTime: msg.CreateTime.Unix(),
			FromUserID: int64(msg.FromUserID),
			ID:         int64(msg.ID),
			ToUserID:   int64(msg.ToUserID),
		})
	}
	// 获取成功
	response.MessageChatResponseHandler(c, "0", "获取聊天记录成功", msgsResp)

}
