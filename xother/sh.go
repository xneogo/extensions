/*
 @Time    : 2024/9/6 -- 10:49
 @Description: sh.go
*/

package xother

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func RunShWithScanner(_ string, name string, args ...string) error {
	// 创建一个 buffer 用于存储脚本的输入
	// input := []byte("some input data\n")

	// 创建命令
	cmd := exec.Command(name, args...)

	// 可选：设置脚本需要的环境变量
	// cmd.Env = append(os.Environ(), "ENV_VAR=value")

	// 获取命令的输入流
	// stdin, err := cmd.StdinPipe()
	// if err != nil {
	// 	fmt.Printf("Error creating stdin pipe: %s\n", err)
	// 	return nil
	// }

	// 获取命令的输出流
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error creating stdout pipe: %s\n", err)
		return err
	}

	// 获取命令的错误输出流
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Printf("Error creating stderr pipe: %s\n", err)
		return err
	}

	// 启动命令
	if err := cmd.Start(); err != nil {
		fmt.Printf("Error starting command: %s\n", err)
		return err
	}

	// 向命令的 stdin 写入数据
	// go func() {
	// 	_, err := stdin.Write(input)
	// 	if err != nil {
	// 		fmt.Printf("Error writing to stdin: %s\n", err)
	// 		return
	// 	}
	// 	_ = stdin.Close() // 关闭写入器，表示完成写入
	// }()

	// 读取命令的 stdout
	go func() {
		_, err := io.Copy(os.Stdout, stdout)
		if err != nil {
			fmt.Printf("Error reading stdout: %s\n", err)
			return
		}
	}()

	// 读取命令的 stderr
	go func() {
		_, err := io.Copy(os.Stderr, stderr)
		if err != nil {
			fmt.Printf("Error reading stderr: %s\n", err)
			return
		}
	}()

	// 等待命令执行完成
	if err := cmd.Wait(); err != nil {
		fmt.Printf("Command finished with error: %s\n", err)
		return err
	}
	return nil
}
