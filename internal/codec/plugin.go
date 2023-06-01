package codec

import (
	"context"

	"github.com/EliasStar/BacoTell/internal/proto/bacotellpb"
	common "github.com/EliasStar/BacoTell/pkg/bacotell_common"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type bacotellPlugin struct {
	plugin.NetRPCUnsupportedPlugin
	common.Plugin
}

var (
	_ plugin.Plugin     = bacotellPlugin{}
	_ plugin.GRPCPlugin = bacotellPlugin{}
)

func NewPlugin(pluginImpl common.Plugin) plugin.Plugin {
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
