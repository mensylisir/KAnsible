package main

import (
	"context"
	"fmt"
	"github.com/mensylisir/KAnsible/kapi"
	"google.golang.org/grpc"
)

func main() {
	// 连接服务器
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接服务端失败: %s\n", err)
		return
	}
	defer conn.Close()
	// 新建一个客户端
	c := kapi.NewAnsibleServerClient(conn)

	helloReq := &kapi.InventoryRequest{
		Item: []*kapi.Node{
			{
				Ip:       "192.168.7.127",
				Port:     "22",
				Password: "Def@u1tpwd",
				Role:     "master",
				Name:     "node1",
			},
			{
				Ip:       "192.168.7.128",
				Port:     "22",
				Password: "Def@u1tpwd",
				Role:     "worker",
				Name:     "node2",
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


	configReq := &kapi.ConfigRequest{
		ClusterName: "aaaaa",
	}

	resp, err := c.CheckConfiguration(context.Background(), configReq)
	if err != nil {
		fmt.Printf("调用服务端代码失败: %s", err)
		return
	}
	fmt.Println(resp.Message)


	//bbb := &kapi.PlayRequests{
	//	Message: "distribute",
	//}
	//
	//r2, err := c.StreamRunPlaybook(context.Background(), bbb)
	//if err != nil {
	//	fmt.Printf("调用服务端代码失败: %s", err)
	//	return
	//}
	//for {
	//	bb, err := r2.Recv()
	//	if err != nil {
	//		log.Fatalf("Error: %v\n", err)
	//	}
	//	data := bb.GetRes()
	//	fmt.Println(data)
	//	if "success" == data {
	//		break
	//	}
	//	if  "failure" == data {
	//		return
	//	}
	//}
	//
	//aaa := &kapi.PlayRequests{
	//	Message: "install",
	//}
	//
	//r1, err := c.StreamRunPlaybook(context.Background(), aaa)
	//if err != nil {
	//	fmt.Printf("调用服务端代码失败: %s", err)
	//	return
	//}
	//for {
	//	bb, err := r1.Recv()
	//	if err != nil {
	//		log.Fatalf("Error: %v\n", err)
	//	}
	//	data := bb.GetRes()
	//	fmt.Println(data)
	//	if "success" == data || "failure" == data {
	//		break
	//	}
	//}
}
