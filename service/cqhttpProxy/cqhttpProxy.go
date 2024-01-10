package cqhttpProxy

import (
	"bytes"
	"chatgpt-im-bot/config"
	"encoding/json"
	"fmt"
	"net/http"
)

type GroupMessage struct {
	GroupId uint64 `json:"group_id"`
	Message string `json:"message"`
}

func (m *GroupMessage) Send() {
	url := config.GetConfig().QQbotConfig.ProxyURL + "send_group_msg"
	body, _ := json.Marshal(m)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		// 处理错误
		fmt.Println("Error:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// 处理错误
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Status:", resp.Status)
}
func (m *GroupMessage) SendImage(filePath string) {

}
