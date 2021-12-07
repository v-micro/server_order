package server

import (
	"context"
	"fmt"
	"server_order/server_common/protobuf/serverorder"
)

type ServerPing struct {}

func (s ServerPing) Get(ctx context.Context, request *serverorder.GetRequest) (*serverorder.GetResponse, error) {
	////标签
	//span, _ := opentracing.StartSpanFromContext(ctx, "Get")
	//defer span.Finish()
	////二级链路
	//Tet(ctx)

	//time.Sleep(time.Second/10)

	fmt.Println("Get....")
	return &serverorder.GetResponse{
		Res:                  "ok1111",
	},nil
}


