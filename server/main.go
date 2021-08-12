package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/mensylisir/KAnsible/ansible"
	"github.com/mensylisir/KAnsible/config"
	"github.com/mensylisir/KAnsible/constant"
	"github.com/mensylisir/KAnsible/kapi"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type server struct{}

func (s *server) StreamRunPlaybook(requests *kapi.PlayRequests, response kapi.AnsibleServer_StreamRunPlaybookServer) error {
	config.GetConfig()
	revMsg := make(chan string)
	switch requests.Message {
	case constant.INSTALL_ACTION:
		go ansible.InstallKubernetes(revMsg)
	case constant.UPDATE_ACTION:
		go ansible.InstallKubernetes(revMsg)
	case constant.RESET_ACTION:
		go ansible.ResetKubernetes(revMsg)
	case constant.DISTRIBUTE_KEY:
		go ansible.DistributePublicKey(revMsg)
	default:
		errMsg := fmt.Sprintf("Unrecognized command: %v", requests.Message)
		fmt.Println(errMsg)
		return errors.New(errMsg)
	}
	for {
		data := <-revMsg
		err := response.Send(&kapi.PlayReply{
			Res: data,
		})
		if err != nil {
			return err
		}
		if "exit" == data {
			break
		}
	}
	return nil
}

func (s *server) RunPlaybook(ctx context.Context, requests *kapi.PlayRequests) (*kapi.PlayReply, error) {
	config.GetConfig()
	revMsg := make(chan string)
	result := &kapi.PlayReply{}
	inventory := constant.HostForKubernetes
	installScript := constant.KubernetesInstallScript
	go ansible.RunPlaybook(revMsg, inventory, installScript)
	for {
		data := <-revMsg
		result.Res = data
		if "exit" == data {
			break
		}
	}
	return result, nil
}

func (s *server) GenerateYaml(ctx context.Context, in *kapi.InventoryRequest) (*kapi.InventoryReply, error) {
	err := ansible.GenHosts(in)
	if err != nil {
        fmt.Printf("Error: %v\n", err)
		return &kapi.InventoryReply{Message: "Generate hosts failure "}, err
	}
	err = ansible.GenYamlHost(in)
	if err != nil {
        fmt.Printf("Error: %v\n", err)
		return &kapi.InventoryReply{Message: "Generate yaml failure "}, err
	}
	return &kapi.InventoryReply{Message: "Generate yaml success "}, nil
}

func main() {
	config.GetConfig()
	host := viper.GetString("bind.host")
	port := viper.GetString("bind.port")
	addr := fmt.Sprintf("%s:%s", host, port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	s := grpc.NewServer()
	kapi.RegisterAnsibleServerServer(s, &server{})
	reflection.Register(s)
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
}
