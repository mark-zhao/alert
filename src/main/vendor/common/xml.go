package common

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/xml"
	"glog"
)

const corpID = "ww1a67bf2e79b10c3b"
const key = "7IT0pmWFpUXIiYXpnv0E3oG9dLexc3NMXRxKigI16pJ"
const token = "PX6TMaLjNyrowakD"

type Req struct{
	XMLName xml.Name `xml:"xml"`
	ToUserName string `xml:"toUserName"`
	AgentID    string `xml:"AgentID"`
	Encrypt    string `xml:"Encrypt"`
}

type Base struct {
	XMLName 		xml.Name 	`xml:"xml"`
	ToUserName  	string		`xml:"ToUserName"`
	FromUserName 	string  	`xml:"FromUserName"`
	CreateTime   	string    	`xml:"CreateTime"`
	MsgType    		string		`xml:"MsgType"`
	Content 		string 		`xml:"Content"`
	MsgId 			string 		`xml:"MsgId"`
	AgentID 		string	 	`xml:"AgentID"`
}

type ResponseBody struct {
	XMLName 		xml.Name 	`xml:"xml"`
	ToUserName  	CDATA 		`xml:"ToUserName"`
	FromUserName 	CDATA  		`xml:"FromUserName"`
	CreateTime   	string    	`xml:"CreateTime"`
	MsgType    		CDATA		`xml:"MsgType"`
	Content 		CDATA 		`xml:"Content"`
}

type EncryptBody struct{
	XMLName 	 xml.Name `xml:"xml"`
	Encrypt 	 CDATA 	  `xml:"Encrypt"`
	MsgSignature CDATA	  `xml:"MsgSignature"`
	TimeStamp    string   `xml:"TimeStamp"`
	Nonce 		 CDATA    `xml:"Nonce"`
 }

type CDATA struct {
	Text string	`xml:",innerxml"`
}

func value2CDATA(v string) CDATA {
	return CDATA{"<![CDATA[" + v + "]]>"}
}

func ParseReq(origin []byte) (*Base, error) {
	req := Req{}
	base := Base{}
	err := xml.Unmarshal(origin, &req)
	if err != nil {
		glog.Error("xml Unmarshal failed; ", err)
		return &base, err
	}

	//Ase Decode
	selfMsg, err := Decrypt(string(req.Encrypt))
	if err != nil {
		glog.Error("decrypt data failed; message:", err)
		return &base, err
	}

	glog.Info("selfMsg: ", string(selfMsg))
	err = xml.Unmarshal(selfMsg, &base)
	if err != nil{
		glog.Error("xml Unmarshal failed; ", err)
		return &base, err
	}
	return &base, nil
}

func EncryptData(toUserName, fromUserName, msgType, content string) (msg []byte, err error){
	//timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	timeStamp := "1539773233"
	encryptXmlData, err := MakeEncryptXmlData(toUserName, fromUserName, msgType, content, timeStamp)
	if err != nil{
		return
	}
	nonce := RandNum(8)
	encryptBody := &EncryptBody{}
	encryptBody.Encrypt = value2CDATA(encryptXmlData)
	encryptBody.Nonce = value2CDATA(string(nonce))
	encryptBody.MsgSignature = value2CDATA(generateSignature(timeStamp, string(nonce), encryptXmlData))
	encryptBody.TimeStamp = timeStamp

	return xml.Marshal(encryptBody)
	}

func MakeEncryptXmlData(toUserName, fromUserName, msgType, content, timeStamp string)(encryptXmlData string, err error){
	responseBody := ResponseBody{
		ToUserName: value2CDATA(toUserName),
		FromUserName: value2CDATA(fromUserName),
		MsgType: value2CDATA(msgType),
		Content: value2CDATA(content),
		CreateTime: timeStamp,
	}
	body, err := xml.Marshal(responseBody)
	if err != nil{
		glog.Error("xml marshal failed; message:", err)
		return
	}
	glog.Info("xml Marshal :", string(body))
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, int32(len(body)))
	if err != nil {
		glog.Error("binary write buf failed; ", err)
		return
	}
	bodyLength := buf.Bytes()
	randomBytes := RandByte(16)
	plainData := bytes.Join([][]byte{randomBytes, bodyLength, body, []byte(corpID)}, nil)
	cipherData, err := aesEncrypt(plainData, AesKey)
	if err != nil {
		glog.Error("AesEncrypt failed; message: ", err)
		return
	}
	return base64.StdEncoding.EncodeToString(cipherData), nil
}
