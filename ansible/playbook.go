package ansible

import (
	util "github.com/mensylisir/KAnsible/common"
	"github.com/mensylisir/KAnsible/constant"
	"fmt"
)

func RunPlaybook(revMsg chan string, inventory, script string) bool {
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
