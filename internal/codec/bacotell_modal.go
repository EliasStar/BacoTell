package codec

import (
	"context"

	"github.com/EliasStar/BacoTell/internal/proto/bacotellpb"
	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

// The server implementation of bacotell_common.Modal.
type modalServer struct {
	bacotellpb.UnimplementedModalServer

	impl   common.Modal
	broker *plugin.GRPCBroker
}

// modalServer implements bacotellpb.ModalServer.
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

// The client implementation of bacotell_common.Modal.
type modalClient struct {
	client bacotellpb.ModalClient
	broker *plugin.GRPCBroker
}

// modalClient implements bacotell_common.Modal.
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
	srv := submitProxyServer{
		interactionProxyServer: interactionProxyServer{
			impl: proxy,
		},
		impl: proxy,
	}

	go c.broker.AcceptAndServe(id, func(opts []grpc.ServerOption) *grpc.Server {
		server = grpc.NewServer(opts...)
		bacotellpb.RegisterSubmitProxyServer(server, srv)
		bacotellpb.RegisterInteractionProxyServer(server, srv)
		return server
	})

	_, err := c.client.Submit(context.Background(), &bacotellpb.ModalSubmitRequest{SubmitProxyId: id})

	server.Stop()
	return err
}

// The server implementation of bacotell_common.SubmitProxy.
type submitProxyServer struct {
	bacotellpb.UnimplementedSubmitProxyServer
	interactionProxyServer

	impl common.SubmitProxy
}

// submitProxyServer implements bacotellpb.InteractionProxyServer, bacotellpb.SubmitProxyServer.
var (
	_ bacotellpb.InteractionProxyServer = submitProxyServer{}
	_ bacotellpb.SubmitProxyServer      = submitProxyServer{}
)

// InputValue implements bacotellpb.SubmitProxyServer.
func (s submitProxyServer) InputValue(_ context.Context, req *bacotellpb.SubmitProxyInputValueRequest) (*bacotellpb.SubmitProxyInputValueResponse, error) {
	val, err := s.impl.InputValue(req.CustomId)
	if err != nil {
		return nil, err
	}

	return &bacotellpb.SubmitProxyInputValueResponse{Value: val}, nil
}

// The client implementation of bacotell_common.SubmitProxy.
type submitProxyClient struct {
	interactionProxyClient

	client bacotellpb.SubmitProxyClient
}

// submitProxyClient implements bacotell_common.InteractionProxy, bacotell_common.SubmitProxy.
var (
	_ common.InteractionProxy = submitProxyClient{}
	_ common.SubmitProxy      = submitProxyClient{}
)

// InputValue implements bacotell_common.SubmitProxy.
func (c submitProxyClient) InputValue(customID string) (string, error) {
	res, err := c.client.InputValue(context.Background(), &bacotellpb.SubmitProxyInputValueRequest{
		CustomId: customID,
	})

	if err != nil {
		return "", err
	}

	return res.Value, nil
}
