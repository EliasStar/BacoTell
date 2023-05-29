package codec

import (
	"context"

	"github.com/EliasStar/BacoTell/internal/proto/bacotellpb"
	"github.com/EliasStar/BacoTell/pkg/bacotell"
)

type interactionProxyServer struct {
	bacotellpb.UnimplementedInteractionProxyServer

	impl bacotell.InteractionProxy
}

var _ bacotellpb.InteractionProxyServer = interactionProxyServer{}

// Defer implements bacotellpb.InteractionProxyServer
func (s interactionProxyServer) Defer(_ context.Context, req *bacotellpb.DeferRequest) (*bacotellpb.DeferResponse, error) {
	return &bacotellpb.DeferResponse{}, s.impl.Defer(req.Ephemeral, req.SuppressEmbeds, req.Tts)
}

// Respond implements bacotellpb.InteractionProxyServer
func (s interactionProxyServer) Respond(_ context.Context, req *bacotellpb.RespondRequest) (*bacotellpb.RespondResponse, error) {
	return &bacotellpb.RespondResponse{}, s.impl.Respond(*decodeResponse(req.Message), req.Ephemeral, req.SuppressEmbeds, req.Tts)
}

// Followup implements bacotellpb.InteractionProxyServer
func (s interactionProxyServer) Followup(_ context.Context, req *bacotellpb.FollowupRequest) (*bacotellpb.FollowupResponse, error) {
	id, err := s.impl.Followup(*decodeResponse(req.Message), req.Ephemeral, req.SuppressEmbeds, req.Tts)
	if err != nil {
		return nil, err
	}

	return &bacotellpb.FollowupResponse{Id: id}, nil
}

// Edit implements bacotellpb.InteractionProxyServer
func (s interactionProxyServer) Edit(_ context.Context, req *bacotellpb.EditRequest) (*bacotellpb.EditResponse, error) {
	return &bacotellpb.EditResponse{}, s.impl.Edit(req.Id, *decodeResponse(req.Message))
}

// Delete implements bacotellpb.InteractionProxyServer
func (s interactionProxyServer) Delete(_ context.Context, req *bacotellpb.DeleteRequest) (*bacotellpb.DeleteResponse, error) {
	return &bacotellpb.DeleteResponse{}, s.impl.Delete(req.Id)
}

type interactionProxyClient struct {
	client bacotellpb.InteractionProxyClient
}

var _ bacotell.InteractionProxy = interactionProxyClient{}

// Defer implements bacotell.InteractionProxy
func (c interactionProxyClient) Defer(ephemeral bool, suppressEmbeds bool, tts bool) error {
	_, err := c.client.Defer(context.Background(), &bacotellpb.DeferRequest{
		Ephemeral:      ephemeral,
		SuppressEmbeds: suppressEmbeds,
		Tts:            tts,
	})

	return err
}

// Respond implements bacotell.InteractionProxy
func (c interactionProxyClient) Respond(message bacotell.Response, ephemeral bool, suppressEmbeds bool, tts bool) error {
	_, err := c.client.Respond(context.Background(), &bacotellpb.RespondRequest{
		Message:        encodeResponse(&message),
		Ephemeral:      ephemeral,
		SuppressEmbeds: suppressEmbeds,
		Tts:            tts,
	})

	return err
}

// Followup implements bacotell.InteractionProxy
func (c interactionProxyClient) Followup(message bacotell.Response, ephemeral bool, suppressEmbeds bool, tts bool) (string, error) {
	res, err := c.client.Followup(context.Background(), &bacotellpb.FollowupRequest{
		Message:        encodeResponse(&message),
		Ephemeral:      ephemeral,
		SuppressEmbeds: suppressEmbeds,
		Tts:            tts,
	})

	if err != nil {
		return "", err
	}

	return res.Id, nil
}

// Edit implements bacotell.InteractionProxy
func (c interactionProxyClient) Edit(id string, message bacotell.Response) error {
	_, err := c.client.Edit(context.Background(), &bacotellpb.EditRequest{
		Id:      id,
		Message: encodeResponse(&message),
	})

	return err
}

// Delete implements bacotell.InteractionProxy
func (c interactionProxyClient) Delete(id string) error {
	_, err := c.client.Delete(context.Background(), &bacotellpb.DeleteRequest{Id: id})

	return err
}

func encodeResponse(response *bacotell.Response) *bacotellpb.Response {
	if response == nil {
		return nil
	}

	return &bacotellpb.Response{
		Content:         response.Content,
		AllowedMentions: encodeMessageAllowedMention(&response.AllowedMentions),
		Components:      encodeMessageComponents(response.Components),
		Embeds:          encodeMessageEmbeds(response.Embeds),
		Files:           encodeFiles(response.Files),
	}
}

func decodeResponse(response *bacotellpb.Response) *bacotell.Response {
	if response == nil {
		return nil
	}

	return &bacotell.Response{
		Content:         response.Content,
		AllowedMentions: *decodeMessageAllowedMention(response.AllowedMentions),
		Components:      decodeMessageComponents(response.Components),
		Embeds:          decodeMessageEmbeds(response.Embeds),
		Files:           decodeFiles(response.Files),
	}
}
