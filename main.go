package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"server_order/server"
	"server_order/server_common/config"
	"server_order/server_common/protobuf/serverorder"
)

func main()  {
	//初始化
	lis, err := net.Listen("tcp", config.ORDERPROT)
	if err != nil {log.Fatalf("启动失败: %v", err)}
	s := grpc.NewServer()

	//服务注入
	serverorder.RegisterPingServer(s,&server.ServerPing{})

	//服务启动
	log.Printf(fmt.Sprintf("服务端口启动成功 %s", config.ORDERPROT))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("当前服务启动失败: %v", err)
	}
}
