package web

import (
	"common"
	"context"
	"glog"
	"io/ioutil"
	"net/http"
	"rpc"
	"strings"
	"time"
	"timer"
)

const uri = "/wechat/ceph/alert"
var upCh = make(chan int64, 1)
var requestChan = make(chan string)
var resultChan = make(chan string)
var alertTimer = 21600 //300 second
var hbTimer = 60     //60 second


func server(w http.ResponseWriter, r *http.Request) {
	requestMap := r.URL.Query()
	msgSignatrue, msgSignatrueExist := requestMap["msg_signature"]
	timeStamp, timeStampExist := requestMap["timestamp"]
	nonce, nonceExist := requestMap["nonce"]
	if !msgSignatrueExist || !timeStampExist || !nonceExist {
		glog.Error("request data missing; message: msg_signatrue exist: ", msgSignatrue, "; timestamp exist: ", timeStamp,
			", nonce exist:", nonceExist)
		http.Error(w, "test request method", 200)
		return
	}
	if common.ValidateUrl(strings.Join(timeStamp, " "), strings.Join(nonce, " "), strings.Join(msgSignatrue, " ")) {
		glog.Error("Invalid request url")
		http.Error(w, "Invalid request", 405)
		return
	}

	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/xml")

		originData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			glog.Error("Invalid request; ", err)
			http.Error(w, "Invalid request", 405)
			return
		}
		//glog.Info("content: ", originData)
		rawData, err := common.ParseReq(originData)
		if err != nil {
			glog.Error("parse request failed;")
			http.Error(w, "parse data failed;", 500)
			return
		}

		glog.Info("self data:", rawData.Content)
		result := rpc.Server(rawData.Content, rawData.FromUserName, upCh)

		msg, err := common.EncryptData(rawData.ToUserName, rawData.FromUserName, rawData.MsgType, result)
		if err != nil {
			glog.Error("encrypt data failed; ", err)
			http.Error(w, "test request method", 200)
			return
		}

		glog.Info("msg: ", string(msg))
		m_size, err := w.Write(msg)
		glog.Info("w.Write return: m_size: ", m_size, "\n, err:", err)
		if err != nil{
			glog.Errorf("passive replay message failed; due to: %s\n request instruction: %s", err.Error(), rawData.Content )
			return
		}

	} else if r.Method == "GET" {
		echoStr, _ := requestMap["echostr"]

		glog.Info("msg_signatrue: ", msgSignatrue, "; timestamp: ", timeStamp, "; nonce:", nonce, "; echostr:", echoStr)
		msg, err := common.Decrypt(strings.Join(echoStr, " "))
		if err != nil {
			glog.Error("decrypt failed; message: ", err)
			return
		}
		w.Write(msg)
	} else {
		glog.Error("Invalid request method.", 200)
		http.Error(w, "Invalid request method", 200)
	}

}

func StartHTTP(ctx context.Context) error {
	http.HandleFunc(uri, server)
	server := &http.Server{
		Addr:         "127.0.0.1:8080",
		Handler:      http.DefaultServeMux,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	go func(s *http.Server) {
		if err := server.ListenAndServe(); err != nil {
			glog.Error("http server listen and Serve failed; message: ", err)
		}
	}(server)
	glog.Info("http server launch sucess!")

	//hbExitCh := make(chan struct{}, 1)
	hb := timer.NewHeartBeat(int64(hbTimer), rpc.HeartBeatTask) // 5 second heartbeat
	go hb.Run()
	glog.Info("Heart Beat set sucess!!!")

	timerCheckStatus2 := timer.NewTimerStatus(int64(hbTimer), upCh, rpc.CephStatusTask_2)
	go timerCheckStatus2.Run()
	glog.Info("timer for check status set ok")

	timerCheckStatus := timer.NewTimerStatus(int64(alertTimer), upCh, rpc.CephStatusTask)
	go timerCheckStatus.Run()
	glog.Info("timer for check status set ok")

	for {
		select {
		case <-ctx.Done():
			server.Close()
			hb.Cancel()
			timerCheckStatus.Cancel()
			return nil
		}
	}
}
