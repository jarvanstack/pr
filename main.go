package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"regexp"
)

const (
	usage = `Usage: pr [filename]`
)

func main() {
	// 参数检查
	if os.Args == nil || len(os.Args) < 2 {
		fmt.Println(usage)
		return
	}

	// 获取输入的第一个参数作为文件名
	filename := os.Args[1]

	// 更新
	if filename == "update" {
		update()
		return
	}

	// 读取文件然后使用正则表达式替换
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 替换
	newBs := replace(string(bs))

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
func update() {
	// 执行命令
	cmd := exec.Command("go", "install", "github.com/dengjiawen8955/pr@latest")
	cmd.Stdout = os.Stdout
	fmt.Printf("cmd: %v\n", cmd.String())
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Run failed: %v\n", err)
		return
	}

	// 打印成功消息
	fmt.Println("Update success!")
}

func replace(data string) string {
	// 句号
	re := regexp.MustCompile(`([\p{Han}])\.[ ]*`)
	output := re.ReplaceAllString(data, `$1。`)

	// 逗号
	re = regexp.MustCompile(`([\p{Han}]),[ ]*`)
	output = re.ReplaceAllString(output, `$1，`)

	// 冒号
	re = regexp.MustCompile(`([\p{Han}]):[ ]*`)
	output = re.ReplaceAllString(output, `$1：`)

	// 分号
	re = regexp.MustCompile(`([\p{Han}]);[ ]*`)
	output = re.ReplaceAllString(output, `$1；`)

	// 问号
	re = regexp.MustCompile(`([\p{Han}])\?[ ]*`)
	output = re.ReplaceAllString(output, `$1？`)

	// 感叹号
	re = regexp.MustCompile(`([\p{Han}])![ ]*`)
	output = re.ReplaceAllString(output, `$1！`)

	// 左括号
	re = regexp.MustCompile(`[ ]*\(([\p{Han}])`)
	output = re.ReplaceAllString(output, `（$1`)

	// 右括号
	re = regexp.MustCompile(`([\p{Han}])\)[ ]*`)
	output = re.ReplaceAllString(output, `$1）`)

	// 左引号
	re = regexp.MustCompile(`[ ]*"([\p{Han}])`)
	output = re.ReplaceAllString(output, `“$1`)

	// 右引号
	re = regexp.MustCompile(`([\p{Han}])"[ ]*`)
	output = re.ReplaceAllString(output, `$1”`)

	// 左小引号
	re = regexp.MustCompile(`[ ]*'([\p{Han}])`)
	output = re.ReplaceAllString(output, `‘$1`)

	// 右小引号
	re = regexp.MustCompile(`([\p{Han}])'[ ]*`)
	output = re.ReplaceAllString(output, `$1’`)

	return output
}
