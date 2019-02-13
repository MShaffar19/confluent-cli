package kafka

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"

	chttp "github.com/confluentinc/ccloud-sdk-go"
	"github.com/confluentinc/cli/shared"
)

// Name description used for registering/disposing GRPC components
const Name = "ccloud-kafka-plugin"

// Plugin mates an interface with Hashicorp plugin object
type Plugin struct {
	plugin.NetRPCUnsupportedPlugin
	Impl chttp.Kafka
}

// GRPCClient registers a GRPC client
func (p *Plugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: NewKafkaClient(c)}, nil
}

// GRPCServer registers a GRPC Server
func (p *Plugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	RegisterKafkaServer(s, &GRPCServer{p.Impl})
	return nil
}

// Check that Plugin satisfies GPRCPlugin interface.
var _ plugin.GRPCPlugin = &Plugin{}

func init() {
	shared.PluginMap[Name] = &Plugin{}
}
