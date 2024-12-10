package client

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/spv-dev/chat-client/internal/converter"
	"github.com/spv-dev/chat-client/internal/model"
	desc "github.com/spv-dev/chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50061"
)

var globalClient desc.ChatV1Client

func Init() {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}

	cl := desc.NewChatV1Client(conn)

	globalClient = cl
}

func CreateChat(ctx context.Context, title string) (*model.Chat, error) {
	res, err := globalClient.CreateChat(ctx, &desc.CreateChatRequest{
		Info: &desc.ChatInfo{
			Title: title,
		},
	})
	if err != nil {
		return &model.Chat{}, err
	}

	return converter.ToChatFromAPI(res.Chat), err
}

func ConnectChat(ctx context.Context, chatID string, userID string) (desc.ChatV1_ConnectChatClient, error) {
	cID, err := strconv.ParseInt(chatID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("can't convert chat id: %v", err)
	}

	uID, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("can't convert user id: %v", err)
	}

	res, err := globalClient.ConnectChat(ctx, &desc.ConnectChatRequest{
		ChatId: cID,
		UserId: uID,
	})
	if err != nil {
		return nil, err
	}

	return res, err
}

func SendMessage(ctx context.Context, chatID string, userID string, body string) (*model.Message, error) {
	cID, err := strconv.ParseInt(chatID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("can't convert chat id: %v", err)
	}

	uID, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("can't convert user id: %v", err)
	}
	res, err := globalClient.SendMessage(ctx, &desc.SendMessageRequest{
		Info: &desc.MessageInfo{
			ChatId: cID,
			UserId: uID,
			Body:   body,
		},
	})
	if err != nil {
		return nil, err
	}

	return converter.ToMessageFromAPI(res.Message), err
}
