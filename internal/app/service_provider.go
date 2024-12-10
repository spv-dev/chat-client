package app

import (
	"google.golang.org/grpc"

	desc "github.com/spv-dev/chat-server/pkg/chat_v1"
)

type serviceProvider struct {
	conn   *grpc.ClientConn
	client desc.ChatV1Client
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}
