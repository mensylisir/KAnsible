package ansible

import (
	"KAnsible/constant"
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
