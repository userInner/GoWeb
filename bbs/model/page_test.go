package model

import (
	"fmt"
	"testing"
)

func TestPage(t *testing.T) {
	t.Run("test get Page", TestGetPagePosts)
}

func TestGetPagePosts(t *testing.T) {
	//Post        []*Post
	//	PageNo      int //当前页
	//	PageSize    int //每页
	//	TotalPageNo int //总页
	//	TotalRecord int //总记录
	Page, _ := GetPagePosts(1, 1)
	for _, v := range Page.Post {
		fmt.Println(v.Content)
	}
	fmt.Println("当前页", Page.PageNo)
	fmt.Println("每页几个", Page.PageSize)
	fmt.Println("总页", Page.TotalPageNo)
	fmt.Println("总记录", Page.TotalRecord)
	fmt.Println("总记录", Page.Post)
	fmt.Println("Post Content", Page.Post[0].Content)
}
