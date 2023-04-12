package qqbot

import (
	"chatgpt-im-bot/service/cqhttpProxy"
	"fmt"
)

func (qqGmsg *QQGroupMessage) Reply(replyMsg string) {

	message := fmt.Sprintf("[CQ:reply,id=%d]%s", qqGmsg.MessageID, replyMsg)
	var CQmsg cqhttpProxy.GroupMessage
	CQmsg.GroupId = uint64(qqGmsg.GroupID)
	CQmsg.Message = message
	CQmsg.Send()
}
