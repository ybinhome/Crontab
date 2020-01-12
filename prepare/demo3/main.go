package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type result struct {
	err    error
	output []byte
}

func main() {
	var (
		ctx        context.Context
		cancelFunc context.CancelFunc
		cmd        *exec.Cmd
		resultChan chan result
		res        result
	)

	ctx, cancelFunc = context.WithCancel(context.TODO())

	resultChan = make(chan result, 1000)

	go func() {
		var (
			err    error
			output []byte
		)

		cmd = exec.CommandContext(ctx, "/bin/bash", "-c", "sleep 2; echo hello")

		output, err = cmd.CombinedOutput()

		resultChan <- result{
			err:    err,
			output: output,
		}
	}()

	time.Sleep(1 * time.Second)

	cancelFunc()

	res = <-resultChan
	fmt.Println(res.err, string(res.output))
}
