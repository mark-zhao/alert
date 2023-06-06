package common

import (
	"fmt"
	"golang.org/x/net/context"
	"os/exec"
	"time"
)

func RunShell(cmd string) (result string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//args := strings.Split(cmd, " ")
	cmdCtx := exec.CommandContext(ctx, "bash", "-c", cmd)
	out, err := cmdCtx.Output()
	if ctx.Err() == context.DeadlineExceeded {
		result = fmt.Sprintf("exec Command timed out")
		return
	}

	return string(out), err
}
