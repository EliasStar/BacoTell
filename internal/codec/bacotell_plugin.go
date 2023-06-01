package codec

import (
	"context"

	"github.com/EliasStar/BacoTell/internal/proto/bacotellpb"
	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type pluginServer struct {
	bacotellpb.UnimplementedPluginServer

	impl   common.Plugin
	broker *plugin.GRPCBroker
}

var _ bacotellpb.PluginServer = pluginServer{}

// Id implements bacotellpb.PluginServer
func (s pluginServer) Id(context.Context, *bacotellpb.IdRequest) (*bacotellpb.IdResponse, error) {
	id, err := s.impl.ID()
	if err != nil {
		return nil, err
	}

	return &bacotellpb.IdResponse{Id: id}, nil
}

// ApplicationCommands implements bacotellpb.PluginServer
func (s pluginServer) ApplicationCommands(context.Context, *bacotellpb.ApplicationCommandsRequest) (*bacotellpb.ApplicationCommandsResponse, error) {
	commands, err := s.impl.ApplicationCommands()
	if err != nil {
		return nil, err
	}

	ids := make([]uint32, len(commands))
	for i, command := range commands {
		ids[i] = s.broker.NextId()
		go s.broker.AcceptAndServe(ids[i], func(opts []grpc.ServerOption) *grpc.Server {
			server := grpc.NewServer(opts...)
			bacotellpb.RegisterCommandServer(server, commandServer{impl: command, broker: s.broker})
			return server
		})
	}

	return &bacotellpb.ApplicationCommandsResponse{CommandIds: ids}, nil
}

// MessageComponents implements bacotellpb.PluginServer
func (s pluginServer) MessageComponents(context.Context, *bacotellpb.MessageComponentsRequest) (*bacotellpb.MessageComponentsResponse, error) {
	components, err := s.impl.MessageComponents()
	if err != nil {
		return nil, err
	}

	ids := make([]uint32, len(components))
	for i, component := range components {
		ids[i] = s.broker.NextId()
		go s.broker.AcceptAndServe(ids[i], func(opts []grpc.ServerOption) *grpc.Server {
			server := grpc.NewServer(opts...)
			bacotellpb.RegisterComponentServer(server, componentServer{impl: component, broker: s.broker})
			return server
		})
	}

	return &bacotellpb.MessageComponentsResponse{ComponentIds: ids}, nil
}

type pluginClient struct {
	client bacotellpb.PluginClient
	broker *plugin.GRPCBroker
}

var _ common.Plugin = pluginClient{}

// ID implements bacotell_common.Plugin
func (c pluginClient) ID() (string, error) {
	res, err := c.client.Id(context.Background(), &bacotellpb.IdRequest{})
	if err != nil {
		return "", err
	}

	return res.Id, nil
}

// ApplicationCommands implements bacotell_common.Plugin
func (c pluginClient) ApplicationCommands() ([]common.Command, error) {
	res, err := c.client.ApplicationCommands(context.Background(), &bacotellpb.ApplicationCommandsRequest{})
	if err != nil {
		return nil, err
	}

	commands := make([]common.Command, len(res.CommandIds))
	for i, id := range res.CommandIds {
		conn, err := c.broker.Dial(id)
		if err != nil {
			return nil, err
		}

		commands[i] = &commandClient{client: bacotellpb.NewCommandClient(conn), broker: c.broker}
	}

	return commands, nil
}

// MessageComponents implements bacotell_common.Plugin
func (c pluginClient) MessageComponents() ([]common.Component, error) {
	res, err := c.client.MessageComponents(context.Background(), &bacotellpb.MessageComponentsRequest{})
	if err != nil {
		return nil, err
	}

	components := make([]common.Component, len(res.ComponentIds))
	for i, id := range res.ComponentIds {
		conn, err := c.broker.Dial(id)
		if err != nil {
			return nil, err
		}

		components[i] = &componentClient{client: bacotellpb.NewComponentClient(conn), broker: c.broker}
	}

	return components, nil
}
