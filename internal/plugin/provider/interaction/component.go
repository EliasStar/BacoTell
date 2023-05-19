package interaction

import (
	"context"

	"github.com/EliasStar/BacoTell/internal/proto/providerpb"
	"github.com/EliasStar/BacoTell/pkg/provider"
	"google.golang.org/protobuf/types/known/emptypb"
)

type componentServer struct {
	providerpb.UnimplementedComponentServer

	impl provider.Component
}

var _ providerpb.ComponentServer = componentServer{}

// CustomId implements providerpb.ComponentServer
func (s componentServer) CustomId(context.Context, *emptypb.Empty) (*providerpb.CustomIdResponse, error) {
	id, err := s.impl.CustomId()
	if err != nil {
		return nil, err
	}

	return &providerpb.CustomIdResponse{CustomId: id}, nil
}

// Handle implements providerpb.ComponentServer
func (s componentServer) Handle(context.Context, *providerpb.HandleRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.impl.Handle(nil)
}

type componentClient struct {
	client providerpb.ComponentClient
}

var _ provider.Component = componentClient{}

// CustomId implements provider.Component
func (c componentClient) CustomId() (string, error) {
	res, err := c.client.CustomId(context.Background(), &emptypb.Empty{})
	if err != nil {
		return "", err
	}

	return res.CustomId, nil
}

// Handle implements provider.Component
func (c componentClient) Handle(provider.InteractionProxy) error {
	_, err := c.client.Handle(context.Background(), &providerpb.HandleRequest{})
	return err
}
