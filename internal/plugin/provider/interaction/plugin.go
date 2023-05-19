package interaction

import (
	"context"

	"github.com/EliasStar/BacoTell/internal/proto/providerpb"
	"github.com/EliasStar/BacoTell/pkg/provider"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type interactionProviderPlugin struct {
	plugin.NetRPCUnsupportedPlugin
	provider.InteractionProvider
}

var (
	_ plugin.Plugin     = interactionProviderPlugin{}
	_ plugin.GRPCPlugin = interactionProviderPlugin{}
)

func NewProviderPlugin(provider provider.InteractionProvider) plugin.Plugin {
	return interactionProviderPlugin{InteractionProvider: provider}
}

// GRPCServer implements plugin.GRPCPlugin
func (p interactionProviderPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	providerpb.RegisterInteractionProviderServer(s, &interactionProviderServer{
		impl:   p.InteractionProvider,
		broker: broker,
	})

	return nil
}

// GRPCClient implements plugin.GRPCPlugin
func (p interactionProviderPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return interactionProviderClient{
		client: providerpb.NewInteractionProviderClient(c),
		broker: broker,
	}, nil
}
