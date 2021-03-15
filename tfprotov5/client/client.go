package tf5client

import (
	"context"

	"google.golang.org/grpc"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/fromproto"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/toproto"
)

func NewClient(conn grpc.ClientConnInterface) *ProviderClient {
	return &ProviderClient{
		client: tfplugin5.NewProviderClient(conn),
	}
}

type ProviderClient struct {
	client tfplugin5.ProviderClient
}

func (c *ProviderClient) GetProviderSchema(ctx context.Context, req *tfprotov5.GetProviderSchemaRequest) (*tfprotov5.GetProviderSchemaResponse, error) {
	r, err := toproto.GetProviderSchema_Request(req)
	if err != nil {
		return nil, err
	}
	rr, err := c.client.GetSchema(ctx, r)
	if err != nil {
		return nil, err
	}
	return fromproto.GetProviderSchemaResponse(rr)
}

func (c *ProviderClient) PrepareProviderConfig(ctx context.Context, req *tfprotov5.PrepareProviderConfigRequest) (*tfprotov5.PrepareProviderConfigResponse, error) {
	r, err := toproto.PrepareProviderConfig_Request(req)
	if err != nil {
		return nil, err
	}
	rr, err := c.client.PrepareProviderConfig(ctx, r)
	if err != nil {
		return nil, err
	}
	return fromproto.PrepareProviderConfigResponse(rr)
}

func (c *ProviderClient) Configure(ctx context.Context, req *tfprotov5.ConfigureProviderRequest) (*tfprotov5.ConfigureProviderResponse, error) {
	r, err := toproto.Configure_Request(req)
	if err != nil {
		return nil, err
	}
	rr, err := c.client.Configure(ctx, r)
	if err != nil {
		return nil, err
	}
	return fromproto.ConfigureProviderResponse(rr)
}

func (c *ProviderClient) ReadResource(ctx context.Context, req *tfprotov5.ReadResourceRequest) (*tfprotov5.ReadResourceResponse, error) {
	r, err := toproto.ReadResource_Request(req)
	if err != nil {
		return nil, err
	}
	rr, err := c.client.ReadResource(ctx, r)
	if err != nil {
		return nil, err
	}
	return fromproto.ReadResourceResponse(rr)
}
