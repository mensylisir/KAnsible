package ansible

import (
	"github.com/mensylisir/KAnsible/constant"
)

func InstallKubernetes(revMsg chan string) bool {
	inventory := constant.HostForKubernetes
	installScript := constant.KubernetesInstallScript
	ok := RunPlaybook(revMsg, inventory, installScript, nil)
	if ok {
		return true
	}
	return false
}

func ResetKubernetes(revMsg chan string) bool {
	inventory := constant.HostForKubernetes
	resetScript := constant.KubernetesResetScript
	ok := RunPlaybook(revMsg, inventory, resetScript, nil)
	if ok {
		return true
	}
	return false
}

func DistributePublicKey(revMsg chan string) bool {
	authorizeScript := constant.AuthorizeKeysScript
	ok := RunPlaybook(revMsg, "", authorizeScript, nil)
	if ok {
		return true
	}
	return false
}

func BackupEtcd(revMsg chan string) bool {
	inventory := constant.HostForKubernetes
	backupEtcdScript := constant.BackupEtcdScript
	ok := RunPlaybook(revMsg, inventory, backupEtcdScript, nil)
	if ok {
		return true
	}
	return false
}

func RestoreEtcd(revMsg chan string) bool {
	inventory := constant.HostForKubernetes
	restoreEtcdScript := constant.RestorEtcdScript
	ok := RunPlaybook(revMsg, inventory, restoreEtcdScript, nil)
	if ok {
		return true
	}
	return false
}

func CreateNFS(revMsg chan string) bool {
	inventory := constant.HostForKubernetes
	createNFSScript := constant.CreateNFSScript
	ok := RunPlaybook(revMsg, inventory, createNFSScript, nil)
	if ok {
		return true
	}
	return false
}

func AddNode(revMsg chan string, limit string) bool {
	inventory := constant.HostForKubernetes
	addnodeScript := constant.KubernetesAddNodeScript
	vars := make(map[string]string)
	vars["limit"] = limit
	ok := RunPlaybook(revMsg, inventory, addnodeScript, vars)
	if ok {
		return true
	}
	return false
}

func RemoveNode(revMsg chan string, node string) bool {
	inventory := constant.HostForKubernetes
	removenodeScript := constant.KubernetesRemoveNodeScript
	vars := make(map[string]string)
	vars["remove_node"] = node
	ok := RunPlaybook(revMsg, inventory, removenodeScript, vars)
	if ok {
		return true
	}
	return false
}
