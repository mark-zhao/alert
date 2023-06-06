
package common

import (
	"encoding/xml"
)

type AlertInfo struct{
	Region 		string
	No     		string
	TimeStamp   string
}

type Request struct{
	XMLName xml.Name  `xml:"xml"`
	ToUserName string `xml:"ToUserName"`
	AgentID  string   `xml:"AgentID"`
	Encrypt  string   `xml:"Encrypt"`
}

type Token struct {
	Errcode int 		`json:"errcode"`
	Errmsg  string 		`json:"errmsg"`
	AccessToken string  `json:"access_token"`
	ExpiresIn int64 		`json:"expires_in"`
}

type text struct{
	Content string `json:"content"`
}

type WeChatAlert struct {
	ToUser  string  			`json:"touser"`
	MsgType string  			`json:"msgtype"`
	AgentId int     			`json:"agentid"`
	Text    text  				`json:"text"`
	Safe    int 				`json:"safe"`
}

func NewWeChatAlert(toUser, msgType string, agentId, safe int, content string) *WeChatAlert{
	t := text{
		Content: content,
	}
	return &WeChatAlert{
		ToUser: toUser,
		MsgType: msgType,
		AgentId: agentId,
		Safe: safe,
		Text: t,
	}
}