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
	greq, err := toproto.GetProviderSchema_Request(req)
	if err != nil {
		return nil, err
	}
	gresp, err := c.client.GetSchema(ctx, greq)
	if err != nil {
		return nil, err
	}
	return fromproto.GetProviderSchemaResponse(gresp)
}
