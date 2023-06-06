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
func (s pluginServer) Id(context.Context, *bacotellpb.PluginIdRequest) (*bacotellpb.PluginIdResponse, error) {
	id, err := s.impl.ID()
	if err != nil {
		return nil, err
	}

	return &bacotellpb.PluginIdResponse{Id: id}, nil
}

// ApplicationCommands implements bacotellpb.PluginServer
func (s pluginServer) ApplicationCommands(context.Context, *bacotellpb.PluginApplicationCommandsRequest) (*bacotellpb.PluginApplicationCommandsResponse, error) {
	commands, err := s.impl.ApplicationCommands()
	if err != nil {
		return nil, err
	}

	ids := make([]uint32, len(commands))
	for i, command := range commands {
		ids[i] = s.broker.NextId()
		srv := commandServer{impl: command, broker: s.broker}

		go s.broker.AcceptAndServe(ids[i], func(opts []grpc.ServerOption) *grpc.Server {
			server := grpc.NewServer(opts...)
			bacotellpb.RegisterCommandServer(server, srv)
			return server
		})
	}

	return &bacotellpb.PluginApplicationCommandsResponse{CommandIds: ids}, nil
}

// MessageComponents implements bacotellpb.PluginServer
func (s pluginServer) MessageComponents(context.Context, *bacotellpb.PluginMessageComponentsRequest) (*bacotellpb.PluginMessageComponentsResponse, error) {
	components, err := s.impl.MessageComponents()
	if err != nil {
		return nil, err
	}

	ids := make([]uint32, len(components))
	for i, component := range components {
		ids[i] = s.broker.NextId()
		srv := componentServer{impl: component, broker: s.broker}

		go s.broker.AcceptAndServe(ids[i], func(opts []grpc.ServerOption) *grpc.Server {
			server := grpc.NewServer(opts...)
			bacotellpb.RegisterComponentServer(server, srv)
			return server
		})
	}

	return &bacotellpb.PluginMessageComponentsResponse{ComponentIds: ids}, nil
}

// Modals implements bacotellpb.PluginServer.
func (s pluginServer) Modals(context.Context, *bacotellpb.PluginModalsRequest) (*bacotellpb.PluginModalsResponse, error) {
	modals, err := s.impl.Modals()
	if err != nil {
		return nil, err
	}

	ids := make([]uint32, len(modals))
	for i, modal := range modals {
		ids[i] = s.broker.NextId()
		srv := modalServer{impl: modal, broker: s.broker}

		go s.broker.AcceptAndServe(ids[i], func(opts []grpc.ServerOption) *grpc.Server {
			server := grpc.NewServer(opts...)
			bacotellpb.RegisterModalServer(server, srv)
			return server
		})
	}

	return &bacotellpb.PluginModalsResponse{ModalIds: ids}, nil
}

type pluginClient struct {
	client bacotellpb.PluginClient
	broker *plugin.GRPCBroker
}

var _ common.Plugin = pluginClient{}

// ID implements bacotell_common.Plugin
func (c pluginClient) ID() (string, error) {
	res, err := c.client.Id(context.Background(), &bacotellpb.PluginIdRequest{})
	if err != nil {
		return "", err
	}

	return res.Id, nil
}

// ApplicationCommands implements bacotell_common.Plugin
func (c pluginClient) ApplicationCommands() ([]common.Command, error) {
	res, err := c.client.ApplicationCommands(context.Background(), &bacotellpb.PluginApplicationCommandsRequest{})
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
	res, err := c.client.MessageComponents(context.Background(), &bacotellpb.PluginMessageComponentsRequest{})
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

// Modals implements bacotell_common.Plugin.
func (c pluginClient) Modals() ([]common.Modal, error) {
	res, err := c.client.Modals(context.Background(), &bacotellpb.PluginModalsRequest{})
	if err != nil {
		return nil, err
	}

	modals := make([]common.Modal, len(res.ModalIds))
	for i, id := range res.ModalIds {
		conn, err := c.broker.Dial(id)
		if err != nil {
			return nil, err
		}

		modals[i] = &modalClient{client: bacotellpb.NewModalClient(conn), broker: c.broker}
	}

	return modals, nil
}
