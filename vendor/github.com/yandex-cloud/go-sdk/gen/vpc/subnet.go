// Code generated by sdkgen. DO NOT EDIT.

//nolint
package vpc

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1"
)

//revive:disable

// SubnetServiceClient is a vpc.SubnetServiceClient with
// lazy GRPC connection initialization.
type SubnetServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

var _ vpc.SubnetServiceClient = &SubnetServiceClient{}

// Create implements vpc.SubnetServiceClient
func (c *SubnetServiceClient) Create(ctx context.Context, in *vpc.CreateSubnetRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return vpc.NewSubnetServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements vpc.SubnetServiceClient
func (c *SubnetServiceClient) Delete(ctx context.Context, in *vpc.DeleteSubnetRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return vpc.NewSubnetServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements vpc.SubnetServiceClient
func (c *SubnetServiceClient) Get(ctx context.Context, in *vpc.GetSubnetRequest, opts ...grpc.CallOption) (*vpc.Subnet, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return vpc.NewSubnetServiceClient(conn).Get(ctx, in, opts...)
}

// List implements vpc.SubnetServiceClient
func (c *SubnetServiceClient) List(ctx context.Context, in *vpc.ListSubnetsRequest, opts ...grpc.CallOption) (*vpc.ListSubnetsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return vpc.NewSubnetServiceClient(conn).List(ctx, in, opts...)
}

// ListOperations implements vpc.SubnetServiceClient
func (c *SubnetServiceClient) ListOperations(ctx context.Context, in *vpc.ListSubnetOperationsRequest, opts ...grpc.CallOption) (*vpc.ListSubnetOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return vpc.NewSubnetServiceClient(conn).ListOperations(ctx, in, opts...)
}

// Update implements vpc.SubnetServiceClient
func (c *SubnetServiceClient) Update(ctx context.Context, in *vpc.UpdateSubnetRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return vpc.NewSubnetServiceClient(conn).Update(ctx, in, opts...)
}
