package ansible

import (
	"encoding/json"
	"fmt"
	file "github.com/mensylisir/KAnsible/common"
	"github.com/mensylisir/KAnsible/config"
	"github.com/mensylisir/KAnsible/constant"
	"github.com/mensylisir/KAnsible/kapi"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strconv"
	"strings"
)

type JsonStruct struct {
}

func (jst *JsonStruct) LoadFile(filename string, v interface{}) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}

func (jst *JsonStruct) Load(data []byte, v interface{}) {
	err := json.Unmarshal(data, v)
	if err != nil {
		return
	}
}

func GenNode(nodes []*kapi.Node) map[string]interface{} {
	ansibleHosts := make(map[string]interface{})
	for _, node := range nodes {
		ansibleNode := make(map[string]interface{})
		ansibleNode["ansible_host"] = node.Ip
		ansibleNode["ansible_port"], _ = strconv.Atoi(node.Port)
		ansibleNode["ip"] = node.Ip
		ansibleNode["access_ip"] = node.Ip
		ansibleHosts[node.Name] = ansibleNode
	}

	return ansibleHosts
}

func GenChildren(nodes []*kapi.Node) map[string]interface{} {

	var t interface{} = nil

	mHosts := make(map[string]interface{})
	wHosts := make(map[string]interface{})
	for _, node := range nodes {
		if "master" == strings.ToLower(node.Role) {
			mHosts[node.Name] = t
		}
		wHosts[node.Name] = t
	}

	var kMaster = make(map[string]interface{})
	kMaster["hosts"] = mHosts

	var kWorker = make(map[string]interface{})
	kWorker["hosts"] = wHosts

	var childMaster = make(map[string]interface{})
	childMaster["kube_control_plane"] = t
	childMaster["kube_node"] = t
	var clusterChild = make(map[string]interface{})
	clusterChild["children"] = childMaster

	var calico = make(map[string]interface{})
	calico["hosts"] = make(map[string]interface{})

	var children = make(map[string]interface{})
	children["kube_control_plane"] = kMaster
	children["kube_node"] = kWorker
	children["etcd"] = kMaster
	children["k8s_cluster"] = clusterChild
	children["calico_rr"] = calico
	return children
}

func Write2Yaml(content map[string]interface{}, inventory string) {
	res, err := yaml.Marshal(content)
	if err != nil {
		errMsg := fmt.Sprintf("Error: %v", err)
		fmt.Println(errMsg)
		_ = file.WriteLog(errMsg)
		return
	}
	res2 := strings.Replace(string(res), "null", "", -1)
	fmt.Println(res2)
	err = ioutil.WriteFile(inventory, []byte(res2), 0777)
	if err != nil {
		errMsg := fmt.Sprintf("ioutil.WriteFile failure, err=[%v]\n", err)
		fmt.Println(errMsg)
		_ = file.WriteLog(errMsg)
		return
	}
}

func GenYamlHost(request *kapi.InventoryRequest) error {
	config.GetConfig()
	inventory := constant.HostForKubernetes
	items := request.GetItem()
	node := GenNode(items)
	children := GenChildren(items)
	all := make(map[string]interface{})
	all["hosts"] = node
	all["children"] = children
	data := make(map[string]interface{})
	data["all"] = all
	Write2Yaml(data, inventory)
	return nil
}

func GenerateYamlHost(request *kapi.PlaybookRequests) error {
	config.GetConfig()
	inventory := constant.HostForKubernetes
	items := request.GetItem()
	node := GenNode(items)
	children := GenChildren(items)
	all := make(map[string]interface{})
	all["hosts"] = node
	all["children"] = children
	data := make(map[string]interface{})
	data["all"] = all
	Write2Yaml(data, inventory)
	return nil
}

func GenerateHosts(request *kapi.PlaybookRequests) error {
	ansibleHosts := constant.AnsibleHosts
	items := request.GetItem()
	fmt.Println(items)
	for _, item := range items {
		host := ""
		host += item.Ip + " "
		host += "ansible_ssh_port=" + item.Port + " "
		host += "ansible_ssh_user=root" + " "
		host += "ansible_ssh_pass=" + item.Password
		err := file.Append(host, ansibleHosts)
		if err != nil {
			errMsg := fmt.Sprintf("Error: %v", err)
			fmt.Println(errMsg)
			_ = file.WriteLog(errMsg)
			return err
		}
	}
	return nil
}

func DeleteHosts(request *kapi.PlaybookRequests) error {
	ansibleHosts := constant.AnsibleHosts
	items := request.GetItem()
	for _, item := range items {
		host := ""
		host += item.Ip + " "
		host += "ansible_ssh_port=" + item.Port + " "
		host += "ansible_ssh_user=root" + " "
		host += "ansible_ssh_pass=" + item.Password
		err := file.DeleteLine(host, ansibleHosts)
		if err != nil {
			errMsg := fmt.Sprintf("Error: %v", err)
			fmt.Println(errMsg)
			_ = file.WriteLog(errMsg)
			return err
		}
	}
	return nil
}

func GenHosts(request *kapi.InventoryRequest) error {
	ansibleHosts := constant.AnsibleHosts
	items := request.GetItem()
	for _, item := range items {
		host := ""
		host += item.Ip + " "
		host += "ansible_ssh_port=" + item.Port + " "
		host += "ansible_ssh_user=root" + " "
		host += "ansible_ssh_pass=" + item.Password
		err := file.Append(host, ansibleHosts)
		if err != nil {
			errMsg := fmt.Sprintf("Error: %v", err)
			fmt.Println(errMsg)
			_ = file.WriteLog(errMsg)
			return err
		}
	}
	return nil
}

func WriteConfig(request *kapi.ConfigRequest) error {
	if cluster := request.GetClusterName(); cluster != "" {
		items := make(map[string]string)
		backupEtcdVars := constant.BackupEtcdVars
		items["cluster_name"] = cluster
		err := config.WriteConfig(backupEtcdVars, items)
		if err != nil {
			fmt.Printf("error: %v", err)
			return err
		}
		restoreEtcdVars := constant.RestoreEtcdVars
		err = config.WriteConfig(restoreEtcdVars, items)
		if err != nil {
			fmt.Printf("error: %v", err)
			return err
		}
	}
	privisionerName := request.GetNfsProvisionerName()
	nfsServer := request.GetNfsServer()
	nfsPath := request.GetNfsServerPath()
	if privisionerName != "" && nfsServer != "" && nfsPath != "" {
		items := make(map[string]string)
		nfsClusterVars := constant.NfsClusterVars
		items["storage_nfs_provisioner_name"] = privisionerName
		items["storage_nfs_server"] = nfsServer
		items["storage_nfs_server_path"] = nfsPath
		err := config.WriteConfig(nfsClusterVars, items)
		if err != nil {
			fmt.Printf("error: %v", err)
			return err
		}
	}
	if kubeVersion := request.GetKubeVersion(); kubeVersion != "" {
		kubernetesClusterVars := constant.KubernetesClusterVars
		items := make(map[string]string)
		items["kube_version"] = kubeVersion
		err := config.WriteConfig(kubernetesClusterVars, items)
		if err != nil {
			fmt.Printf("error: %v", err)
			return err
		}
	}
	if containerNetwork := request.GetContainerNetwork(); containerNetwork != "" {
		err := WriteNetworkConfig(containerNetwork)
		if err != nil {
			return err
		}
	}
	if networkMode := request.GetNetworkMode(); networkMode != "" {
		err := WriteNetworkConfig(networkMode)
		if err != nil {
			return err
		}
	}
	return nil
}

func WriteNetworkConfig(config string) error {
	return nil
}

func WriteConfig2(request *kapi.Config) error {
	if cluster := request.GetClusterName(); cluster != "" {
		items := make(map[string]string)
		backupEtcdVars := constant.BackupEtcdVars
		items["cluster_name"] = cluster
		err := config.WriteConfig(backupEtcdVars, items)
		if err != nil {
			fmt.Printf("error: %v", err)
			return err
		}
		restoreEtcdVars := constant.RestoreEtcdVars
		err = config.WriteConfig(restoreEtcdVars, items)
		if err != nil {
			fmt.Printf("error: %v", err)
			return err
		}
	}
	privisionerName := request.GetNfsProvisionerName()
	nfsServer := request.GetNfsServer()
	nfsPath := request.GetNfsServerPath()
	if privisionerName != "" && nfsServer != "" && nfsPath != "" {
		items := make(map[string]string)
		nfsClusterVars := constant.NfsClusterVars
		items["storage_nfs_provisioner_name"] = privisionerName
		items["storage_nfs_server"] = nfsServer
		items["storage_nfs_server_path"] = nfsPath
		err := config.WriteConfig(nfsClusterVars, items)
		if err != nil {
			fmt.Printf("error: %v", err)
			return err
		}
	}
	//if kubeVersion := request.GetKubeVersion(); kubeVersion != "" {
	//	kubernetesClusterVars := constant.KubernetesClusterVars
	//	items := make(map[string]string)
	//	items["kube_version"] = kubeVersion
	//	err := config.WriteConfig(kubernetesClusterVars, items)
	//	if err != nil {
	//		fmt.Printf("error: %v", err)
	//		return err
	//	}
	//}
	//if containerNetwork := request.GetContainerNetwork();containerNetwork != "" {
	//	err := WriteNetworkConfig(containerNetwork)
	//	if err != nil {
	//		return err
	//	}
	//}
	//if networkMode := request.GetNetworkMode();networkMode != "" {
	//	err := WriteNetworkConfig(networkMode)
	//	if err != nil {
	//		return err
	//	}
	//}
	return nil
}
