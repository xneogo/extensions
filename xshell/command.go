/*
 *  ┏┓      ┏┓
 *┏━┛┻━━━━━━┛┻┓
 *┃　　　━　　  ┃
 *┃   ┳┛ ┗┳   ┃
 *┃           ┃
 *┃     ┻     ┃
 *┗━━━┓     ┏━┛
 *　　 ┃　　　┃神兽保佑
 *　　 ┃　　　┃代码无BUG！
 *　　 ┃　　　┗━━━┓
 *　　 ┃         ┣┓
 *　　 ┃         ┏┛
 *　　 ┗━┓┓┏━━┳┓┏┛
 *　　   ┃┫┫  ┃┫┫
 *      ┗┻┛　 ┗┻┛
 @Time    : 2025/7/3 -- 11:23
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2025 亓官竹
 @Description: xother xother/command.go
*/

package xshell

import (
	"context"
	"fmt"
	oexec "os/exec"
	"strings"

	"github.com/spf13/cast"
	"github.com/xneogo/extensions/colorlog"
)

func command(ctx context.Context, str string) ([]string, error) {
	cmd := oexec.Command("sh", "-c", str)
	output, err := cmd.Output()
	if err != nil {
		colorlog.Errorf(ctx, "cmd: %s, output: %s", str, output)
		return nil, err
	} else {
		colorlog.Debugf(ctx, "cmd: %s, output: %s", str, output)
	}
	ss := strings.Split(string(output), "\n")
	res := make([]string, 0, len(ss))
	for _, s := range ss {
		res = append(res, strings.TrimRight(s, "\n"))
	}
	return res, nil
}

func Exec(ctx context.Context, host string, str string) ([]string, error) {
	if host != "" {
		str = fmt.Sprintf("ssh root@%s %s", host, str)
	}
	return command(ctx, str)
}

func LS(ctx context.Context, host string, path string) ([][]string, error) {
	path = strings.TrimRight(path, "/")
	str := fmt.Sprintf("ls %s", path)
	if host != "" {
		str = fmt.Sprintf("ssh root@%s %s", host, str)
	}
	ss, err := command(ctx, str)
	if err != nil {
		return nil, err
	}
	res := make([][]string, 0, len(ss))
	for _, s := range ss {
		if s == "" {
			continue
		}
		if s[0] == '/' {
			res = append(res, []string{s, s[strings.LastIndex(s, "/")+1:]})
		} else {
			res = append(res, []string{path + "/" + s, s})
		}
	}
	return res, nil
}

func RM(ctx context.Context, host string, path string) error {
	if path == "/" {
		return nil
	}
	str := fmt.Sprintf("rm -rf %s", path)
	if host != "" {
		str = fmt.Sprintf("ssh root@%s %s", host, str)
	}
	_, err := command(ctx, str)
	if err != nil {
		return err
	}
	return nil
}

func MK(ctx context.Context, host string, path string) error {
	str := fmt.Sprintf("mkdir -p %s", path)
	if host != "" {
		str = fmt.Sprintf("ssh root@%s %s", host, str)
	}
	_, err := command(ctx, str)
	if err != nil {
		return err
	}
	return nil
}

func MV(ctx context.Context, host string, oldPath, newPath string) error {
	str := fmt.Sprintf("mv -f %s %s", oldPath, newPath)
	if host != "" {
		str = fmt.Sprintf("ssh root@%s %s", host, str)
	}
	_, err := command(ctx, str)
	if err != nil {
		return err
	}
	return nil
}

func CP(ctx context.Context, host string, oldPath, newPath string) error {
	str := fmt.Sprintf("cp -rf %s %s", oldPath, newPath)
	if host != "" {
		str = fmt.Sprintf("ssh root@%s %s", host, str)
	}
	_, err := command(ctx, str)
	if err != nil {
		return err
	}
	return nil
}

func Lines(ctx context.Context, host string, path string) int64 {
	str := fmt.Sprintf("wc -l %s | awk '{print $1}'", path)
	if host != "" {
		str = fmt.Sprintf("ssh root@%s %s", host, str)
	}
	ss, err := command(ctx, str)
	if err != nil || len(ss) == 0 {
		return 0
	}
	return cast.ToInt64(ss[0])
}

func PIDs(ctx context.Context, host string, keyWord string) []int64 {
	str := fmt.Sprintf("ps -ef | grep %s | grep -v grep | awk '{print $2}'", keyWord)
	if host != "" {
		str = fmt.Sprintf("ssh root@%s %s", host, str)
	}
	ss, err := command(ctx, str)
	if err != nil || len(ss) == 0 {
		return nil
	}
	res := make([]int64, 0, len(ss))
	for _, s := range ss {
		if s == "" {
			continue
		}
		res = append(res, cast.ToInt64(s))
	}
	return res
}

func Kill(ctx context.Context, host string, keyWord string) error {
	str := fmt.Sprintf("ps -ef | grep %s | grep -v grep | awk '{print $2}' | xargs kill -15", keyWord)
	if host != "" {
		str = fmt.Sprintf("ssh root@%s %s", host, str)
	}
	_, err := command(ctx, str)
	if err != nil {
		return err
	}
	return nil
}

func ScpFrom(ctx context.Context, host string, remotePath string, localPath string) error {
	str := fmt.Sprintf("scp -r root@%s:%s %s", host, remotePath, localPath)
	_, err := command(ctx, str)
	if err != nil {
		return err
	}
	return nil
}

func ScpTo(ctx context.Context, localPath string, host string, remotePath string) error {
	str := fmt.Sprintf("scp -r %s root@%s:%s", localPath, host, remotePath)
	_, err := command(ctx, str)
	if err != nil {
		return err
	}
	return nil
}
