package codec

import (
	"context"

	"github.com/EliasStar/BacoTell/internal/proto/bacotellpb"
	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	"github.com/bwmarrin/discordgo"
)

// The server implementation of bacotell_common.InteractionProxy.
type interactionProxyServer struct {
	bacotellpb.UnimplementedInteractionProxyServer

	impl common.InteractionProxy
}

// interactionProxyServer implements bacotellpb.InteractionProxyServer.
var _ bacotellpb.InteractionProxyServer = interactionProxyServer{}

// Defer implements bacotellpb.InteractionProxyServer.
func (s interactionProxyServer) Defer(_ context.Context, req *bacotellpb.InteractionProxyDeferRequest) (*bacotellpb.InteractionProxyDeferResponse, error) {
	return &bacotellpb.InteractionProxyDeferResponse{}, s.impl.Defer(req.Ephemeral)
}

// Respond implements bacotellpb.InteractionProxyServer.
func (s interactionProxyServer) Respond(_ context.Context, req *bacotellpb.InteractionProxyRespondRequest) (*bacotellpb.InteractionProxyRespondResponse, error) {
	return &bacotellpb.InteractionProxyRespondResponse{}, s.impl.Respond(*decodeResponse(req.Message), req.Ephemeral)
}

// Modal implements bacotellpb.InteractionProxyServer.
func (s interactionProxyServer) Modal(_ context.Context, req *bacotellpb.InteractionProxyModalRequest) (*bacotellpb.InteractionProxyModalResponse, error) {
	return &bacotellpb.InteractionProxyModalResponse{}, s.impl.Modal(req.CustomId, req.Title, decodeMessageComponents(req.Components)...)
}

// Followup implements bacotellpb.InteractionProxyServer.
func (s interactionProxyServer) Followup(_ context.Context, req *bacotellpb.InteractionProxyFollowupRequest) (*bacotellpb.InteractionProxyFollowupResponse, error) {
	id, err := s.impl.Followup(*decodeResponse(req.Message), req.Ephemeral)
	if err != nil {
		return nil, err
	}

	return &bacotellpb.InteractionProxyFollowupResponse{Id: id}, nil
}

// Edit implements bacotellpb.InteractionProxyServer.
func (s interactionProxyServer) Edit(_ context.Context, req *bacotellpb.InteractionProxyEditRequest) (*bacotellpb.InteractionProxyEditResponse, error) {
	return &bacotellpb.InteractionProxyEditResponse{}, s.impl.Edit(req.Id, *decodeResponse(req.Message))
}

// Delete implements bacotellpb.InteractionProxyServer.
func (s interactionProxyServer) Delete(_ context.Context, req *bacotellpb.InteractionProxyDeleteRequest) (*bacotellpb.InteractionProxyDeleteResponse, error) {
	return &bacotellpb.InteractionProxyDeleteResponse{}, s.impl.Delete(req.Id)
}

// GuildId implements bacotellpb.InteractionProxyServer.
func (s interactionProxyServer) GuildId(context.Context, *bacotellpb.InteractionProxyGuildIdRequest) (*bacotellpb.InteractionProxyGuildIdResponse, error) {
	id, err := s.impl.GuildID()
	if err != nil {
		return nil, err
	}

	return &bacotellpb.InteractionProxyGuildIdResponse{Id: id}, nil
}

// ChannelId implements bacotellpb.InteractionProxyServer.
func (s interactionProxyServer) ChannelId(context.Context, *bacotellpb.InteractionProxyChannelIdRequest) (*bacotellpb.InteractionProxyChannelIdResponse, error) {
	id, err := s.impl.ChannelID()
	if err != nil {
		return nil, err
	}

	return &bacotellpb.InteractionProxyChannelIdResponse{Id: id}, nil
}

// UserLocale implements bacotellpb.InteractionProxyServer.
func (s interactionProxyServer) UserLocale(context.Context, *bacotellpb.InteractionProxyUserLocaleRequest) (*bacotellpb.InteractionProxyUserLocaleResponse, error) {
	locale, err := s.impl.UserLocale()
	if err != nil {
		return nil, err
	}

	return &bacotellpb.InteractionProxyUserLocaleResponse{Locale: string(locale)}, nil
}

// GuildLocale implements bacotellpb.InteractionProxyServer.
func (s interactionProxyServer) GuildLocale(context.Context, *bacotellpb.InteractionProxyGuildLocaleRequest) (*bacotellpb.InteractionProxyGuildLocaleResponse, error) {
	locale, err := s.impl.GuildLocale()
	if err != nil {
		return nil, err
	}

	return &bacotellpb.InteractionProxyGuildLocaleResponse{Locale: string(locale)}, nil
}

// User implements bacotellpb.InteractionProxyServer.
func (s interactionProxyServer) User(context.Context, *bacotellpb.InteractionProxyUserRequest) (*bacotellpb.InteractionProxyUserResponse, error) {
	user, err := s.impl.User()
	if err != nil {
		return nil, err
	}

	return &bacotellpb.InteractionProxyUserResponse{User: encodeUser(user)}, nil
}

// Member implements bacotellpb.InteractionProxyServer.
func (s interactionProxyServer) Member(context.Context, *bacotellpb.InteractionProxyMemberRequest) (*bacotellpb.InteractionProxyMemberResponse, error) {
	member, err := s.impl.Member()
	if err != nil {
		return nil, err
	}

	return &bacotellpb.InteractionProxyMemberResponse{Member: encodeMember(member)}, nil
}

// Message implements bacotellpb.InteractionProxyServer.
func (s interactionProxyServer) Message(context.Context, *bacotellpb.InteractionProxyMessageRequest) (*bacotellpb.InteractionProxyMessageResponse, error) {
	message, err := s.impl.Message()
	if err != nil {
		return nil, err
	}

	return &bacotellpb.InteractionProxyMessageResponse{Message: encodeMessage(message)}, nil
}

// Permissions implements bacotellpb.InteractionProxyServer.
func (s interactionProxyServer) Permissions(context.Context, *bacotellpb.InteractionProxyPermissionsRequest) (*bacotellpb.InteractionProxyPermissionsResponse, error) {
	permissions, err := s.impl.Permissions()
	if err != nil {
		return nil, err
	}

	return &bacotellpb.InteractionProxyPermissionsResponse{Permissions: permissions}, nil
}

// The client implementation of bacotell_common.InteractionProxy.
type interactionProxyClient struct {
	client bacotellpb.InteractionProxyClient
}

// interactionProxyClient implements bacotell_common.InteractionProxy.
var _ common.InteractionProxy = interactionProxyClient{}

// Defer implements bacotell_common.InteractionProxy.
func (c interactionProxyClient) Defer(ephemeral bool) error {
	_, err := c.client.Defer(context.Background(), &bacotellpb.InteractionProxyDeferRequest{Ephemeral: ephemeral})

	return err
}

// Respond implements bacotell_common.InteractionProxy.
func (c interactionProxyClient) Respond(message common.Response, ephemeral bool) error {
	_, err := c.client.Respond(context.Background(), &bacotellpb.InteractionProxyRespondRequest{
		Message:   encodeResponse(&message),
		Ephemeral: ephemeral,
	})

	return err
}

// Modal implements bacotell_common.InteractionProxy.
func (c interactionProxyClient) Modal(customId string, title string, components ...discordgo.MessageComponent) error {
	_, err := c.client.Modal(context.Background(), &bacotellpb.InteractionProxyModalRequest{
		CustomId:   customId,
		Title:      title,
		Components: encodeMessageComponents(components),
	})

	return err
}

// Followup implements bacotell_common.InteractionProxy.
func (c interactionProxyClient) Followup(message common.Response, ephemeral bool) (string, error) {
	res, err := c.client.Followup(context.Background(), &bacotellpb.InteractionProxyFollowupRequest{
		Message:   encodeResponse(&message),
		Ephemeral: ephemeral,
	})

	if err != nil {
		return "", err
	}

	return res.Id, nil
}

// Edit implements bacotell_common.InteractionProxy.
func (c interactionProxyClient) Edit(id string, message common.Response) error {
	_, err := c.client.Edit(context.Background(), &bacotellpb.InteractionProxyEditRequest{
		Id:      id,
		Message: encodeResponse(&message),
	})

	return err
}

// Delete implements bacotell_common.InteractionProxy.
func (c interactionProxyClient) Delete(id string) error {
	_, err := c.client.Delete(context.Background(), &bacotellpb.InteractionProxyDeleteRequest{Id: id})

	return err
}

// GuildID implements bacotell_common.InteractionProxy.
func (c interactionProxyClient) GuildID() (string, error) {
	res, err := c.client.GuildId(context.Background(), &bacotellpb.InteractionProxyGuildIdRequest{})
	if err != nil {
		return "", err
	}

	return res.Id, nil
}

// ChannelID implements bacotell_common.InteractionProxy.
func (c interactionProxyClient) ChannelID() (string, error) {
	res, err := c.client.ChannelId(context.Background(), &bacotellpb.InteractionProxyChannelIdRequest{})
	if err != nil {
		return "", err
	}

	return res.Id, nil
}

// UserLocale implements bacotell_common.InteractionProxy.
func (c interactionProxyClient) UserLocale() (discordgo.Locale, error) {
	res, err := c.client.UserLocale(context.Background(), &bacotellpb.InteractionProxyUserLocaleRequest{})
	if err != nil {
		return "", err
	}

	return discordgo.Locale(res.Locale), nil
}

// GuildLocale implements bacotell_common.InteractionProxy.
func (c interactionProxyClient) GuildLocale() (discordgo.Locale, error) {
	res, err := c.client.GuildLocale(context.Background(), &bacotellpb.InteractionProxyGuildLocaleRequest{})
	if err != nil {
		return "", err
	}

	return discordgo.Locale(res.Locale), nil
}

// User implements bacotell_common.InteractionProxy.
func (c interactionProxyClient) User() (*discordgo.User, error) {
	res, err := c.client.User(context.Background(), &bacotellpb.InteractionProxyUserRequest{})
	if err != nil {
		return nil, err
	}

	return decodeUser(res.User), nil
}

// Member implements bacotell_common.InteractionProxy.
func (c interactionProxyClient) Member() (*discordgo.Member, error) {
	res, err := c.client.Member(context.Background(), &bacotellpb.InteractionProxyMemberRequest{})
	if err != nil {
		return nil, err
	}

	return decodeMember(res.Member), nil
}

// Message implements bacotell_common.InteractionProxy.
func (c interactionProxyClient) Message() (*discordgo.Message, error) {
	res, err := c.client.Message(context.Background(), &bacotellpb.InteractionProxyMessageRequest{})
	if err != nil {
		return nil, err
	}

	return decodeMessage(res.Message), nil
}

// Permissions implements bacotell_common.InteractionProxy.
func (c interactionProxyClient) Permissions() (int64, error) {
	res, err := c.client.Permissions(context.Background(), &bacotellpb.InteractionProxyPermissionsRequest{})
	if err != nil {
		return 0, err
	}

	return res.Permissions, nil
}

// encodeResponse encodes a bacotell_common.Response into a bacotellpb.Response.
func encodeResponse(response *common.Response) *bacotellpb.Response {
	if response == nil {
		return nil
	}

	return &bacotellpb.Response{
		Content:         response.Content,
		SuppressEmbeds:  response.SuppressEmbeds,
		Tts:             response.TTS,
		AllowedMentions: encodeMessageAllowedMention(&response.AllowedMentions),
		Components:      encodeMessageComponents(response.Components),
		Embeds:          encodeMessageEmbeds(response.Embeds),
		Files:           encodeFiles(response.Files),
	}
}

// decodeResponse decodes a bacotellpb.Response into a bacotell_common.Response.
func decodeResponse(response *bacotellpb.Response) *common.Response {
	if response == nil {
		return nil
	}

	return &common.Response{
		Content:         response.Content,
		SuppressEmbeds:  response.SuppressEmbeds,
		TTS:             response.Tts,
		AllowedMentions: *decodeMessageAllowedMention(response.AllowedMentions),
		Components:      decodeMessageComponents(response.Components),
		Embeds:          decodeMessageEmbeds(response.Embeds),
		Files:           decodeFiles(response.Files),
	}
}
