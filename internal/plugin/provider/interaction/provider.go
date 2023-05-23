package interaction

import (
	"context"

	"github.com/EliasStar/BacoTell/internal/proto/providerpb"
	"github.com/EliasStar/BacoTell/pkg/provider"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type interactionProviderServer struct {
	providerpb.UnimplementedInteractionProviderServer

	impl   provider.InteractionProvider
	broker *plugin.GRPCBroker
}

var _ providerpb.InteractionProviderServer = interactionProviderServer{}

// GetPrefix implements providerpb.InteractionProviderServer
func (s interactionProviderServer) Prefix(context.Context, *emptypb.Empty) (*providerpb.GetPrefixResponse, error) {
	prefix, err := s.impl.Prefix()
	if err != nil {
		return nil, err
	}

	return &providerpb.GetPrefixResponse{Prefix: prefix}, nil
}

// GetApplicationCommands implements providerpb.InteractionProviderServer
func (s interactionProviderServer) ApplicationCommands(context.Context, *emptypb.Empty) (*providerpb.GetApplicationCommandsResponse, error) {
	commands, err := s.impl.ApplicationCommands()
	if err != nil {
		return nil, err
	}

	ids := make([]uint32, len(commands))
	for i, command := range commands {
		ids[i] = s.broker.NextId()
		go s.broker.AcceptAndServe(ids[i], func(opts []grpc.ServerOption) *grpc.Server {
			server := grpc.NewServer(opts...)
			providerpb.RegisterCommandServer(server, commandServer{impl: command, broker: s.broker})
			return server
		})
	}

	return &providerpb.GetApplicationCommandsResponse{CommandIds: ids}, nil
}

// GetMessageComponents implements providerpb.InteractionProviderServer
func (s interactionProviderServer) MessageComponents(context.Context, *emptypb.Empty) (*providerpb.GetMessageComponentsResponse, error) {
	components, err := s.impl.MessageComponents()
	if err != nil {
		return nil, err
	}

	ids := make([]uint32, len(components))
	for i, component := range components {
		ids[i] = s.broker.NextId()
		go s.broker.AcceptAndServe(ids[i], func(opts []grpc.ServerOption) *grpc.Server {
			server := grpc.NewServer(opts...)
			providerpb.RegisterComponentServer(server, componentServer{impl: component, broker: s.broker})
			return server
		})
	}

	return &providerpb.GetMessageComponentsResponse{ComponentIds: ids}, nil
}

type interactionProviderClient struct {
	client providerpb.InteractionProviderClient
	broker *plugin.GRPCBroker
}

var _ provider.InteractionProvider = interactionProviderClient{}

// GetPrefix implements provider.InteractionProvider
func (c interactionProviderClient) Prefix() (string, error) {
	res, err := c.client.Prefix(context.Background(), &emptypb.Empty{})
	if err != nil {
		return "", err
	}

	return res.Prefix, nil
}

// GetApplicationCommands implements provider.InteractionProvider
func (c interactionProviderClient) ApplicationCommands() ([]provider.Command, error) {
	res, err := c.client.ApplicationCommands(context.Background(), &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	commands := make([]provider.Command, len(res.CommandIds))
	for i, id := range res.CommandIds {
		conn, err := c.broker.Dial(id)
		if err != nil {
			return nil, err
		}

		commands[i] = &commandClient{client: providerpb.NewCommandClient(conn), broker: c.broker}
	}

	return commands, nil
}

// GetMessageComponents implements provider.InteractionProvider
func (c interactionProviderClient) MessageComponents() ([]provider.Component, error) {
	res, err := c.client.MessageComponents(context.Background(), &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	components := make([]provider.Component, len(res.ComponentIds))
	for i, id := range res.ComponentIds {
		conn, err := c.broker.Dial(id)
		if err != nil {
			return nil, err
		}

		components[i] = &componentClient{client: providerpb.NewComponentClient(conn), broker: c.broker}
	}

	return components, nil
}
