package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"server_order/server"
	"server_order/server_common/comutil"
	"server_order/server_common/protobuf/serverorder"
)

func main()  {
	grpcUrl := ":10003"

	//服务租约
	var endpoints = []string{"192.168.59.131:2379"}
	_, err := comutil.NewServiceRegister(endpoints, "server_order", grpcUrl, 5)
	if err != nil {
		log.Fatalln(err)
	}
	////监听续租相应chan
	//go ser.ListenLeaseRespChan()
	//select {
	//	case <-time.After(20 * time.Second):
	//    ser.Close()
	//}

	//初始化
	lis, err := net.Listen("tcp", grpcUrl)
	if err != nil {log.Fatalf("启动失败: %v", err)}
	s := grpc.NewServer()

	//服务注入
	serverorder.RegisterPingServer(s,&server.ServerPing{})

	//服务启动
	log.Printf(fmt.Sprintf("服务端口启动成功 %s", grpcUrl))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("当前服务启动失败: %v", err)
	}
}

