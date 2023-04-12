package config

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	QQbotConfig  QQbotConfig  `json:"qq_bot_config"`
	OpenAIConfig OpenAIConfig `json:"openai_config"`
}
type QQbotConfig struct {
	AllowQQGroup []uint64 `json:"allow_qq_group"`
	BotQQNumber  string   `json:"bot_qq_number"`
	ProxyURL     string   `json:"proxy_url"`
}
type OpenAIConfig struct {
	OpenAIKey string `json:"openai_key"`
}

var globalConfig *Config

func init() {

	jsonFile, err := os.Open("config/config.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	configByte, err := io.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}
	globalConfig = new(Config)
	err = json.Unmarshal(configByte, globalConfig)
	if err != nil {
		panic(err)
	}

}

func GetConfig() *Config {
	return globalConfig
}
