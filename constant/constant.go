package constant

import (
	"log"
	"os/exec"
	"path"
	"path/filepath"
)

const (
	AnsiblePlaybookBinPath = "ansible-playbook"
	AnsibleBinPath         = "ansible"
)

var (
	AnsibleHosts            = path.Join("/", "etc", "ansible", "hosts")
	AuthorizeKeysScript     = path.Join("data", "kubespray", "kubespray", "authorized_keys.yaml")
	KubernetesInstallScript = path.Join("data", "kubespray", "kubespray", "cluster.yml")
	KubernetesUpdateScript  = path.Join("data", "kubespray", "kubespray", "cluster.yml")
	KubernetesResetScript   = path.Join("data", "kubespray", "kubespray", "reset.yml")
	BackupEtcdScript        = path.Join("data", "kubespray", "kubespray", "etcdbackup.yml")
	RestorEtcdScript        = path.Join("data", "kubespray", "kubespray", "etcdrestore.yml")
	CreateNFSScript         = path.Join("data", "kubespray", "kubespray", "nfs.yml")
	HostForKubernetes       = path.Join("data", "kubespray", "kubespray", "inventory", "mycluster", "hosts.yaml")
)

var (
	BackupEtcdVars        = path.Join("data", "kubespray", "kubespray", "roles", "etcdbackup", "vars", "main.yml")
	RestoreEtcdVars       = path.Join("data", "kubespray", "kubespray", "roles", "etcdrestore", "vars", "main.yml")
	KubernetesClusterVars = path.Join("data", "kubespray", "kubespray", "inventory", "mycluster", "group_vars", "k8s_cluster", "k8s-cluster.yml")
	NfsClusterVars        = path.Join("data", "kubespray", "kubespray", "roles", "cluster-storage", "nfs", "vars", "main.yml")
)

var (
	LogPath    = path.Join("data", "log")
	DateFormat = "2006-01-02"
	TimeFormat = "20060102150405"
	LineFeed   = "\r\n"
)

var (
	DISTRIBUTE_KEY = "distribute"
	INSTALL_ACTION = "install"
	UPDATE_ACTION  = "update"
	RESET_ACTION   = "reset"
	BACKUP_ETCD    = "etcdbackup"
	RESTORE_ETCD   = "etcdrestore"
	CREATE_NFS     = "nfscreate"
	ADD_NODE       = "addnode"
)

func GetBinPath(binFile string) (string, error) {
	file, err := exec.LookPath(binFile)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return "", err
	}
	return path, nil
}

func LookUpAnsiblePlaybookBinPath() (string, error) {
	return GetBinPath(AnsiblePlaybookBinPath)
}

func LookUpAnsibleBinPath() (string, error) {
	return GetBinPath(AnsibleBinPath)
}
