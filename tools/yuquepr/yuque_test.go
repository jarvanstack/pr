package main

import (
	"fmt"
	"testing"

	"github.com/dengjiawen8955/pr/util"
	"github.com/wujiyu115/yuqueg"
)

var yu *yuqueg.Service = yuqueg.NewService("")
var slug string = "fr9zuigpvhe72hlu"

func Test_yuque_doc_list(t *testing.T) {
	bd, err := yu.Doc.List("dengjiawen8955/dsne7d")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(bd)
}
func Test_yuque_doc_get(t *testing.T) {
	bd, err := yu.Doc.Get("dengjiawen8955/dsne7d", slug, &yuqueg.DocGet{Raw: 1})
	if err != nil {
		t.Fatal(err)
	}
	// 转化为 json
	fmt.Printf("bd.Data.Body: %v\n", bd.Data.Body)
	// 调用 replace 方法
}

func Test_yuque_doc_update(t *testing.T) {
	bd, err := yu.Doc.Get("dengjiawen8955/dsne7d", slug, &yuqueg.DocGet{Raw: 1})
	if err != nil {
		t.Fatal(err)
	}
	// 转化为 json
	fmt.Printf("bd.Data.Body: %v\n", bd.Data.Body)
	// 调用 replace 方法
	newBody := util.Replace(bd.Data.Body)
	dd, err := yu.Doc.Update("dengjiawen8955/dsne7d", fmt.Sprintf("%d", bd.Data.ID), &yuqueg.DocCreate{
		Title:  bd.Data.Title,
		Slug:   bd.Data.Slug,
		Public: bd.Data.Public,
		Format: bd.Data.Format,
		Body:   newBody,
	})
	if err != nil {
		t.Fatal(err)
		// {"status":400,"message":"抱歉，语雀不允许通过 API 修改富文本格式文档，请到语雀进行操作。"}
	}
	println("=====================================")
	t.Log(dd)
}
func Test_yuque_doc_create(t *testing.T) {
	bd, err := yu.Doc.Get("dengjiawen8955/dsne7d", slug, &yuqueg.DocGet{Raw: 1})
	if err != nil {
		t.Fatal(err)
	}
	// 转化为 json
	fmt.Printf("bd.Data.Body: %v\n", bd.Data.Body)
	// 调用 replace 方法
	newBody := util.Replace(bd.Data.Body)
	dd, err := yu.Doc.Create("dengjiawen8955/dsne7d", &yuqueg.DocCreate{
		Title:  bd.Data.Title + "-replace",
		Slug:   bd.Data.Slug + "-replace",
		Format: "markdown",
		Public: bd.Data.Public,
		Body:   newBody,
	})
	if err != nil {
		t.Fatal(err)
	}
	println("=====================================")
	t.Log(dd)
	fmt.Printf("dd.Data.Title: %v\n", dd.Data.Title)
	fmt.Printf("dd.Data.Slug: %v\n", dd.Data.Slug)
}
