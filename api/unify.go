package api

import (
	"chatgpt-im-bot/service/openaiProxy"
	"context"
)

type UnifyCommand struct {
	Command string
	Data    []byte
}

type UnifyResponse struct {
	Type string
	Data []byte
}

func (c *UnifyCommand) Exec(ctx context.Context) (UnifyResponse, error) {
	var rsp UnifyResponse
	if c.Command == "GPT" {
		msg := string(c.Data)
		chatGPTrsp, err := openaiProxy.GetChatGPTSimpleReply(ctx, msg)
		if err != nil {
			return rsp, err
		}
		rsp.Type = "MSG"
		rsp.Data = []byte(chatGPTrsp)
		return rsp, nil
	}
	return rsp, nil
}
