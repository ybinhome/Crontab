package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var (
		cmd *exec.Cmd
		err error
	)

	// 调用 os/exec 包中的 Command 函数创建一个命令对象
	cmd = exec.Command("/bin/bash", "-c", "echo 1;echo 2")

	// 执行创建好的命令对象，并接收命令执行的错误信息
	err = cmd.Run()
	fmt.Println(err)
}
