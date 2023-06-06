package rpc

import (
	"fmt"
	"google.golang.org/grpc"
	"glog"
	"context"
	"pb"
	"strings"
	"time"
)

func HeartBeatTask() (alertMsg string) {
	for region, addr := range infos {
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			glog.Error("grpc.Dial failed; ", err, "; region:", region, "; addr:", addr)
			alertMsg = alertMsg + fmt.Sprintf("region: %s HeartBeatTask grpc dail failed; due to %s \n", region, err.Error())
			continue
		}
		defer conn.Close()

		c := pb.NewCephServiceClient(conn)

		r, err := c.HeartBeat(context.Background(), &pb.CephHeartBeatRequest{Ping: "ping"})
		if err != nil {
			glog.Error("rpc call HeartBeat failed;", err)
			//hb missing and send message to client;
			alertMsg = alertMsg + fmt.Sprintf("region: %s HeartBeatTask request; due to %s \n", region, err.Error())
		}else {
			glog.Infof("heartBeat region: %s , heartBeat result: %s", region, r.Pong)
		}
	}
	return
}

func CephStatusTask_2(status map[string]string){
	var alertMsg string
	for region, addr := range infos {
		result, err := rpcClusterStatus(region, addr, timeInstruction)
		if err != nil{
			glog.Errorf("region: %s timerTask failed; \n due to: %s", region, err.Error())
		}
		if !strings.Contains(result,"HEALTH_OK") {
			alertMsg = fmt.Sprintf("region: %s\ntimeStamp: %s\nstatus: %s", region, time.Now().Format("2006-01-02 15:04:05"), result)
			glog.Info(alertMsg)
			status[region] = alertMsg
		}
		/*
			if !strings.Contains(r.GetResult(), normalState) {
				//glog.Infof("region: %s\n timeStamp: %s;\n status: \n %s ", region, time.Now().Format("2006-01-02 15:04:05"), r.GetResult())

			}*/
	}
}