package codec

import (
	"context"

	"github.com/EliasStar/BacoTell/internal/proto/bacotellpb"
	"github.com/EliasStar/BacoTell/pkg/bacotell"
	"github.com/bwmarrin/discordgo"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type commandServer struct {
	bacotellpb.UnimplementedCommandServer

	impl   bacotell.Command
	broker *plugin.GRPCBroker
}

var _ bacotellpb.CommandServer = commandServer{}

// CommandData implements bacotellpb.CommandServer
func (s commandServer) CommandData(context.Context, *bacotellpb.CommandDataRequest) (*bacotellpb.CommandDataResponse, error) {
	data, err := s.impl.CommandData()
	if err != nil {
		return nil, err
	}

	return &bacotellpb.CommandDataResponse{Data: encodeApplicationCommand(data)}, nil
}

// Execute implements bacotellpb.CommandServer
func (s commandServer) Execute(_ context.Context, req *bacotellpb.ExecuteRequest) (*bacotellpb.ExecuteResponse, error) {
	conn, err := s.broker.Dial(req.ExecuteProxyId)
	if err != nil {
		return nil, err
	}

	err = s.impl.Execute(executeProxyClient{
		interactionProxyClient: interactionProxyClient{
			client: bacotellpb.NewInteractionProxyClient(conn),
		},
		client: bacotellpb.NewExecuteProxyClient(conn),
	})

	conn.Close()
	return &bacotellpb.ExecuteResponse{}, err
}

type commandClient struct {
	client bacotellpb.CommandClient
	broker *plugin.GRPCBroker
}

var _ bacotell.Command = commandClient{}

// CommandData implements bacotell.Command
func (c commandClient) CommandData() (discordgo.ApplicationCommand, error) {
	res, err := c.client.CommandData(context.Background(), &bacotellpb.CommandDataRequest{})
	if err != nil {
		return discordgo.ApplicationCommand{}, err
	}

	return decodeApplicationCommand(res.Data), nil
}

// Execute implements bacotell.Command
func (c commandClient) Execute(proxy bacotell.ExecuteProxy) error {
	var server *grpc.Server

	id := c.broker.NextId()
	go c.broker.AcceptAndServe(id, func(opts []grpc.ServerOption) *grpc.Server {
		server = grpc.NewServer(opts...)

		srv := executeProxyServer{
			interactionProxyServer: interactionProxyServer{
				impl: proxy,
			},
			impl: proxy,
		}

		bacotellpb.RegisterExecuteProxyServer(server, srv)
		bacotellpb.RegisterInteractionProxyServer(server, srv)

		return server
	})

	_, err := c.client.Execute(context.Background(), &bacotellpb.ExecuteRequest{ExecuteProxyId: id})

	server.Stop()
	return err
}

type executeProxyServer struct {
	bacotellpb.UnimplementedExecuteProxyServer
	interactionProxyServer

	impl bacotell.ExecuteProxy
}

var (
	_ bacotellpb.InteractionProxyServer = executeProxyServer{}
	_ bacotellpb.ExecuteProxyServer     = executeProxyServer{}
)

// StringOption implements bacotellpb.ExecuteProxyServer
func (s executeProxyServer) StringOption(_ context.Context, req *bacotellpb.StringOptionRequest) (*bacotellpb.StringOptionResponse, error) {
	val, err := s.impl.StringOption(req.Name)
	if err != nil {
		return nil, err
	}

	return &bacotellpb.StringOptionResponse{Value: val}, nil
}

// IntegerOption implements bacotellpb.ExecuteProxyServer
func (s executeProxyServer) IntegerOption(_ context.Context, req *bacotellpb.IntegerOptionRequest) (*bacotellpb.IntegerOptionResponse, error) {
	val, err := s.impl.IntegerOption(req.Name)
	if err != nil {
		return nil, err
	}

	return &bacotellpb.IntegerOptionResponse{Value: val}, nil
}

// NumberOption implements bacotellpb.ExecuteProxyServer
func (s executeProxyServer) NumberOption(_ context.Context, req *bacotellpb.NumberOptionRequest) (*bacotellpb.NumberOptionResponse, error) {
	val, err := s.impl.NumberOption(req.Name)
	if err != nil {
		return nil, err
	}

	return &bacotellpb.NumberOptionResponse{Value: val}, nil
}

// BooleanOption implements bacotellpb.ExecuteProxyServer
func (s executeProxyServer) BooleanOption(_ context.Context, req *bacotellpb.BooleanOptionRequest) (*bacotellpb.BooleanOptionResponse, error) {
	val, err := s.impl.BooleanOption(req.Name)
	if err != nil {
		return nil, err
	}

	return &bacotellpb.BooleanOptionResponse{Value: val}, nil
}

// UserOption implements bacotellpb.ExecuteProxyServer
func (s executeProxyServer) UserOption(_ context.Context, req *bacotellpb.UserOptionRequest) (*bacotellpb.UserOptionResponse, error) {
	val, err := s.impl.UserOption(req.Name)
	if err != nil {
		return nil, err
	}

	return &bacotellpb.UserOptionResponse{Value: encodeUser(val)}, nil
}

// RoleOption implements bacotellpb.ExecuteProxyServer
func (s executeProxyServer) RoleOption(_ context.Context, req *bacotellpb.RoleOptionRequest) (*bacotellpb.RoleOptionResponse, error) {
	val, err := s.impl.RoleOption(req.Name)
	if err != nil {
		return nil, err
	}

	return &bacotellpb.RoleOptionResponse{Value: encodeRole(val)}, nil
}

// ChannelOption implements bacotellpb.ExecuteProxyServer
func (s executeProxyServer) ChannelOption(_ context.Context, req *bacotellpb.ChannelOptionRequest) (*bacotellpb.ChannelOptionResponse, error) {
	val, err := s.impl.ChannelOption(req.Name)
	if err != nil {
		return nil, err
	}

	return &bacotellpb.ChannelOptionResponse{Value: encodeChannel(val)}, nil
}

// AttachmentOption implements bacotellpb.ExecuteProxyServer
func (s executeProxyServer) AttachmentOption(_ context.Context, req *bacotellpb.AttachmentOptionRequest) (*bacotellpb.AttachmentOptionResponse, error) {
	val, err := s.impl.AttachmentOption(req.Name)
	if err != nil {
		return nil, err
	}

	return &bacotellpb.AttachmentOptionResponse{Value: encodeAttachment(val)}, nil
}

type executeProxyClient struct {
	interactionProxyClient

	client bacotellpb.ExecuteProxyClient
}

var (
	_ bacotell.InteractionProxy = executeProxyClient{}
	_ bacotell.ExecuteProxy     = executeProxyClient{}
)

// StringOption implements bacotell.ExecuteProxy
func (c executeProxyClient) StringOption(name string) (string, error) {
	res, err := c.client.StringOption(context.Background(), &bacotellpb.StringOptionRequest{
		Name: name,
	})

	if err != nil {
		return "", err
	}

	return res.Value, nil
}

// IntegerOption implements bacotell.ExecuteProxy
func (c executeProxyClient) IntegerOption(name string) (int64, error) {
	res, err := c.client.IntegerOption(context.Background(), &bacotellpb.IntegerOptionRequest{
		Name: name,
	})

	if err != nil {
		return 0, err
	}

	return res.Value, nil
}

// NumberOption implements bacotell.ExecuteProxy
func (c executeProxyClient) NumberOption(name string) (float64, error) {
	res, err := c.client.NumberOption(context.Background(), &bacotellpb.NumberOptionRequest{
		Name: name,
	})

	if err != nil {
		return 0, err
	}

	return res.Value, nil
}

// BooleanOption implements bacotell.ExecuteProxy
func (c executeProxyClient) BooleanOption(name string) (bool, error) {
	res, err := c.client.BooleanOption(context.Background(), &bacotellpb.BooleanOptionRequest{
		Name: name,
	})

	if err != nil {
		return false, err
	}

	return res.Value, nil
}

// UserOption implements bacotell.ExecuteProxy
func (c executeProxyClient) UserOption(name string) (*discordgo.User, error) {
	res, err := c.client.UserOption(context.Background(), &bacotellpb.UserOptionRequest{
		Name: name,
	})

	if err != nil {
		return nil, err
	}

	return decodeUser(res.Value), nil
}

// RoleOption implements bacotell.ExecuteProxy
func (c executeProxyClient) RoleOption(name string) (*discordgo.Role, error) {
	res, err := c.client.RoleOption(context.Background(), &bacotellpb.RoleOptionRequest{
		Name: name,
	})

	if err != nil {
		return nil, err
	}

	return decodeRole(res.Value), nil
}

// ChannelOption implements bacotell.ExecuteProxy
func (c executeProxyClient) ChannelOption(name string) (*discordgo.Channel, error) {
	res, err := c.client.ChannelOption(context.Background(), &bacotellpb.ChannelOptionRequest{
		Name: name,
	})

	if err != nil {
		return nil, err
	}

	return decodeChannel(res.Value), nil
}

// AttachmentOption implements bacotell.ExecuteProxy
func (c executeProxyClient) AttachmentOption(name string) (*discordgo.MessageAttachment, error) {
	res, err := c.client.AttachmentOption(context.Background(), &bacotellpb.AttachmentOptionRequest{
		Name: name,
	})

	if err != nil {
		return nil, err
	}

	return decodeAttachment(res.Value), nil
}
