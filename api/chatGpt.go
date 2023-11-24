package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"log"
)

const openaiApiUrl = "https://api.openai.com/v1/engines/text-davinci-003/completions"

func LinkToAi(c *gin.Context) {
	// 设置本地代理服务器地址
	proxyAddress := "http://127.0.0.1:10809"

	// 设置 OpenAI API Key
	apiKey := "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"

	// 创建 resty 客户端，并设置代理
	client := resty.New().
		SetProxy(proxyAddress)

	// 设置 OpenAI API Key
	client.SetHeader("Authorization", "Bearer "+apiKey)

	// 构建 API 请求
	payload := map[string]interface{}{
		"prompt":            "请帮我生成6个随机的常用词", //用于提示模型生成文本的起始内容。可以提供一个或多个句子作为提示
		"max_tokens":        100,             //限制生成文本的最大令牌数。令牌是模型处理文本的基本单位。通过控制令牌数量，可以控制生成文本的长度。
		"temperature":       1.7,             //控制生成文本的多样性。较高的温度值（大于1.0）会使输出更随机和多样化，而较低的温度值（小于1.0）会使输出更加确定和一致。
		"stop":              nil,             //指定一个或多个字符串作为停止标记，当模型生成包含停止标记的文本时，会停止生成并返回结果。
		"n":                 1,               //指定要生成的候选响应数量。模型会生成多个可能的响应，并根据其概率进行排序。默认情况下，`n` 的值为1，即返回最高概率的响应。
		"frequency_penalty": 0,
		"presence_penalty":  0,
	}

	// 发送请求
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post(openaiApiUrl)

	if err != nil {
		log.Fatal("Error sending request:", err)
	}

	// 处理 API 响应
	fmt.Println("Response Status:", resp.Status())
	fmt.Println("Response Body:", resp.String())

	// 检查 API 是否成功
	if resp.IsSuccess() {
		// 在这里处理 API 成功的情况
	} else {
		// 在这里处理 API 失败的情况
		log.Fatal("API request failed:", resp.Status())
	}
}
