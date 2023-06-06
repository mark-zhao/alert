package pb

import (
	"common"
	"context"
	"errors"
	"glog"
)

type Ceph struct {
}

func (c *Ceph) GetClusterStatus(ctx context.Context, req *CephRequest) (*CephResponse, error) {
	r, err := common.RunShell(req.Instruction)
	glog.Info("process instruction: ", req.Instruction)
	if err != nil {
		glog.Errorf("cmd run failed; instruction: %s, result: %s", req.Instruction, r)
	}
	return &CephResponse{Result: r}, err
}

func (c *Ceph) HeartBeat(ctx context.Context, req *CephHeartBeatRequest) (*CephHeartBeatResponse, error) {
	if req.Ping != "ping" {
		glog.Errorf("HeartBeat invalid request")
		return &CephHeartBeatResponse{Pong: "invalid request"}, errors.New("invalid params")
	}
	return &CephHeartBeatResponse{Pong: "pong"}, nil
}
