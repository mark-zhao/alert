package rpc

import (
	"fmt"
	"google.golang.org/grpc"
	"glog"
	"context"
	"pb"
)

func HeartBeatTask() (alertMsg string) {
	for region, addr := range infos {
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			glog.Error("grpc.Dial failed; ", err, "; region:", region, "; addr:", addr)
			alertMsg = fmt.Sprintf("region: %s HeartBeatTask grpc dail failed; due to %s \n", region, err.Error())
		}
		defer conn.Close()

		c := pb.NewCephServiceClient(conn)

		r, err := c.HeartBeat(context.Background(), &pb.CephHeartBeatRequest{Ping: "ping"})
		if err != nil {
			glog.Error("rpc call HeartBeat failed;", err)
			//hb missing and send message to client;
			alertMsg = fmt.Sprintf("region: %s HeartBeatTask request; due to %s \n", region, err.Error())
		}else {
			glog.Infof("heartBeat region: %s , heartBeat result: %s", region, r.Pong)
		}
	}
	return
}
