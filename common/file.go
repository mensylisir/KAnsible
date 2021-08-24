package common

import (
	"bufio"
	"fmt"
	"io"
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

func ModifyMultiConfig(path string, items map[string]string) error {
	for key, value := range items {
		err := ModifyConfig(path, key, value)
		if err != nil {
			errMsg := fmt.Sprintf("Write config error[%v: %v = %v]: %v", path, key, value, err)
			_ = WriteLog(errMsg)
			return err
		}
	}
	return nil
}

func ModifyConfig(path, key, value string) error{
	file, err:= os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		errMsg := fmt.Sprintf("open config file[%v]: %v", path, err)
		_ = WriteLog(errMsg)
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	buf := bufio.NewReader(file)
	output := make([]byte, 0)
	exist := false
	for {
		line, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		if strings.HasPrefix(string(line), "#") {
			continue
		}
		if strings.Contains(string(line), key){
			exist = true
			newline := key + ": " + "\"" + value + "\""
			line = []byte(newline)
		}
		output = append(output, line...)
		output = append(output, []byte("\n")...)
	}
	if !exist {
		newline := key + ": " + "\"" + value + "\""
		line := []byte(newline)
		output = append(output, line...)
		output = append(output, []byte("\n")...)
	}

	if err := writeToFile(path ,output);err != nil{
		errMsg := fmt.Sprintf("write config file[%v]: %v", path, err)
		_ = WriteLog(errMsg)
		return err
	}
	return nil
}

func writeToFile(filePath string, outPut []byte) error {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0600)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	if err != nil {
		errMsg := fmt.Sprintf("open config file[%v]: %v", path, err)
		_ = WriteLog(errMsg)
		return err
	}
	writer := bufio.NewWriter(f)
	_, err = writer.Write(outPut)
	if err != nil {
		return err
	}
	err = writer.Flush()
	if err != nil {
		errMsg := fmt.Sprintf("write config file[%v]: %v", path, err)
		_ = WriteLog(errMsg)
		return err
	}
	return nil
}