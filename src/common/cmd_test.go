package common

import (
	"testing"
)

func TestRunShell(t *testing.T){
	var cmd = "ls /root -l"
	output, err := RunShell(cmd)
	if err != nil{
		t.Fatal("run shell failed; cmd: ", cmd, "; msg:", err)
	}
	t.Log(output)
}
