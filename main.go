package main

import (
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
	"log"
	"net"
	"server_order/server"
	"server_order/server_common/comutil"
	"server_order/server_common/protobuf/serverorder"
)

func main()  {
	var err error
	var servOpts []grpc.ServerOption
	grpcUrl := ":9991"

	//服务租约
	_, err = comutil.NewServiceRegister([]string{"192.168.59.131:2379"}, "server_order", grpcUrl, 5)
	if err != nil {
		log.Fatalf("服务租约启动失败：",err)
	}
	//监听续租相应chan
	//go ser.ListenLeaseRespChan()

	//链路全局化
	tracer, closer, err := comutil.InitJaeger("server_order","192.168.59.131:6831")
	if err != nil || tracer == nil {
		log.Fatalf("链路追踪启动失败：",err)
		return
	}
	defer closer.Close()
	servOpts = append(servOpts,grpc.UnaryInterceptor(grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithTracer(tracer))))

	//初始化
	lis, err := net.Listen("tcp", grpcUrl)
	if err != nil {
		log.Fatalf("服务端口监听失败: %v", err)
	}

	//新建服务
	s := grpc.NewServer(servOpts...)
	//服务注入
	serverorder.RegisterPingServer(s,&server.ServerPing{})

	log.Print("服务运行成功:" + grpcUrl)
	err = s.Serve(lis)
	if  err != nil {
		log.Fatalf("服务运行失败: %v", err)
	}
}

