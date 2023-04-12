package router

import (
	"chatgpt-im-bot/api/qqbot"
	"chatgpt-im-bot/logger"
	"github.com/gin-gonic/gin"
)

const (
	port          = ":10101"
	smartbotGroup = "smartbot"
)

func RunServer() {
	engine := gin.Default()

	router := engine.Group(smartbotGroup)
	router.POST("receiveMessage", qqbot.ReceiveMessage)
	logger.Info("server running")
	engine.Run(port)
}
