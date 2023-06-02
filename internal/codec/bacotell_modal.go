package codec

import (
	"context"

	"github.com/EliasStar/BacoTell/internal/proto/bacotellpb"
	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type modalServer struct {
	bacotellpb.UnimplementedModalServer

	impl   common.Modal
	broker *plugin.GRPCBroker
}

var _ bacotellpb.ModalServer = modalServer{}

// CustomId implements bacotellpb.ModalServer.
func (s modalServer) CustomId(context.Context, *bacotellpb.ModalCustomIdRequest) (*bacotellpb.ModalCustomIdResponse, error) {
	id, err := s.impl.CustomID()
	if err != nil {
		return nil, err
	}

	return &bacotellpb.ModalCustomIdResponse{CustomId: id}, nil
}

// Submit implements bacotellpb.ModalServer.
func (s modalServer) Submit(_ context.Context, req *bacotellpb.ModalSubmitRequest) (*bacotellpb.ModalSubmitResponse, error) {
	conn, err := s.broker.Dial(req.SubmitProxyId)
	if err != nil {
		return nil, err
	}

	err = s.impl.Submit(submitProxyClient{
		interactionProxyClient: interactionProxyClient{
			client: bacotellpb.NewInteractionProxyClient(conn),
		},
		client: bacotellpb.NewSubmitProxyClient(conn),
	})

	conn.Close()
	return &bacotellpb.ModalSubmitResponse{}, err
}

type modalClient struct {
	client bacotellpb.ModalClient
	broker *plugin.GRPCBroker
}

var _ common.Modal = modalClient{}

// CustomID implements bacotell_common.Modal.
func (c modalClient) CustomID() (string, error) {
	res, err := c.client.CustomId(context.Background(), &bacotellpb.ModalCustomIdRequest{})
	if err != nil {
		return "", err
	}

	return res.CustomId, nil
}

// Submit implements bacotell_common.Modal.
func (c modalClient) Submit(proxy common.SubmitProxy) error {
	var server *grpc.Server

	id := c.broker.NextId()
	go c.broker.AcceptAndServe(id, func(opts []grpc.ServerOption) *grpc.Server {
		server = grpc.NewServer(opts...)

		srv := submitProxyServer{
			interactionProxyServer: interactionProxyServer{
				impl: proxy,
			},
			impl: proxy,
		}

		bacotellpb.RegisterSubmitProxyServer(server, srv)
		bacotellpb.RegisterInteractionProxyServer(server, srv)

		return server
	})

	_, err := c.client.Submit(context.Background(), &bacotellpb.ModalSubmitRequest{SubmitProxyId: id})

	server.Stop()
	return err
}

type submitProxyServer struct {
	bacotellpb.UnimplementedSubmitProxyServer
	interactionProxyServer

	impl common.SubmitProxy
}

var (
	_ bacotellpb.InteractionProxyServer = submitProxyServer{}
	_ bacotellpb.SubmitProxyServer      = submitProxyServer{}
)

type submitProxyClient struct {
	interactionProxyClient

	client bacotellpb.SubmitProxyClient
}

var (
	_ common.InteractionProxy = submitProxyClient{}
	_ common.SubmitProxy      = submitProxyClient{}
)
