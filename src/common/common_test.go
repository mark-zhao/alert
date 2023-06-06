package common

import "testing"

func TestRandByte(t *testing.T){
	rands := RandByte(8)
	t.Log("RandByte pass; 8 RandByte: ", string(rands))
}

func TestRandNum(t *testing.T){
	numRands := RandNum(8)
	t.Log("RandNum pass; 8 rand num: ", string(numRands))
}