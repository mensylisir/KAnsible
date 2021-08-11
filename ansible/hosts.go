package ansible

import (
	"github.com/mensylisir/KAnsible/api"
	file "github.com/mensylisir/KAnsible/common"
	"github.com/mensylisir/KAnsible/config"
	"github.com/mensylisir/KAnsible/constant"
	"encoding/json"
	"fmt"
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

func GenNode(nodes []*api.Node) map[string]interface{} {
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

func GenChildren(nodes []*api.Node) map[string]interface{} {

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
	childMaster["kube-master"] = t
	childMaster["kube-node"] = t
	var clusterChild = make(map[string]interface{})
	clusterChild["children"] = childMaster

	var calico = make(map[string]interface{})
	calico["hosts"] = make(map[string]interface{})

	var children = make(map[string]interface{})
	children["kube-master"] = kMaster
	children["kube-node"] = kWorker
	children["etcd"] = kMaster
	children["k8s-cluster"] = clusterChild
	children["calico-rr"] = calico
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

func GenYamlHost(request *api.InventoryRequest) error {
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

func GenHosts(request *api.InventoryRequest) error {
	config.GetConfig()
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
