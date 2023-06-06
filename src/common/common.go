package common

import (
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrsquvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandByte(n int) []byte {
	b := make([]byte, n)
	for index := range b {
		b[index] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return b
}

const num = "123456789"

func RandNum(n int) []byte {
	b := make([]byte, n)
	for index := range b {
		b[index] = num[rand.Int63()%int64(len(num))]
	}
	return b
}
