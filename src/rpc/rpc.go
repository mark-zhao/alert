package rpc

import (
	"context"
	"fmt"
	"github.com/buger/jsonparser"
	"glog"
	"google.golang.org/grpc"
	"io/ioutil"
	"os"
	"pb"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const cephAlertConfigLinux = "/etc/ceph/cephAlert.json"
const cephAlertConfigPathWin = "D:/goproject/src/project/cephAlert.json"
const statusInstruction = "ceph health detail"

var infos = make(map[string]string) //key: region value: part info
var instructionSet = make(map[string]string)
var whiteList = make(map[string]string)

func init() {
	var config string

	if runtime.GOOS == "linux" {
		config = cephAlertConfigLinux
	} else if runtime.GOOS == "windows" {
		config = cephAlertConfigPathWin
	}

	glog.Warningf("config file path:%s", config)
	_, err := os.Stat(config)
	if err != nil && os.IsNotExist(err) {
		glog.Error("config  file stat failed; message:", err)
		os.Exit(1)
	}

	data, err := ioutil.ReadFile(config)
	if err != nil {
		glog.Error("read config file failed; message:", err)
		os.Exit(1)
	}
	region, err := jsonparser.GetString(data, "regions")
	if err != nil {
		glog.Error("jsonparser parse failed; message:", err)
		os.Exit(1)
	}

	regions := strings.Split(region, ";")

	for _, v := range regions {
		addr, err := jsonparser.GetString(data, v, "addr")
		if err != nil {
			glog.Error("jsonparser failed; key:", v, "; message:", err)
			os.Exit(1)
		}
		infos[v] = addr
	}

	err = jsonparser.ObjectEach(data, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		glog.Infof("instructionSet key:%s value:%s Type:%s\n", string(key), string(value), dataType)
		instructionSet[string(key)] = string(value)
		return nil
	}, "instructionSet")

	err = jsonparser.ObjectEach(data, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		glog.Infof("whiteList key:%s value:%s Type:%s\n", string(key), string(value), dataType)
		whiteList[string(key)] = string(value)
		return nil
	}, "whiteList")
	if err != nil {
		glog.Error("json objectEach failed; ", err)
		os.Exit(1)
	}
}

func Server(ch string, username string, upCh chan int64) (result string) {
	if _, ok := whiteList[username]; !ok {
		result = "permission denied"
		return
	}

	regionAndInstructions := strings.Split(ch, ":")
	if regionAndInstructions[0] == "help" {
		result = instructionSet["help"]
		return
	}

	if regionAndInstructions[0] == "status" {
		for region, addr := range infos {
			r, err := rpcClusterStatus(region, addr, statusInstruction)
			if err != nil {
				glog.Errorf("region: %s timerTask failed; \n due to: %s", region, err.Error())
			}
			result = fmt.Sprintf("%s \nregion: %s\ntimestamp: %s \nstatus: %s", result, region, time.Now().Format("2006-01-02 15:04:05"), r)
		}
		return
	}

	if len(regionAndInstructions) != 2 {
		result = "invalid request: request fatal;"
		return
	}

	// add:srcName and from username must be admin
	if regionAndInstructions[0] == "add" && whiteList[username] == "admin" {
		whiteList[regionAndInstructions[1]] = "member"
		result = " authorization " + regionAndInstructions[1] + " success"
		return
	}

	//parse adjust params:
	// adjust:jiaDing@300
	if regionAndInstructions[0] == "adjust" {
		adjustInfo := strings.Split(regionAndInstructions[1], "@")
		if len(adjustInfo) == 2 {
			if _, ok := infos[adjustInfo[0]]; ok {
				if interval, err := strconv.ParseInt(adjustInfo[1], 10, 64); err == nil {
					upCh <- interval
					result = "set alert interval:" + strconv.FormatInt(interval,10)
				} else {
					result = err.Error()
				}
			} else {
				result = "mean to adjust alert interval; but invalid region"
			}
		} else {
			result = "mean to adjust alert interval; but invalid params"
		}
		return
	}

	//regionAndInstructions[0]: region or help or add
	if _, ok := infos[regionAndInstructions[0]]; !ok {
		result = "invalid request: region not found;"
		return
	}

	//regionAndInstructions[1]: command
	if _, ok := instructionSet[regionAndInstructions[1]]; !ok {
		result = "invalid request: instruction not found;"
		return
	}

	result, _ = rpcClusterStatus(regionAndInstructions[0], infos[regionAndInstructions[0]], instructionSet[regionAndInstructions[1]])
	glog.Infof("rpc call return:%s \n", result)
	return
}

func rpcClusterStatus(region, target, instruction string) (result string, err error) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		glog.Error("grpc.Dial failed; ", err, "; region:", region, "; addr:", target)
		return
	}
	defer conn.Close()

	c := pb.NewCephServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	r, err := c.GetClusterStatus(ctx, &pb.CephRequest{Instruction: instruction})
	if err != nil {
		glog.Error("rpc call GetClusterStatus failed;", err)
		result = fmt.Sprintf("rpc call failed; due to %s", err.Error())
		return
	}
	result = r.GetResult()
	return
}
