package server

import (
	"context"
	"fmt"
	"server_order/server_common/protobuf/serverorder"
)

type ServerPing struct {}

func (s ServerPing) Get(ctx context.Context, request *serverorder.GetRequest) (*serverorder.GetResponse, error) {
	fmt.Println("Get....")

	return &serverorder.GetResponse{
		Res:                  "ok",
	},nil
}


