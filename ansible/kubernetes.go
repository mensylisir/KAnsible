package ansible

import (
	"github.com/mensylisir/KAnsible/constant"
)

func InstallKubernetes(revMsg chan string) bool {
	inventory := constant.HostForKubernetes
	installScript := constant.KubernetesInstallScript
	ok := RunPlaybook(revMsg, inventory, installScript)
	if ok {
		return true
	}
	return false
}

func ResetKubernetes(revMsg chan string) bool {
	inventory := constant.HostForKubernetes
	resetScript := constant.KubernetesResetScript
	ok := RunPlaybook(revMsg, inventory, resetScript)
	if ok {
		return true
	}
	return false
}

func DistributePublicKey(revMsg chan string) bool {
	authorizeScript := constant.AuthorizeKeysScript
	ok := RunPlaybook(revMsg, "", authorizeScript)
	if ok {
		return true
	}
	return false
}

func BackupEtcd(revMsg chan string) bool {
	inventory := constant.HostForKubernetes
	backupEtcdScript := constant.BackupEtcdScript
	ok := RunPlaybook(revMsg, inventory, backupEtcdScript)
	if ok {
		return true
	}
	return false
}

func RestoreEtcd(revMsg chan string) bool {
	inventory := constant.HostForKubernetes
	restoreEtcdScript := constant.RestorEtcdScript
	ok := RunPlaybook(revMsg, inventory, restoreEtcdScript)
	if ok {
		return true
	}
	return false
}

func CreateNFS(revMsg chan string) bool {
	inventory := constant.HostForKubernetes
	createNFSScript := constant.CreateNFSScript
	ok := RunPlaybook(revMsg, inventory, createNFSScript)
	if ok {
		return true
	}
	return false
}