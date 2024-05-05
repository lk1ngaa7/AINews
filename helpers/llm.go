package helpers

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

func GptGen(str string) (ret string, err error) {
	config := openai.DefaultConfig("token")
	config.BaseURL = "https://api.gptapi.us/v1"
	c := openai.NewClientWithConfig(config)
	ctx := context.Background()

	req := openai.CompletionRequest{
		Model:  openai.GPT3Dot5Turbo,
		Prompt: str,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		BuzzLogger.Warn(fmt.Sprintf("gpt gen Error: %v", err))
		return
	}
	// resp.Choices[0].Text 是否存在
	if len(resp.Choices) > 0 {
		ret = resp.Choices[0].Text
	}
	return
}
