package codec

import (
	"context"

	"github.com/EliasStar/BacoTell/internal/proto/bacotellpb"
	"github.com/EliasStar/BacoTell/pkg/bacotell"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type bacotellPlugin struct {
	plugin.NetRPCUnsupportedPlugin
	bacotell.Plugin
}

var (
	_ plugin.Plugin     = bacotellPlugin{}
	_ plugin.GRPCPlugin = bacotellPlugin{}
)

func NewBacoTellPlugin(pluginImpl bacotell.Plugin) plugin.Plugin {
	return bacotellPlugin{Plugin: pluginImpl}
}

// GRPCServer implements plugin.GRPCPlugin
func (p bacotellPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	bacotellpb.RegisterPluginServer(s, &pluginServer{
		impl:   p.Plugin,
		broker: broker,
	})

	return nil
}

// GRPCClient implements plugin.GRPCPlugin
func (p bacotellPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return pluginClient{
		client: bacotellpb.NewPluginClient(c),
		broker: broker,
	}, nil
}
