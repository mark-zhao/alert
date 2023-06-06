package rpc

import (
	"fmt"
	"glog"
	"time"
)

const timeInstruction = "ceph health detail"

func CephStatusTask(status map[string]string){
	var alertMsg string
	for region, addr := range infos {
		result, err := rpcClusterStatus(region, addr, timeInstruction)
		if err != nil{
			glog.Errorf("region: %s timerTask failed; \n due to: %s", region, err.Error())
		}

		alertMsg = fmt.Sprintf("region: %s\ntimeStamp: %s\nstatus: %s", region, time.Now().Format("2006-01-02 15:04:05"), result)
		glog.Info(alertMsg)
		status[region] = alertMsg
		/*
		if !strings.Contains(r.GetResult(), normalState) {
			//glog.Infof("region: %s\n timeStamp: %s;\n status: \n %s ", region, time.Now().Format("2006-01-02 15:04:05"), r.GetResult())

		}*/
	}
}
