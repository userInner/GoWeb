package model

import "fmt"

type Page struct {
	Post          []*Post
	PageNo        int    //当前页
	PageSize      int    //每页
	TotalPageNo   int    //总页
	TotalRecord   int    //总记录
	Totalstr      string //html字符串
	TotalUpHref   string //当前上一页链接地址
	TotalNextHref string //当前上一页链接地址
}

//get翻页Post信息
func GetPagePosts(pageNo int, topic_id int) (*Page, error) {
	sqlStr := "select count(*) from post where topic_id = ?"
	var totalRecord int
	row := DB.QueryRow(sqlStr, topic_id)
	row.Scan(&totalRecord)

	//每页显示十条
	var pageSize int = 10
	var totalPageNo int //总页数
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}

	//获取当前页
	sqlStr2 := "select post_id,title,content,last_post,user_id,topic_id from post where topic_id = ? limit ?,?"
	rows, err := DB.Query(sqlStr2, topic_id, (pageNo-1)*pageSize, pageSize)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var posts []*Post
	for rows.Next() {
		post := &Post{}
		rows.Scan(&post.Post_id, &post.Title, &post.Content, &post.Last_post, &post.User_id, &post.Topic_id)
		posts = append(posts, post)
	}
	Page := &Page{
		Post:        posts,
		PageNo:      pageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	return Page, nil
}
