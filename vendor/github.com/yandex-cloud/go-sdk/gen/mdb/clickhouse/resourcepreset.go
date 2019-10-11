// Code generated by sdkgen. DO NOT EDIT.

//nolint
package clickhouse

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/clickhouse/v1"
)

//revive:disable

// ResourcePresetServiceClient is a clickhouse.ResourcePresetServiceClient with
// lazy GRPC connection initialization.
type ResourcePresetServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

var _ clickhouse.ResourcePresetServiceClient = &ResourcePresetServiceClient{}

// Get implements clickhouse.ResourcePresetServiceClient
func (c *ResourcePresetServiceClient) Get(ctx context.Context, in *clickhouse.GetResourcePresetRequest, opts ...grpc.CallOption) (*clickhouse.ResourcePreset, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewResourcePresetServiceClient(conn).Get(ctx, in, opts...)
}

// List implements clickhouse.ResourcePresetServiceClient
func (c *ResourcePresetServiceClient) List(ctx context.Context, in *clickhouse.ListResourcePresetsRequest, opts ...grpc.CallOption) (*clickhouse.ListResourcePresetsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewResourcePresetServiceClient(conn).List(ctx, in, opts...)
}
