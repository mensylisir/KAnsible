package common

import (
	"bufio"
	"errors"
	"fmt"
	"os/exec"
)

func RunCommand(revMsg chan string, name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout

	if err != nil {
		return err
	}

	if err = cmd.Start(); err != nil {
		return err
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		if err != nil {
			return err
		}
		fmt.Print(string(tmp))
		revMsg <- string(tmp)
		err = WriteLog(string(tmp))
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		return err
	}
	revMsg <- "exit"
	return nil
}

func RunCMD(revMsg chan string, name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	stderr, _ := cmd.StderrPipe()
	stdout, _ := cmd.StdoutPipe()
	done := make(chan bool)
	err := cmd.Start()
	if err != nil {
		return err
	}
	go func() {
		errScanner := bufio.NewScanner(stderr)
		errScanner.Split(bufio.ScanLines)
		for errScanner.Scan() {
			select {
			case <-done:
				revMsg <- "done"
				return
			default:
				m := errScanner.Text()
				err := WriteLog(m)
				if err != nil {
					return
				}
				revMsg <- m
				fmt.Println(m)
			}

		}
	}()

	go func() {
		stdScanner := bufio.NewScanner(stdout)
		stdScanner.Split(bufio.ScanLines)
		for stdScanner.Scan() {
			select {
			case <-done:
				revMsg <- "done"
				return
			default:
				m := stdScanner.Text()
				err := WriteLog(m)
				if err != nil {
					return
				}
				revMsg <- m
				fmt.Println(m)
			}
		}
	}()
	err = cmd.Wait()
	if !cmd.ProcessState.Success() {
		revMsg <- "failure"
		errMsg := fmt.Sprintf("Command execution failed")
		_ = WriteLog(errMsg)
		return errors.New(errMsg)
	}
	revMsg <- "success"
	_ = WriteLog("Command execution success")
	return nil
}
