package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/dengjiawen8955/pr/util"
	"github.com/wujiyu115/yuqueg"
)

var (
	ErrParseUrl = errors.New("ErrParseUrl")
)

var (
	postfix = "_replace"
)

var usage string = "Usage yuquepr TOKEN FROMURL"

func main() {
	// 读取参数, 第一个是 token 第二个是 fromUrl
	params := os.Args
	if len(params) < 3 {
		fmt.Println(usage)
		return
	}
	token := params[1]
	fromUrl := params[2]

	fmt.Printf("token: %v\n", token)
	fmt.Printf("fromUrl: %v\n", fromUrl)

	// 解析 URL 获得 namespace 和 slug
	namespace, slug, err := ParseUrl(fromUrl)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 初始化服务
	yu := yuqueg.NewService(token)

	bd, err := yu.Doc.Get(namespace, slug, &yuqueg.DocGet{Raw: 1})
	if err != nil {
		fmt.Println(err)
		return
	}

	newBody := util.Replace(bd.Data.Body)

	dd, err := yu.Doc.Create(namespace, &yuqueg.DocCreate{
		Title:  bd.Data.Title + postfix,
		Slug:   bd.Data.Slug + postfix,
		Public: bd.Data.Public,
		Body:   newBody,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// 打印拼接好的 url
	fmt.Println()
	fmt.Printf("Success! New Doc URL: https://www.yuque.com/%s/%s \n\n", namespace, dd.Data.Slug)
}

// ParseUrl 解析 URL 获得 namespace 和 slug
// https://www.yuque.com/dengjiawen8955/dsne7d/fr9zuigpvhe72hlu-replace
// https://www.yuque.com/{namespace}/{slug}
func ParseUrl(urlStr string) (namespace string, slug string, err error) {
	ss1 := strings.Split(urlStr, "yuque.com/")
	if len(ss1) != 2 {
		return "", "", ErrParseUrl
	}
	ss2 := strings.Split(ss1[1], "/")
	if len(ss2) != 3 {
		return "", "", ErrParseUrl
	}
	return ss2[0] + "/" + ss2[1], ss2[2], nil
}
