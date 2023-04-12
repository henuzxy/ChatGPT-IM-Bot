package main

import (
	"chatgpt-im-bot/router"
	"chatgpt-im-bot/service/openaiProxy"
)

func main() {
	openaiProxy.RegisterOpenAI()
	router.RunServer()
	//TestServer()
}

func TestServer() {
	//fmt.Printf("%+v\n", config.GetConfig())
	//openaiProxy.RegisterOpenAI()
	//ctx := context.Background()
	//msg, err := openaiProxy.GetChatGPTSimpleReply(ctx, "请问你如何看待GPT4的推出？")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(msg)
}
