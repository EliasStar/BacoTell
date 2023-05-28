package codec

import (
	"context"

	"github.com/EliasStar/BacoTell/internal/proto/bacotellpb"
	"github.com/EliasStar/BacoTell/pkg/bacotell"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type componentServer struct {
	bacotellpb.UnimplementedComponentServer

	impl   bacotell.Component
	broker *plugin.GRPCBroker
}

var _ bacotellpb.ComponentServer = componentServer{}

// CustomId implements bacotellpb.ComponentServer
func (s componentServer) CustomId(context.Context, *bacotellpb.CustomIdRequest) (*bacotellpb.CustomIdResponse, error) {
	id, err := s.impl.CustomID()
	if err != nil {
		return nil, err
	}

	return &bacotellpb.CustomIdResponse{CustomId: id}, nil
}

// Handle implements bacotellpb.ComponentServer
func (s componentServer) Handle(_ context.Context, req *bacotellpb.HandleRequest) (*bacotellpb.HandleResponse, error) {
	conn, err := s.broker.Dial(req.HandleProxyId)
	if err != nil {
		return nil, err
	}

	err = s.impl.Handle(handleProxyClient{
		interactionProxyClient: interactionProxyClient{
			client: bacotellpb.NewInteractionProxyClient(conn),
		},
		client: bacotellpb.NewHandleProxyClient(conn),
	})

	conn.Close()
	return &bacotellpb.HandleResponse{}, err
}

type componentClient struct {
	client bacotellpb.ComponentClient
	broker *plugin.GRPCBroker
}

var _ bacotell.Component = componentClient{}

// CustomID implements bacotell.Component
func (c componentClient) CustomID() (string, error) {
	res, err := c.client.CustomId(context.Background(), &bacotellpb.CustomIdRequest{})
	if err != nil {
		return "", err
	}

	return res.CustomId, nil
}

// Handle implements bacotell.Component
func (c componentClient) Handle(proxy bacotell.HandleProxy) error {
	var server *grpc.Server

	id := c.broker.NextId()
	go c.broker.AcceptAndServe(id, func(opts []grpc.ServerOption) *grpc.Server {
		server = grpc.NewServer(opts...)

		srv := handleProxyServer{
			interactionProxyServer: interactionProxyServer{
				impl: proxy,
			},
			impl: proxy,
		}

		bacotellpb.RegisterHandleProxyServer(server, srv)
		bacotellpb.RegisterInteractionProxyServer(server, srv)

		return server
	})

	_, err := c.client.Handle(context.Background(), &bacotellpb.HandleRequest{HandleProxyId: id})

	server.Stop()
	return err
}

type handleProxyServer struct {
	bacotellpb.UnimplementedHandleProxyServer
	interactionProxyServer

	impl bacotell.HandleProxy
}

var (
	_ bacotellpb.InteractionProxyServer = handleProxyServer{}
	_ bacotellpb.HandleProxyServer      = handleProxyServer{}
)

type handleProxyClient struct {
	interactionProxyClient

	client bacotellpb.HandleProxyClient
}

var (
	_ bacotell.InteractionProxy = handleProxyClient{}
	_ bacotell.HandleProxy      = handleProxyClient{}
)
