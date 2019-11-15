package api

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"
	"math"
)

const _ = proto.ProtoPackageIsVersion3

type ExampleService interface {
	Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error)
}

type exampleService struct {
	c    client.Client
	name string
}

// create a new example service
func NewExampleService(name string, c client.Client) ExampleService {
	// if c is nil, create a new client
	if c == nil {
		c = client.NewClient()
	}
	// default service name is "example"
	if len(name) == 0 {
		name = "example"
	}
	return &exampleService{
		c:    c,
		name: name,
	}
}

func (service *exampleService) Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error) {
	req := service.c.NewRequest(service.name, "Example.Call", in)
	out := new(CallResponse)
	err := service.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
