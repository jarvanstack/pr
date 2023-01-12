package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

	// 读取文件然后使用正则表达式替换
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 替换
	newBs := replace(string(bs))

	// 写入文件
	newFileName := "replace." + filename
	err = ioutil.WriteFile(newFileName, []byte(newBs), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 打印成功消息
	fmt.Printf("Success! New file: %s", newFileName)
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
