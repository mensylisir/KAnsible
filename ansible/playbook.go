package ansible

import (
	"fmt"
	util "github.com/mensylisir/KAnsible/common"
	"github.com/mensylisir/KAnsible/constant"
)

func RunPlaybook(revMsg chan string, inventory, script string, vars map[string]string) bool {
	ansiblePlaybookPath, err := constant.LookUpAnsiblePlaybookBinPath()
	if err != nil {
		errMsg := fmt.Sprintf("Cannot find ansibleplaybook: %v\n", err)
		fmt.Printf(errMsg)
		_ = util.WriteLog(errMsg)
		return false
	}
	command := []string{ansiblePlaybookPath}
	command = append(command, script)
	if inventory != "" {
		command = append(command, "-i")
		command = append(command, inventory)
		command = append(command, "--become")
		command = append(command, "--become-user=root")
	}
	for key, value := range vars {
		if key == "limit" {
			command = append(command, "--limit")
			command = append(command, value)
		}
		if key == "remove_node" {
			value = fmt.Sprintf("node=%s", value)
			command = append(command, "-e")
			command = append(command, value)
		}

	}
	fmt.Println(command)
	err = util.RunCMD(revMsg, command[0], command[1:]...)
	if err != nil {
		errMsg := fmt.Sprintf("Error :%v", err)
		fmt.Println(errMsg)
		_ = util.WriteLog(errMsg)
		return false
	}
	return true
}
