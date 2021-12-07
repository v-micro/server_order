package main

import (
	"fmt"
	 "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
	"log"
	"net"
	"server_order/server"
	"server_order/server_common/comutil"
	"server_order/server_common/protobuf/serverorder"
)

func main()  {
	var servOpts []grpc.ServerOption
	var grpcUrl = ":3333"


	//服务租约
	var endpoints = []string{"192.168.59.131:2379"}
	_, err := comutil.NewServiceRegister(endpoints, "server_order", grpcUrl, 5)
	if err != nil {
		log.Fatalln(err)
	}
	//监听续租相应chan
	//go ser.ListenLeaseRespChan()

	//链路全局化
	tracer, closer, err := comutil.InitJaeger("server_order","192.168.59.131:6831")
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	defer closer.Close()
	servOpts = append(servOpts,grpc.UnaryInterceptor(grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithTracer(tracer))))

	//初始化
	lis, err := net.Listen("tcp", grpcUrl)
	if err != nil {log.Fatalf("启动失败: %v", err)}

	//新启动rpc服务
	s := grpc.NewServer(servOpts...)

	//服务注入
	serverorder.RegisterPingServer(s,&server.ServerPing{})

	//服务启动
	log.Printf(fmt.Sprintf("服务端口启动成功 %s", grpcUrl))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("当前服务启动失败: %v", err)
	}
}

