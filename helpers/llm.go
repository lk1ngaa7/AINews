package helpers

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

func getLLMChat(msgList []openai.ChatCompletionMessage, modelName, token, baseUrl string) (ret string, err error) {
	config := openai.DefaultConfig(token)
	config.BaseURL = baseUrl
	c := openai.NewClientWithConfig(config)
	ctx := context.Background()
	req := openai.ChatCompletionRequest{
		Model:    modelName,
		Messages: msgList,
	}
	resp, err := c.CreateChatCompletion(ctx, req)
	if err != nil {
		BuzzLogger.Warn(fmt.Sprintf("gpt gen Error: %v", err))
		return
	}
	if len(resp.Choices) > 0 {
		ret = resp.Choices[0].Message.Content
	} else {
		return "", fmt.Errorf("no response")
	}
	return
}

func LLMDeepSeek(queStr, systemPrompt string) (ret string, err error) {
	token := "sk-10798117588b48979af4e41cbc89a692"
	baseUrl := "https://api.deepseek.com/v1"
	modelName := "deepseek-chat"
	msgList := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: systemPrompt,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: queStr,
		},
	}
	ret, err = getLLMChat(msgList, modelName, token, baseUrl)
	return
}
