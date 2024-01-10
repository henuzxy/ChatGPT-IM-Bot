package qqbot

import (
	"chatgpt-im-bot/api"
	"chatgpt-im-bot/config"
	"chatgpt-im-bot/logger"
	"chatgpt-im-bot/utils"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func ReceiveMessage(c *gin.Context) {
	api.ResponseSuccess(c)
	var msg QQGroupMessage
	err := c.ShouldBind(&msg)
	if err != nil {
		logger.Errorf("QQbot RecieveMessage ShouldBind error:%s", err)
		return
	}
	logger.Info(fmt.Sprintf("%+v", msg))
	str, ok := AuthMsg(&msg)
	if !ok {
		logger.Errorf("QQbot AuthMsg error:%s", err)
		return
	}

	ctx := context.Background()
	unifyCommand, err := api.TrasnformToUnifyCommand(str)
	if err != nil {
		msg.Reply(err.Error())
		return
	}
	unifyResponse, err := unifyCommand.Exec(ctx)
	if err != nil {
		msg.Reply(err.Error())
		return
	}

	if unifyResponse.Type == "MSG" {
		replyMsg := string(unifyResponse.Data)
		msg.Reply(replyMsg)
	} else if unifyResponse.Type == "IMAGE" {
		filePath := string(unifyResponse.Data)
		msg.ReplyImage(filePath)
	}
	return
}
func isDirectTargetToBot(msg *QQGroupMessage) bool {
	if len(msg.Message) < 2 {
		return false
	}
	if msg.Message[0].Type != "at" {
		return false
	}
	if msg.Message[1].Type != "text" {
		return false
	}
	targetQQ, _ := msg.Message[0].Data["qq"]
	if targetQQ != config.GetConfig().QQbotConfig.BotQQNumber {
		return false
	}
	return true
}

// isNonDirectTargetToBot 没有使用@功能
func isNonDirectTargetToBot(msg *QQGroupMessage) bool {
	if len(msg.Message) != 1 {
		return false
	}
	if msg.Message[0].Type != "text" {
		return false
	}
	str := msg.Message[0].Data["text"]
	if strings.HasPrefix(str, "gpt") || strings.HasPrefix(str, "GPT") {
		return true
	}
	if strings.HasPrefix(str, "draw") || strings.HasPrefix(str, "DRAW") {
		return true
	}
	return true
}
func isTargetToBot(msg *QQGroupMessage) (string, bool) {
	if isDirectTargetToBot(msg) {
		return msg.Message[1].Data["text"], true
	}
	if isNonDirectTargetToBot(msg) {
		return msg.Message[0].Data["text"], true
	}
	return "", false
}
func AuthMsg(msg *QQGroupMessage) (string, bool) {
	if msg.MessageType != "group" {
		return "", false
	}
	if !utils.IsCointainUint64(config.GetConfig().QQbotConfig.AllowQQGroup, uint64(msg.GroupID)) {
		return "", false
	}
	return isTargetToBot(msg)
}
