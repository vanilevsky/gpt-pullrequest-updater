package openai

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

//go:embed prompts/review
var PromptReview string

const (
	PromptDescribeChanges = "Below is the code patch, Generate a GitHub pull request description based on the following comments without basic prefix\n%s\n"
	PromptOverallDescribe = "Below comments are generated by AI, Generate a GitHub pull request description based on the following comments without basic prefix in markdown format with ### Description and ### Changes blocks:\n%s\n"
)

type Client struct {
	client *openai.Client
}

func NewClient(token string) *Client {
	return &Client{
		client: openai.NewClient(token),
	}
}

func (o *Client) ChatCompletion(ctx context.Context, messages []openai.ChatCompletionMessage) (string, error) {
	resp, err := o.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:       openai.GPT3Dot5Turbo,
			Messages:    messages,
			Temperature: 0.1,
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
