package common

import (
	"KAnsible/constant"
	"fmt"
	"io"
	"os"
	"time"
)

var path = constant.LogPath + "/" + time.Now().Format(constant.DateFormat) + "/"
var fileName = time.Now().Format(constant.TimeFormat) + ".log"

func WriteLog(msg string) error {
	if !IsExist(path) {
		err := CreateDir(path)
		if err != nil {
			fmt.Println("Create directory failure")
		}
	}
	f, err := os.OpenFile(path + fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	_, err = io.WriteString(f, constant.LineFeed+msg)
	if err != nil {
		return err
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	return nil
}

func CreateDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	err = os.Chmod(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}
