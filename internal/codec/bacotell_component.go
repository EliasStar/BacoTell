package codec

import (
	"context"

	"github.com/EliasStar/BacoTell/internal/proto/bacotellpb"
	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	"github.com/bwmarrin/discordgo"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type componentServer struct {
	bacotellpb.UnimplementedComponentServer

	impl   common.Component
	broker *plugin.GRPCBroker
}

var _ bacotellpb.ComponentServer = componentServer{}

// CustomId implements bacotellpb.ComponentServer
func (s componentServer) CustomId(context.Context, *bacotellpb.ComponentCustomIdRequest) (*bacotellpb.ComponentCustomIdResponse, error) {
	id, err := s.impl.CustomID()
	if err != nil {
		return nil, err
	}

	return &bacotellpb.ComponentCustomIdResponse{CustomId: id}, nil
}

// Handle implements bacotellpb.ComponentServer
func (s componentServer) Handle(_ context.Context, req *bacotellpb.ComponentHandleRequest) (*bacotellpb.ComponentHandleResponse, error) {
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
	return &bacotellpb.ComponentHandleResponse{}, err
}

type componentClient struct {
	client bacotellpb.ComponentClient
	broker *plugin.GRPCBroker
}

var _ common.Component = componentClient{}

// CustomID implements bacotell_common.Component
func (c componentClient) CustomID() (string, error) {
	res, err := c.client.CustomId(context.Background(), &bacotellpb.ComponentCustomIdRequest{})
	if err != nil {
		return "", err
	}

	return res.CustomId, nil
}

// Handle implements bacotell_common.Component
func (c componentClient) Handle(proxy common.HandleProxy) error {
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

	_, err := c.client.Handle(context.Background(), &bacotellpb.ComponentHandleRequest{HandleProxyId: id})

	server.Stop()
	return err
}

type handleProxyServer struct {
	bacotellpb.UnimplementedHandleProxyServer
	interactionProxyServer

	impl common.HandleProxy
}

var (
	_ bacotellpb.InteractionProxyServer = handleProxyServer{}
	_ bacotellpb.HandleProxyServer      = handleProxyServer{}
)

// ComponentType implements bacotellpb.HandleProxyServer
func (s handleProxyServer) ComponentType(context.Context, *bacotellpb.HandleProxyComponentTypeRequest) (*bacotellpb.HandleProxyComponentTypeResponse, error) {
	typ, err := s.impl.ComponentType()
	if err != nil {
		return nil, err
	}

	return &bacotellpb.HandleProxyComponentTypeResponse{Type: uint32(typ)}, nil
}

// SelectedValues implements bacotellpb.HandleProxyServer
func (s handleProxyServer) SelectedValues(context.Context, *bacotellpb.HandleProxySelectedValuesRequest) (*bacotellpb.HandleProxySelectedValuesResponse, error) {
	values, err := s.impl.SelectedValues()
	if err != nil {
		return nil, err
	}

	return &bacotellpb.HandleProxySelectedValuesResponse{Values: values}, nil
}

type handleProxyClient struct {
	interactionProxyClient

	client bacotellpb.HandleProxyClient
}

var (
	_ common.InteractionProxy = handleProxyClient{}
	_ common.HandleProxy      = handleProxyClient{}
)

// ComponentType implements bacotell_common.HandleProxy
func (c handleProxyClient) ComponentType() (discordgo.ComponentType, error) {
	res, err := c.client.ComponentType(context.Background(), &bacotellpb.HandleProxyComponentTypeRequest{})
	if err != nil {
		return 0, err
	}

	return discordgo.ComponentType(res.Type), nil
}

// SelectedValues implements bacotell_common.HandleProxy
func (c handleProxyClient) SelectedValues() ([]string, error) {
	res, err := c.client.SelectedValues(context.Background(), &bacotellpb.HandleProxySelectedValuesRequest{})
	if err != nil {
		return nil, err
	}

	return res.Values, nil
}
