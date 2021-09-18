package server

import (
	"context"
	"server_order/server_common/protobuf/serverorder"
)

type ServerPing struct {}

func (s ServerPing) Get(ctx context.Context, request *serverorder.GetRequest) (*serverorder.GetResponse, error) {
	panic("implement me")
}


