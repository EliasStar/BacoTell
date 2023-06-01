package codec

import (
	"context"

	"github.com/EliasStar/BacoTell/internal/proto/bacotellpb"
	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
)

type interactionProxyServer struct {
	bacotellpb.UnimplementedInteractionProxyServer

	impl common.InteractionProxy
}

var _ bacotellpb.InteractionProxyServer = interactionProxyServer{}

// Defer implements bacotellpb.InteractionProxyServer
func (s interactionProxyServer) Defer(_ context.Context, req *bacotellpb.DeferRequest) (*bacotellpb.DeferResponse, error) {
	return &bacotellpb.DeferResponse{}, s.impl.Defer(req.Ephemeral)
}

// Respond implements bacotellpb.InteractionProxyServer
func (s interactionProxyServer) Respond(_ context.Context, req *bacotellpb.RespondRequest) (*bacotellpb.RespondResponse, error) {
	return &bacotellpb.RespondResponse{}, s.impl.Respond(*decodeResponse(req.Message), req.Ephemeral)
}

// Followup implements bacotellpb.InteractionProxyServer
func (s interactionProxyServer) Followup(_ context.Context, req *bacotellpb.FollowupRequest) (*bacotellpb.FollowupResponse, error) {
	id, err := s.impl.Followup(*decodeResponse(req.Message), req.Ephemeral)
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

var _ common.InteractionProxy = interactionProxyClient{}

// Defer implements bacotell_common.InteractionProxy
func (c interactionProxyClient) Defer(ephemeral bool) error {
	_, err := c.client.Defer(context.Background(), &bacotellpb.DeferRequest{Ephemeral: ephemeral})

	return err
}

// Respond implements bacotell_common.InteractionProxy
func (c interactionProxyClient) Respond(message common.Response, ephemeral bool) error {
	_, err := c.client.Respond(context.Background(), &bacotellpb.RespondRequest{
		Message:   encodeResponse(&message),
		Ephemeral: ephemeral,
	})

	return err
}

// Followup implements bacotell_common.InteractionProxy
func (c interactionProxyClient) Followup(message common.Response, ephemeral bool) (string, error) {
	res, err := c.client.Followup(context.Background(), &bacotellpb.FollowupRequest{
		Message:   encodeResponse(&message),
		Ephemeral: ephemeral,
	})

	if err != nil {
		return "", err
	}

	return res.Id, nil
}

// Edit implements bacotell_common.InteractionProxy
func (c interactionProxyClient) Edit(id string, message common.Response) error {
	_, err := c.client.Edit(context.Background(), &bacotellpb.EditRequest{
		Id:      id,
		Message: encodeResponse(&message),
	})

	return err
}

// Delete implements bacotell_common.InteractionProxy
func (c interactionProxyClient) Delete(id string) error {
	_, err := c.client.Delete(context.Background(), &bacotellpb.DeleteRequest{Id: id})

	return err
}

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
