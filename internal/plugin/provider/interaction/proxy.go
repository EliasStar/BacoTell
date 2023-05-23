package interaction

import (
	context "context"

	"github.com/EliasStar/BacoTell/internal/proto/providerpb"
	"github.com/EliasStar/BacoTell/pkg/provider"
	"github.com/bwmarrin/discordgo"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type interactionProxyServer struct {
	providerpb.UnimplementedInteractionProxyServer

	impl provider.InteractionProxy
}

var _ providerpb.InteractionProxyServer = interactionProxyServer{}

// Defer implements providerpb.InteractionProxyServer
func (s interactionProxyServer) Defer(_ context.Context, req *providerpb.DeferRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.impl.Defer(req.Ephemeral, req.SuppressEmbeds, req.Tts)
}

// Respond implements providerpb.InteractionProxyServer
func (s interactionProxyServer) Respond(_ context.Context, req *providerpb.RespondRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.impl.Respond(req.Message, req.Ephemeral, req.SuppressEmbeds, req.Tts)
}

// Followup implements providerpb.InteractionProxyServer
func (s interactionProxyServer) Followup(_ context.Context, req *providerpb.FollowupRequest) (*providerpb.FollowupResponse, error) {
	id, err := s.impl.Followup(req.Message, req.Ephemeral, req.SuppressEmbeds, req.Tts)
	if err != nil {
		return nil, err
	}

	return &providerpb.FollowupResponse{Id: id}, nil
}

// Edit implements providerpb.InteractionProxyServer
func (s interactionProxyServer) Edit(_ context.Context, req *providerpb.EditRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.impl.Edit(req.Id, req.Message)
}

// Delete implements providerpb.InteractionProxyServer
func (s interactionProxyServer) Delete(_ context.Context, req *providerpb.DeleteRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.impl.Delete(req.Id)
}

type interactionProxyClient struct {
	client providerpb.InteractionProxyClient
}

var _ provider.InteractionProxy = interactionProxyClient{}

// Defer implements provider.InteractionProxy
func (c interactionProxyClient) Defer(ephemeral bool, suppressEmbeds bool, tts bool) error {
	_, err := c.client.Defer(context.Background(), &providerpb.DeferRequest{
		Ephemeral:      ephemeral,
		SuppressEmbeds: suppressEmbeds,
		Tts:            tts,
	})

	return err
}

// Respond implements provider.InteractionProxy
func (c interactionProxyClient) Respond(message string, ephemeral bool, suppressEmbeds bool, tts bool) error {
	_, err := c.client.Respond(context.Background(), &providerpb.RespondRequest{
		Message:        message,
		Ephemeral:      ephemeral,
		SuppressEmbeds: suppressEmbeds,
		Tts:            tts,
	})

	return err
}

// Followup implements provider.InteractionProxy
func (c interactionProxyClient) Followup(message string, ephemeral bool, suppressEmbeds bool, tts bool) (string, error) {
	res, err := c.client.Followup(context.Background(), &providerpb.FollowupRequest{
		Message:        message,
		Ephemeral:      ephemeral,
		SuppressEmbeds: suppressEmbeds,
		Tts:            tts,
	})

	if err != nil {
		return "", err
	}

	return res.Id, nil
}

// Edit implements provider.InteractionProxy
func (c interactionProxyClient) Edit(id string, message string) error {
	_, err := c.client.Edit(context.Background(), &providerpb.EditRequest{
		Id:      id,
		Message: message,
	})

	return err
}

// Delete implements provider.InteractionProxy
func (c interactionProxyClient) Delete(id string) error {
	_, err := c.client.Delete(context.Background(), &providerpb.DeleteRequest{
		Id: id,
	})

	return err
}

type executeProxyServer struct {
	providerpb.UnimplementedExecuteProxyServer

	impl provider.ExecuteProxy
}

var _ providerpb.ExecuteProxyServer = executeProxyServer{}

// StringOption implements providerpb.ExecuteProxyServer
func (s executeProxyServer) StringOption(_ context.Context, req *providerpb.OptionRequest) (*providerpb.StringOptionResponse, error) {
	val, err := s.impl.StringOption(req.Name)
	if err != nil {
		return nil, err
	}

	return &providerpb.StringOptionResponse{
		Value: val,
	}, nil
}

// IntegerOption implements providerpb.ExecuteProxyServer
func (s executeProxyServer) IntegerOption(_ context.Context, req *providerpb.OptionRequest) (*providerpb.IntegerOptionResponse, error) {
	val, err := s.impl.IntegerOption(req.Name)
	if err != nil {
		return nil, err
	}

	return &providerpb.IntegerOptionResponse{
		Value: val,
	}, nil
}

// NumberOption implements providerpb.ExecuteProxyServer
func (s executeProxyServer) NumberOption(_ context.Context, req *providerpb.OptionRequest) (*providerpb.NumberOptionResponse, error) {
	val, err := s.impl.NumberOption(req.Name)
	if err != nil {
		return nil, err
	}

	return &providerpb.NumberOptionResponse{
		Value: val,
	}, nil
}

// BooleanOption implements providerpb.ExecuteProxyServer
func (s executeProxyServer) BooleanOption(_ context.Context, req *providerpb.OptionRequest) (*providerpb.BooleanOptionResponse, error) {
	val, err := s.impl.BooleanOption(req.Name)
	if err != nil {
		return nil, err
	}

	return &providerpb.BooleanOptionResponse{
		Value: val,
	}, nil
}

type executeProxyClient struct {
	interactionProxyClient

	client providerpb.ExecuteProxyClient
}

var _ provider.ExecuteProxy = executeProxyClient{}

// StringOption implements provider.ExecuteProxy
func (c executeProxyClient) StringOption(name string) (string, error) {
	res, err := c.client.StringOption(context.Background(), &providerpb.OptionRequest{
		Name: name,
	})

	if err != nil {
		return "", err
	}

	return res.Value, nil
}

// IntegerOption implements provider.ExecuteProxy
func (c executeProxyClient) IntegerOption(name string) (int64, error) {
	res, err := c.client.IntegerOption(context.Background(), &providerpb.OptionRequest{
		Name: name,
	})

	if err != nil {
		return 0, err
	}

	return res.Value, nil
}

// NumberOption implements provider.ExecuteProxy
func (c executeProxyClient) NumberOption(name string) (float64, error) {
	res, err := c.client.NumberOption(context.Background(), &providerpb.OptionRequest{
		Name: name,
	})

	if err != nil {
		return 0, err
	}

	return res.Value, nil
}

// BooleanOption implements provider.ExecuteProxy
func (c executeProxyClient) BooleanOption(name string) (bool, error) {
	res, err := c.client.BooleanOption(context.Background(), &providerpb.OptionRequest{
		Name: name,
	})

	if err != nil {
		return false, err
	}

	return res.Value, nil
}

// UserOption implements provider.ExecuteProxy
func (c executeProxyClient) UserOption(name string) (discordgo.User, error) {
	panic("unimplemented")
}

// RoleOption implements provider.ExecuteProxy
func (c executeProxyClient) RoleOption(name string) (discordgo.Role, error) {
	panic("unimplemented")
}

// ChannelOption implements provider.ExecuteProxy
func (c executeProxyClient) ChannelOption(name string) (discordgo.Channel, error) {
	panic("unimplemented")
}

// AttachmentOption implements provider.ExecuteProxy
func (c executeProxyClient) AttachmentOption(name string) (discordgo.MessageAttachment, error) {
	panic("unimplemented")
}

type handleProxyServer struct {
	providerpb.UnimplementedHandleProxyServer

	impl provider.HandleProxy
}

var _ providerpb.HandleProxyServer = handleProxyServer{}

type handleProxyClient struct {
	interactionProxyClient

	client providerpb.HandleProxyClient
}

var _ provider.HandleProxy = handleProxyClient{}
