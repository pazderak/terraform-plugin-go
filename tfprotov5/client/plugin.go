package tf5client

import (
	"context"
	"errors"
	"net/rpc"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

// GRPCProviderPlugin is an implementation of the
// github.com/hashicorp/go-plugin#Plugin and
// github.com/hashicorp/go-plugin#GRPCPlugin interfaces, indicating how to
// serve tfprotov5.ProviderServers as gRPC plugins for go-plugin.
type GRPCProviderPlugin struct {
}

// Server always returns an error; we're only implementing the GRPCPlugin
// interface, not the Plugin interface.
func (p *GRPCProviderPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return nil, errors.New("terraform-plugin-go only implements gRPC clients")
}

// Client always returns an error; we're only implementing the GRPCPlugin
// interface, not the Plugin interface.
func (p *GRPCProviderPlugin) Client(*plugin.MuxBroker, *rpc.Client) (interface{}, error) {
	return nil, errors.New("terraform-plugin-go only implements gRPC clients")
}

// GRPCClient always returns an error; we're only implementing the server half
// of the interface.
func (p *GRPCProviderPlugin) GRPCClient(ctx context.Context, b *plugin.GRPCBroker, cc *grpc.ClientConn) (interface{}, error) {
	return NewClient(cc), nil
}

// GRPCServer registers the gRPC provider server with the gRPC server that
// go-plugin is standing up.
func (p *GRPCProviderPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	return errors.New("terraform-plugin-go only implements gRPC clients")
}
