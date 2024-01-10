package api

import (
	"chatgpt-im-bot/service/openaiProxy"
	"context"
	"errors"
	"strings"
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
	} else if c.Command == "DRAW" {
		chatGPTrsp := "D:\\project\\ChatGPT-IM-Bot\\sources\\generateImage\\imag1.jpg"
		rsp.Type = "IMAGE"
		rsp.Data = []byte(chatGPTrsp)
		return rsp, nil
	}

	return rsp, nil
}

func TrasnformToUnifyCommand(str string) (*UnifyCommand, error) {
	if strings.HasPrefix(str, "GPT") || strings.HasPrefix(str, "gpt") {
		if len(str) < 4 {
			return nil, errors.New("msg empty")
		}
		command := "GPT"
		msg := str[3:]
		return &UnifyCommand{
			Command: command,
			Data:    []byte(msg),
		}, nil
	}
	if strings.HasPrefix(str, " GPT") || strings.HasPrefix(str, " gpt") {
		if len(str) < 5 {
			return nil, errors.New("msg empty")
		}
		command := "GPT"
		msg := str[4:]
		return &UnifyCommand{
			Command: command,
			Data:    []byte(msg),
		}, nil
	}
	if strings.HasPrefix(str, "draw") || strings.HasPrefix(str, "DRAW") {
		if len(str) < 5 {
			return nil, errors.New("msg empty")
		}
		command := "DRAW"
		msg := str[4:]
		return &UnifyCommand{
			Command: command,
			Data:    []byte(msg),
		}, nil
	}
	return nil, errors.New("command not found")
}
