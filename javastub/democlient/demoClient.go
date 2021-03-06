// Code generated by goctl. DO NOT EDIT!
// Source: demo.proto

package democlient

import (
	"context"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"

	"demoGo/javastub/demopb"
)

type (
	HelloReq  = demopb.HelloReq
	HelloResp = demopb.HelloResp

	DemoClient interface {
		SayHello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloResp, error)
	}

	defaultDemoClient struct {
		cli zrpc.Client
	}
)

func NewDemoClient(cli zrpc.Client) DemoClient {
	return &defaultDemoClient{
		cli: cli,
	}
}

func (m *defaultDemoClient) SayHello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloResp, error) {
	client := demopb.NewDemoClientClient(m.cli.Conn())
	return client.SayHello(ctx, in, opts...)
}
