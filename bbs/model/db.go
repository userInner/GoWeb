package model

import (
	"database/sql"
	"fmt"
)

var (
	DB  *sql.DB
	err error
)

//user表
type Users struct {
	Id       int
	UserName string
	Pwd      string
	Birthday string
	Sex      string
	HeadImg  string
	IsLogin  bool
	Btn      bool //提示发帖成功
}

type Sessions struct {
	Sessions_id string
	Username    string
	User_id     int
}

type Board struct {
	Board_id        int
	Board_name      string
	Board_introduce string
}

type Topic struct {
	Topic_id            int
	Topic_name          string
	Topic_introduce     string
	Board_id            int
	Boardtopic_lastpost *Post //该主题的最后一条帖子
	Topic_sum           int   //该主题的帖子数量
	Topic_post_username string
}

type Post struct {
	Post_id   int
	Title     string
	Content   string
	Last_post string
	User_id   int
	Topic_id  int
	Author    string
}

type Reply struct {
	Reply_id      int
	Reply_user_id int
	Reply_content string
	Reply_time    string
	Reply_post_id string
}

func init() {
	DB, err = sql.Open("mysql", "root:root@(localhost:3306)/bbs")
	if err != nil {
		fmt.Println(err)
		return
	}
}
