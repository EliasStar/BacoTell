package codec

import (
	"context"

	"github.com/EliasStar/BacoTell/internal/proto/bacotellpb"
	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	"github.com/bwmarrin/discordgo"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/structpb"
)

type commandServer struct {
	bacotellpb.UnimplementedCommandServer

	impl   common.Command
	broker *plugin.GRPCBroker
}

var _ bacotellpb.CommandServer = commandServer{}

// Data implements bacotellpb.CommandServer
func (s commandServer) Data(context.Context, *bacotellpb.CommandDataRequest) (*bacotellpb.CommandDataResponse, error) {
	data, err := s.impl.Data()
	if err != nil {
		return nil, err
	}

	return &bacotellpb.CommandDataResponse{Data: encodeApplicationCommand(&data)}, nil
}

// Execute implements bacotellpb.CommandServer
func (s commandServer) Execute(_ context.Context, req *bacotellpb.CommandExecuteRequest) (*bacotellpb.CommandExecuteResponse, error) {
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
	return &bacotellpb.CommandExecuteResponse{}, err
}

// Autocomplete implements bacotellpb.CommandServer
func (s commandServer) Autocomplete(_ context.Context, req *bacotellpb.CommandAutocompleteRequest) (*bacotellpb.CommandAutocompleteResponse, error) {
	conn, err := s.broker.Dial(req.AutocompleteProxyId)
	if err != nil {
		return nil, err
	}

	err = s.impl.Autocomplete(autocompleteProxyClient{client: bacotellpb.NewAutocompleteProxyClient(conn)})

	conn.Close()
	return &bacotellpb.CommandAutocompleteResponse{}, err
}

type commandClient struct {
	client bacotellpb.CommandClient
	broker *plugin.GRPCBroker
}

var _ common.Command = commandClient{}

// Data implements bacotell_common.Command
func (c commandClient) Data() (discordgo.ApplicationCommand, error) {
	res, err := c.client.Data(context.Background(), &bacotellpb.CommandDataRequest{})
	if err != nil {
		return discordgo.ApplicationCommand{}, err
	}

	return *decodeApplicationCommand(res.Data), nil
}

// Execute implements bacotell_common.Command
func (c commandClient) Execute(proxy common.ExecuteProxy) error {
	var server *grpc.Server

	id := c.broker.NextId()
	srv := executeProxyServer{
		interactionProxyServer: interactionProxyServer{
			impl: proxy,
		},
		impl: proxy,
	}

	go c.broker.AcceptAndServe(id, func(opts []grpc.ServerOption) *grpc.Server {
		server = grpc.NewServer(opts...)
		bacotellpb.RegisterExecuteProxyServer(server, srv)
		bacotellpb.RegisterInteractionProxyServer(server, srv)
		return server
	})

	_, err := c.client.Execute(context.Background(), &bacotellpb.CommandExecuteRequest{ExecuteProxyId: id})

	server.Stop()
	return err
}

// Autocomplete implements bacotell_common.Command
func (c commandClient) Autocomplete(proxy common.AutocompleteProxy) error {
	var server *grpc.Server

	id := c.broker.NextId()
	srv := autocompleteProxyServer{impl: proxy}

	go c.broker.AcceptAndServe(id, func(opts []grpc.ServerOption) *grpc.Server {
		server = grpc.NewServer(opts...)
		bacotellpb.RegisterAutocompleteProxyServer(server, srv)
		return server
	})

	_, err := c.client.Autocomplete(context.Background(), &bacotellpb.CommandAutocompleteRequest{AutocompleteProxyId: id})

	server.Stop()
	return err
}

type executeProxyServer struct {
	bacotellpb.UnimplementedExecuteProxyServer
	interactionProxyServer

	impl common.ExecuteProxy
}

var (
	_ bacotellpb.InteractionProxyServer = executeProxyServer{}
	_ bacotellpb.ExecuteProxyServer     = executeProxyServer{}
)

// StringOption implements bacotellpb.ExecuteProxyServer
func (s executeProxyServer) StringOption(_ context.Context, req *bacotellpb.ExecuteProxyStringOptionRequest) (*bacotellpb.ExecuteProxyStringOptionResponse, error) {
	val, err := s.impl.StringOption(req.Name)
	if err != nil {
		return nil, err
	}

	return &bacotellpb.ExecuteProxyStringOptionResponse{Value: val}, nil
}

// IntegerOption implements bacotellpb.ExecuteProxyServer
func (s executeProxyServer) IntegerOption(_ context.Context, req *bacotellpb.ExecuteProxyIntegerOptionRequest) (*bacotellpb.ExecuteProxyIntegerOptionResponse, error) {
	val, err := s.impl.IntegerOption(req.Name)
	if err != nil {
		return nil, err
	}

	return &bacotellpb.ExecuteProxyIntegerOptionResponse{Value: val}, nil
}

// NumberOption implements bacotellpb.ExecuteProxyServer
func (s executeProxyServer) NumberOption(_ context.Context, req *bacotellpb.ExecuteProxyNumberOptionRequest) (*bacotellpb.ExecuteProxyNumberOptionResponse, error) {
	val, err := s.impl.NumberOption(req.Name)
	if err != nil {
		return nil, err
	}

	return &bacotellpb.ExecuteProxyNumberOptionResponse{Value: val}, nil
}

// BooleanOption implements bacotellpb.ExecuteProxyServer
func (s executeProxyServer) BooleanOption(_ context.Context, req *bacotellpb.ExecuteProxyBooleanOptionRequest) (*bacotellpb.ExecuteProxyBooleanOptionResponse, error) {
	val, err := s.impl.BooleanOption(req.Name)
	if err != nil {
		return nil, err
	}

	return &bacotellpb.ExecuteProxyBooleanOptionResponse{Value: val}, nil
}

// UserOption implements bacotellpb.ExecuteProxyServer
func (s executeProxyServer) UserOption(_ context.Context, req *bacotellpb.ExecuteProxyUserOptionRequest) (*bacotellpb.ExecuteProxyUserOptionResponse, error) {
	val, err := s.impl.UserOption(req.Name)
	if err != nil {
		return nil, err
	}

	return &bacotellpb.ExecuteProxyUserOptionResponse{Value: encodeUser(val)}, nil
}

// RoleOption implements bacotellpb.ExecuteProxyServer
func (s executeProxyServer) RoleOption(_ context.Context, req *bacotellpb.ExecuteProxyRoleOptionRequest) (*bacotellpb.ExecuteProxyRoleOptionResponse, error) {
	val, err := s.impl.RoleOption(req.Name)
	if err != nil {
		return nil, err
	}

	return &bacotellpb.ExecuteProxyRoleOptionResponse{Value: encodeRole(val)}, nil
}

// ChannelOption implements bacotellpb.ExecuteProxyServer
func (s executeProxyServer) ChannelOption(_ context.Context, req *bacotellpb.ExecuteProxyChannelOptionRequest) (*bacotellpb.ExecuteProxyChannelOptionResponse, error) {
	val, err := s.impl.ChannelOption(req.Name)
	if err != nil {
		return nil, err
	}

	return &bacotellpb.ExecuteProxyChannelOptionResponse{Value: encodeChannel(val)}, nil
}

// AttachmentOption implements bacotellpb.ExecuteProxyServer
func (s executeProxyServer) AttachmentOption(_ context.Context, req *bacotellpb.ExecuteProxyAttachmentOptionRequest) (*bacotellpb.ExecuteProxyAttachmentOptionResponse, error) {
	val, err := s.impl.AttachmentOption(req.Name)
	if err != nil {
		return nil, err
	}

	return &bacotellpb.ExecuteProxyAttachmentOptionResponse{Value: encodeMessageAttachment(val)}, nil
}

type executeProxyClient struct {
	interactionProxyClient

	client bacotellpb.ExecuteProxyClient
}

var (
	_ common.InteractionProxy = executeProxyClient{}
	_ common.ExecuteProxy     = executeProxyClient{}
)

// StringOption implements bacotell_common.ExecuteProxy
func (c executeProxyClient) StringOption(name string) (string, error) {
	res, err := c.client.StringOption(context.Background(), &bacotellpb.ExecuteProxyStringOptionRequest{
		Name: name,
	})

	if err != nil {
		return "", err
	}

	return res.Value, nil
}

// IntegerOption implements bacotell_common.ExecuteProxy
func (c executeProxyClient) IntegerOption(name string) (int64, error) {
	res, err := c.client.IntegerOption(context.Background(), &bacotellpb.ExecuteProxyIntegerOptionRequest{
		Name: name,
	})

	if err != nil {
		return 0, err
	}

	return res.Value, nil
}

// NumberOption implements bacotell_common.ExecuteProxy
func (c executeProxyClient) NumberOption(name string) (float64, error) {
	res, err := c.client.NumberOption(context.Background(), &bacotellpb.ExecuteProxyNumberOptionRequest{
		Name: name,
	})

	if err != nil {
		return 0, err
	}

	return res.Value, nil
}

// BooleanOption implements bacotell_common.ExecuteProxy
func (c executeProxyClient) BooleanOption(name string) (bool, error) {
	res, err := c.client.BooleanOption(context.Background(), &bacotellpb.ExecuteProxyBooleanOptionRequest{
		Name: name,
	})

	if err != nil {
		return false, err
	}

	return res.Value, nil
}

// UserOption implements bacotell_common.ExecuteProxy
func (c executeProxyClient) UserOption(name string) (*discordgo.User, error) {
	res, err := c.client.UserOption(context.Background(), &bacotellpb.ExecuteProxyUserOptionRequest{
		Name: name,
	})

	if err != nil {
		return nil, err
	}

	return decodeUser(res.Value), nil
}

// RoleOption implements bacotell_common.ExecuteProxy
func (c executeProxyClient) RoleOption(name string) (*discordgo.Role, error) {
	res, err := c.client.RoleOption(context.Background(), &bacotellpb.ExecuteProxyRoleOptionRequest{
		Name: name,
	})

	if err != nil {
		return nil, err
	}

	return decodeRole(res.Value), nil
}

// ChannelOption implements bacotell_common.ExecuteProxy
func (c executeProxyClient) ChannelOption(name string) (*discordgo.Channel, error) {
	res, err := c.client.ChannelOption(context.Background(), &bacotellpb.ExecuteProxyChannelOptionRequest{
		Name: name,
	})

	if err != nil {
		return nil, err
	}

	return decodeChannel(res.Value), nil
}

// AttachmentOption implements bacotell_common.ExecuteProxy
func (c executeProxyClient) AttachmentOption(name string) (*discordgo.MessageAttachment, error) {
	res, err := c.client.AttachmentOption(context.Background(), &bacotellpb.ExecuteProxyAttachmentOptionRequest{
		Name: name,
	})

	if err != nil {
		return nil, err
	}

	return decodeMessageAttachment(res.Value), nil
}

type autocompleteProxyServer struct {
	bacotellpb.UnimplementedAutocompleteProxyServer

	impl common.AutocompleteProxy
}

var _ bacotellpb.AutocompleteProxyServer = autocompleteProxyServer{}

// Respond implements bacotellpb.AutocompleteProxyServer
func (s autocompleteProxyServer) Respond(_ context.Context, req *bacotellpb.AutocompleteProxyRespondRequest) (*bacotellpb.AutocompleteProxyRespondResponse, error) {
	return &bacotellpb.AutocompleteProxyRespondResponse{}, s.impl.Respond(decodeApplicationCommandOptionChoices(req.Choices)...)
}

// FocusedOption implements bacotellpb.AutocompleteProxyServer
func (s autocompleteProxyServer) FocusedOption(_ context.Context, req *bacotellpb.AutocompleteProxyFocusedOptionRequest) (*bacotellpb.AutocompleteProxyFocusedOptionResponse, error) {
	name, val, err := s.impl.FocusedOption()
	if err != nil {
		return nil, err
	}

	value, err := structpb.NewValue(val)
	if err != nil {
		return nil, err
	}

	return &bacotellpb.AutocompleteProxyFocusedOptionResponse{Name: name, Value: value}, nil
}

type autocompleteProxyClient struct {
	client bacotellpb.AutocompleteProxyClient
}

var _ common.AutocompleteProxy = autocompleteProxyClient{}

// Complete implements bacotell_common.AutocompleteProxy
func (c autocompleteProxyClient) Respond(choices ...*discordgo.ApplicationCommandOptionChoice) error {
	_, err := c.client.Respond(context.Background(), &bacotellpb.AutocompleteProxyRespondRequest{Choices: encodeApplicationCommandOptionChoices(choices)})
	return err
}

// FocusedOption implements bacotell_common.AutocompleteProxy
func (c autocompleteProxyClient) FocusedOption() (string, any, error) {
	res, err := c.client.FocusedOption(context.Background(), &bacotellpb.AutocompleteProxyFocusedOptionRequest{})
	if err != nil {
		return "", nil, err
	}

	return res.Name, res.Value.AsInterface(), nil
}
