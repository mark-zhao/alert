package common

import (
	"crypto/sha1"
	"fmt"
	"io"
	"sort"
	"strings"
)

func ValidateUrl(timestamp, nonce, signatureIn string) bool {
	signatureGen := MakeSignature(timestamp, nonce)
	if signatureGen != signatureIn {
		return false
	}
	return true
}

func MakeSignature(timestamp, nonce string) string {
	sl := []string{token, timestamp, nonce}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}

func ValidateMsg(timestamp, nonce, msgEncrypt, msgSignatureIn string) bool {
	msgSignatureGen := MakeMsgSignature(timestamp, nonce, msgEncrypt)
	if msgSignatureGen != msgSignatureIn {
		return false
	}
	return true
}

func MakeMsgSignature(timestamp, nonce, msg_encrypt string) string {
	sl := []string{token, timestamp, nonce, msg_encrypt}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}