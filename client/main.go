package main

import (
	"context"
	"fmt"
	"github.com/mensylisir/KAnsible/api"
	"google.golang.org/grpc"
	"log"
)

func main() {
	// 连接服务器
	conn, err := grpc.Dial("localhost:12800", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接服务端失败: %s\n", err)
		return
	}
	defer conn.Close()
	// 新建一个客户端
	c := api.NewAnsibleServerClient(conn)

	helloReq := &api.InventoryRequest{
		Item: []*api.Node{
			{
				Ip:       "192.168.7.128",
				Port:     "22",
				Password: "xiaoming98",
				Role:     "master",
				Name:     "node1",
			},
		},
	}
	// 调用服务端函数
	r, err := c.GenerateYaml(context.Background(), helloReq)
	if err != nil {
		fmt.Printf("调用服务端代码失败: %s", err)
		return
	}
	fmt.Println(r.Message)


	bbb := &api.PlayRequests{
		Message: "distribute",
	}

	r2, err := c.StreamRunPlaybook(context.Background(), bbb)
	if err != nil {
		fmt.Printf("调用服务端代码失败: %s", err)
		return
	}
	for {
		bb, err := r2.Recv()
		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}
		data := bb.GetRes()
		fmt.Println(data)
		if "success" == data || "failure" == data {
			break
		}
	}

	aaa := &api.PlayRequests{
		Message: "install",
	}

	r1, err := c.StreamRunPlaybook(context.Background(), aaa)
	if err != nil {
		fmt.Printf("调用服务端代码失败: %s", err)
		return
	}
	for {
		bb, err := r1.Recv()
		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}
		data := bb.GetRes()
		fmt.Println(data)
		if "success" == data || "failure" == data {
			break
		}
	}
}
