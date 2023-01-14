package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"

	"github.com/dengjiawen8955/pr/util"
	"github.com/zeromicro/go-zero/tools/goctl/rpc/execx"
)

var (
	Version = "v1.1"
)

func printUsage() {
	fmt.Println("pr version: " + Version)
	fmt.Println("Usage: pr [filename]")
	fmt.Println("Punctuation replace:    pr test.txt")
	fmt.Println("Update tool:            pr update")
}

func main() {
	// 参数检查
	if os.Args == nil || len(os.Args) < 2 {
		printUsage()
		return
	}

	// 获取输入的第一个参数作为文件名
	filename := os.Args[1]

	// 更新
	if filename == "update" {
		Update()
		return
	}

	// 读取文件然后使用正则表达式替换
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 替换
	newBs := util.Replace(string(bs))

	// 写入文件
	newFileName := filename + ".replace" + path.Ext(filename)
	err = ioutil.WriteFile(newFileName, []byte(newBs), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 打印成功消息
	fmt.Printf("Success! New file: %s\n", newFileName)
}

// 自动更新软件
// 调用命令行 go install github.com/dengjiawen8955/pr@latest
func Update() {
	err := update()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

}
func update() error {
	cmd := `GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go install github.com/dengjiawen8955/pr@latest`
	if runtime.GOOS == "windows" {
		cmd = `set GOPROXY=https://goproxy.cn/,direct && go install github.com/dengjiawen8955/pr@latest`
	}
	info, err := execx.Run(cmd, "")
	if err != nil {
		return err
	}

	fmt.Print(info)
	return nil
}
