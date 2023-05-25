package interaction

import (
	"context"

	"github.com/EliasStar/BacoTell/internal/proto/providerpb"
	"github.com/EliasStar/BacoTell/pkg/provider"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type componentServer struct {
	providerpb.UnimplementedComponentServer

	impl   provider.Component
	broker *plugin.GRPCBroker
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
func (s componentServer) Handle(_ context.Context, req *providerpb.HandleRequest) (*emptypb.Empty, error) {
	conn, err := s.broker.Dial(req.HandleProxyId)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	return &emptypb.Empty{}, s.impl.Handle(handleProxyClient{
		interactionProxyClient: interactionProxyClient{
			client: providerpb.NewInteractionProxyClient(conn),
		},
		client: providerpb.NewHandleProxyClient(conn),
	})
}

type componentClient struct {
	client providerpb.ComponentClient
	broker *plugin.GRPCBroker
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
func (c componentClient) Handle(proxy provider.HandleProxy) error {
	var server *grpc.Server
	defer server.Stop()

	id := c.broker.NextId()
	go c.broker.AcceptAndServe(id, func(opts []grpc.ServerOption) *grpc.Server {
		server = grpc.NewServer(opts...)

		srv := handleProxyServer{
			interactionProxyServer: interactionProxyServer{
				impl: proxy,
			},
			impl: proxy,
		}

		providerpb.RegisterHandleProxyServer(server, srv)
		providerpb.RegisterInteractionProxyServer(server, srv)

		return server
	})

	_, err := c.client.Handle(context.Background(), &providerpb.HandleRequest{
		HandleProxyId: id,
	})

	return err
}
