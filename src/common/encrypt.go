package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"glog"
	"io"
	"sort"
	"strings"
)

type Encrypt interface {

}

var AesKey []byte
func init(){
	AesKey, _ = base64.StdEncoding.DecodeString(key + "=")
}

func generateSignature (timestamp, nonce, msgEncrypt string ) string  {
	var base = []string{token, timestamp, nonce, msgEncrypt}
	sort.Strings(base)
	sha1Hash := sha1.New()
	io.WriteString(sha1Hash, strings.Join(base, ""))
	return fmt.Sprintf("%x", sha1Hash.Sum(nil))
}

func Decrypt(msgEncrypt string) (msg []byte, err error){
	aesMsg, err := base64.StdEncoding.DecodeString(msgEncrypt)
	if err != nil{
		glog.Error("base64 decode failed; message: ", err)
		return
	}
	msg, err = aesDecrypt(aesMsg, AesKey)
	if err != nil{
		glog.Error("AES decrypt failed; message: ", err)
		return
	}

	len := binary.BigEndian.Uint32(msg[16:20])

	return msg[20: 20 + len], nil
	}

func aesDecrypt(cipherData []byte, aesKey []byte) ([]byte, error) {
	k := len(aesKey) //PKCS#7
	glog.Info("len(aesKey): ", k, "; len(cipherData):", len(cipherData))
	if len(cipherData)%k != 0 {
		return nil, errors.New("crypto/cipher: ciphertext size is not multiple of aes key length")
	}

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	plainData := make([]byte, len(cipherData))
	blockMode.CryptBlocks(plainData, cipherData)
	return plainData, nil
}

func aesEncrypt(plainData []byte, aesKey []byte) ([]byte, error) {
	k := len(aesKey)

	if len(plainData)%k != 0 {
		plainData, _ = PKCS7Pad(plainData, k)
	}

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	cipherData := make([]byte, len(plainData))
	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(cipherData, plainData)

	return cipherData, nil
}

// PadLength calculates padding length, from github.com/vgorin/cryptogo
func PadLength(slice_length, blocksize int) (padlen int) {
	padlen = blocksize - slice_length%blocksize
	if padlen == 0 {
		padlen = blocksize
	}
	return padlen
}
//from github.com/vgorin/cryptogo
func PKCS7Pad(message []byte, blocksize int) (padded []byte, err error) {
	// block size must be bigger or equal 2
	if blocksize < 1<<1 {
		err = errors.New("block size is too small (minimum is 2 bytes)")
		return
	}
	// block size up to 255 requires 1 byte padding
	if blocksize < 1<<8 {
		// calculate padding length
		padlen := PadLength(len(message), blocksize)

		// define PKCS7 padding block
		padding := bytes.Repeat([]byte{byte(padlen)}, padlen)

		// apply padding
		padded = append(message, padding...)
	}else {
		// block size bigger or equal 256 is not currently supported
		err = errors.New("unsupported block size")
	}
	return
}
