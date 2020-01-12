package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var (
		cmd    *exec.Cmd
		output []byte
		err    error
	)

	// 生成 Cmd 对象
	cmd = exec.Command("/bin/bash", "-c", "sleep 5; ls -l")

	// 调用 os/exec 包中的 CombinedOutput 方法，启动子进程执行命令，并通过 pipe 捕获输出
	if output, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(output))
}
