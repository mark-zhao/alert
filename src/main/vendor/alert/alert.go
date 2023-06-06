package alert

import (
	"bytes"
	"common"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/buger/jsonparser"
	"glog"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
	"time"
)

const tokenUrl = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"
const msgUrl = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s"
const weChatAgentId = 1000002
const corpid = "ww1a67bf2e79b10c3b"
const secret = "affjMBqrftYBjKITh1zKxa6tpUg7wnAeVlEwB47EOQs"

type WeChatAlert struct {
	corpid    string
	secret    string
	token     string
	expire    int64
	timeStamp time.Time
}

func NewWeChatAlert() *WeChatAlert {
	return &WeChatAlert{
		corpid: corpid,
		secret: secret,
	}
}

func (wc *WeChatAlert) getToken() (err error) {
	url := fmt.Sprintf(tokenUrl, wc.corpid, wc.secret)
	ctx, _ := context.WithTimeout(context.Background(), 16*time.Second)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		glog.Error("NewRequest failed; msg: ", err)
		return
	}
	request = request.WithContext(ctx)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		glog.Error("do request failed; mes: ", err)
		return
	}
	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		return
	}
	content, _ := ioutil.ReadAll(resp.Body)
	var token common.Token
	if err = json.Unmarshal(content, &token); err != nil {
		glog.Error("json Unmarshal failed; msg:", err)
		return
	}
	wc.timeStamp = time.Now()
	wc.token = token.AccessToken
	wc.expire = token.ExpiresIn
	return
}

func (wc *WeChatAlert) Alert(msg string) (err error) {
	if time.Now().Unix()-wc.timeStamp.Unix() > wc.expire {
		if err := wc.getToken(); err != nil {
			glog.Error("weChat getToken failed; due to:", err)
			return err
		}
	}
	url := fmt.Sprintf(msgUrl, wc.token)

	alert := common.NewWeChatAlert("@all", "text", weChatAgentId, 0, msg)
	params, _ := json.Marshal(alert)
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(params))
	if err != nil {
		glog.Error("NewRequest failed; msg: ", err)
		return
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request = request.WithContext(ctx)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		glog.Error("do request failed; mes: ", err)
		return
	}
	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		return
	}

	content, _ := ioutil.ReadAll(resp.Body)
	errcode, err := jsonparser.GetInt(content, "errcode")

	if errcode != 0 || err != nil {
		errmsg, _ := jsonparser.GetString(content, "errmsg")
		glog.Error("weChat alert response; token:", wc.token, "; errmsg: ", errmsg)
		err = errors.New(errmsg)
		return
	}
	return
}
