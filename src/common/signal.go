package common

import (
	"context"
	"errors"
	"fmt"
	"glog"
	"os"
	"os/signal"
	"strconv"
)

const pidPath = "/var/run/ceph/"

func SetSignal(cancelFunc context.CancelFunc) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill)
	glog.Info("set signal success")
	sinalId := <-interrupt
	glog.Info("receive signal for exit; signal: ", sinalId)
	cancelFunc()
}

func savePid(pid string) error {
	if !pathExist(pidPath) {
		if err := os.Mkdir(pidPath, 0644); err != nil{
			glog.Error("mkdir pidPath failed; ", err)
			return err
		}
	}
	pidFile := pidPath + pid
	if pathExist(pidFile){
		fmt.Printf("file %s exist; process running now\n", pidFile)
		return errors.New("process running already")
	}
	ppid := os.Getpid()
	fd, err := os.Create(pidFile)
	if err != nil{
		glog.Error("write pid file failed; message: ", err)
	}
	defer fd.Close()
	fd.Write([]byte(strconv.Itoa(ppid)))
	return err
}

func pathExist(path string ) bool {
	_, err := os.Stat(path)
	if err != nil{
		if os.IsExist(err){
			return true
		}
		return false
	}
	return true
}
