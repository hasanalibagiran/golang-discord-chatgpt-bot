package main

import (
	"context"
	gogpt "github.com/sashabaranov/go-gpt3"
)

var prompt = "The following is a conversation with an AI assistant. The assistant is helpful, creative, clever, and very friendly.\n\nHuman: Hello, who are you?\nAI: I am an AI created by OpenAI. How can I help you today?\n"

func gpt(input string) string {

	c := gogpt.NewClient("Api-Key")
	ctx := context.Background()

	prompt = prompt + "You: " + input + "\n"

	req := gogpt.CompletionRequest{
		Model:            "text-davinci-003",
		MaxTokens:        500,
		Temperature:      0.9,
		Prompt:           prompt,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0.5,
		Stop:             []string{" Human:", " AI:"},
	}

	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		panic(err)
	}

	prompt = prompt + resp.Choices[0].Text + "\n"

	return resp.Choices[0].Text[4:]

}
