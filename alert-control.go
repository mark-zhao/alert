package main

import (
	"common"
	"context"
	"glog"
	//"github.com/golang/glog"
	"web"
)

//init
func init() {
	//@TODO
}

func main() {
	defer glog.Flush()
	//flag.Parse()

	glog.Info("alert control start!")
	ctx, cancel := context.WithCancel(context.Background())
	//set signal
	go common.SetSignal( cancel)

	if err := web.StartHTTP(ctx); err != nil {
		glog.Error("start http service failed; ", err)
	}
}
