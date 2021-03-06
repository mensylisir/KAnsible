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
	case constant.BACKUP_ETCD:
		go ansible.BackupEtcd(revMsg)
	case constant.RESTORE_ETCD:
		go ansible.RestoreEtcd(revMsg)
	case constant.CREATE_NFS:
		go ansible.CreateNFS(revMsg)
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
		if "exit" == data || "failure" == data || "success" == data {
			break
		}
	}
	return nil
}

func (s *server) StreamPlaybook(requests *kapi.PlaybookRequests, response kapi.AnsibleServer_StreamPlaybookServer) error {
	config.GetConfig()
	err := ansible.GenerateHosts(requests)
	if err != nil {
		errMsg := fmt.Sprintf("Error: %v\n", err)
		err := response.Send(&kapi.PlayReply{
			Res: errMsg,
		})
		return err
	}
	err = ansible.GenerateYamlHost(requests)
	if err != nil {
		errMsg := fmt.Sprintf("Error: %v\n", err)
		err := response.Send(&kapi.PlayReply{
			Res: errMsg,
		})
		return err
	}
	err = ansible.WriteConfig2(requests.Config)
	if err != nil {
		errMsg := fmt.Sprintf("Error: %v\n", err)
		err := response.Send(&kapi.PlayReply{
			Res: errMsg,
		})
		return err
	}

	revMsg := make(chan string)
	go ansible.DistributePublicKey(revMsg)
	for {
		data := <-revMsg
		if "success" == data {
			break
		}
		err := response.Send(&kapi.PlayReply{
			Res: data,
		})
		if err != nil {
			return err
		}
	}
	switch requests.GetAction() {
	case constant.INSTALL_ACTION:
		go ansible.InstallKubernetes(revMsg)
	case constant.UPDATE_ACTION:
		go ansible.InstallKubernetes(revMsg)
	case constant.RESET_ACTION:
		go ansible.ResetKubernetes(revMsg)
	case constant.DISTRIBUTE_KEY:
		go ansible.DistributePublicKey(revMsg)
	case constant.BACKUP_ETCD:
		go ansible.BackupEtcd(revMsg)
	case constant.RESTORE_ETCD:
		go ansible.RestoreEtcd(revMsg)
	case constant.CREATE_NFS:
		go ansible.CreateNFS(revMsg)
	case constant.ADD_NODE:
		limit := requests.GetVars()[0]
		go ansible.AddNode(revMsg, limit)
	case constant.REMOVE_NODE:
		node := requests.GetVars()[0]
		go ansible.RemoveNode(revMsg, node)
	default:
		errMsg := fmt.Sprintf("Unrecognized command: %v", requests.Action)
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
		if "exit" == data || "failure" == data || "success" == data {
			err = ansible.DeleteHosts(requests)
			if err != nil {
				errMsg := fmt.Sprintf("Error: %v\n", err)
				err := response.Send(&kapi.PlayReply{
					Res: errMsg,
				})
				return err
			}
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
	go ansible.RunPlaybook(revMsg, inventory, installScript, nil)
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

func (s *server) CheckConfiguration(ctx context.Context, in *kapi.ConfigRequest) (*kapi.ConfigReply, error) {
	err := ansible.WriteConfig(in)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return &kapi.ConfigReply{Message: "Write config failure "}, err
	}
	return &kapi.ConfigReply{Message: "Write config success "}, nil
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
