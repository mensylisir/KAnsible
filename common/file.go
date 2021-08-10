package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func PathExists(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		return err
	}
	return err
}

func Read(path string) ([]byte, error) {
	err := PathExists(path)
	if err != nil {
		errMsg := fmt.Sprintf("File[%v] is not exits: [%v]", path, err)
		_ = WriteLog(errMsg)
		return nil, err
	}
	input, err := ioutil.ReadFile(path)
	if err != nil {
		errMsg := fmt.Sprintf("File[%v] read error: [%v]", string(input), err)
		fmt.Println(errMsg)
		_ = WriteLog(errMsg)
		return nil, err
	}
	return input, nil
}

func ContentExits(host string, content []byte) bool {
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.Contains(line, host) {
			return true
		}
	}
	return false
}

func Append(host, path string) error {
	input, err := Read(path)
	if err != nil {
		errMsg := fmt.Sprintf("Error: %v", err)
		fmt.Println(errMsg)
		_ = WriteLog(errMsg)
		return err
	}
	flag := ContentExits(host, input)
	if !flag {
		lines := strings.Split(string(input), "\n")
		lines = append(lines, host)
		output := strings.Join(lines, "\n")
		err := ioutil.WriteFile(path, []byte(output), 0644)
		if err != nil {
			errMsg := fmt.Sprintf("Error: %v", err)
			fmt.Println(errMsg)
			_ = WriteLog(errMsg)
			return err
		}
	}
	return nil
}
