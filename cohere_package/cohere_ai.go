package coherepackage

import (
	"context"

	cohere "github.com/cohere-ai/cohere-go/v2"
	client "github.com/cohere-ai/cohere-go/v2/client"
)

const api_token string = "Ql450UOM2vGInI88hcxcjTuyZu9E9aCbowh0UDrO"

func GetAiResponse(message string, chat_history []*cohere.Message) (string, []*cohere.Message) {
	client := client.NewClient(client.WithToken(api_token))
	response, err := client.Chat(
		context.TODO(),
		&cohere.ChatRequest{
			ChatHistory: chat_history,
			Message:     message,
		},
	)
	if err != nil {
		return err.Error(), nil
	}

	return response.Text, response.ChatHistory
}
