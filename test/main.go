package main

import (
	"context"
	"log"
	"time"

	v1 "babycare/api/tree/v1"

	"google.golang.org/grpc"
)

func main() {
	// 设置gRPC连接
	conn, err := grpc.Dial("127.0.0.1:9000",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 初始化Tree服务的客户端
	client := v1.NewTreeClient(conn)

	// 设置超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 调用GetTree方法
	resp, err := client.GetTree(ctx, &v1.GetTreeRequest{Id: 123})
	if err != nil {
		log.Fatalf("could not get tree: %v", err)
	}

	// 打印响应
	log.Printf("Tree: %s", resp.GetTree())
}