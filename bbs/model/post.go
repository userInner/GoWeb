package model

import (
	"fmt"
	"strconv"
)

type BoardNews struct {
	Topic        []*Topic //主题
	LinkBoard_id string   //链接地址
}

func InsertPost(post *Post) error {
	sqlStr := "insert into post(title,content,user_id,topic_id,last_post) values(?,?,?,?,?)"

	_, err := DB.Exec(sqlStr, post.Title, post.Content, post.User_id, post.Topic_id, post.Last_post)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func QueryBoardNews(board_id int) (*BoardNews, error) {
	sqlStr := "select topic_id,topic_name,topic_introduce,topic_id from topic where board_id = ?"

	rows, err := DB.Query(sqlStr, board_id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//最后发布的帖子 主题帖子数量 作者
	sqlStrLastPost := "select post_id,title,content,last_post,user_id,topic_id from post where topic_id = ? order by post_id desc limit 0,1"
	sqlStrCountPost := "select count(post_id) from post where topic_id = ?"
	sqlStrLastPostUserName := "select userName from user where id = ?"
	//topic仓库
	var storageTopic []*Topic
	for rows.Next() {
		//主题
		topic := &Topic{}
		rows.Scan(&topic.Topic_id, &topic.Topic_name, &topic.Topic_introduce, &topic.Board_id)

		//该主题最后发表的帖子
		row := DB.QueryRow(sqlStrLastPost, topic.Topic_id)
		lastPost := &Post{}
		row.Scan(&lastPost.Post_id, &lastPost.Title, &lastPost.Content, &lastPost.Last_post, &lastPost.User_id, &lastPost.Topic_id)
		topic.Boardtopic_lastpost = lastPost

		//最后发布的作者
		var lastPostUser string
		lastPostUserNameRow := DB.QueryRow(sqlStrLastPostUserName, lastPost.User_id)
		lastPostUserNameRow.Scan(&lastPostUser)
		topic.Topic_post_username = lastPostUser
		//主题内帖子的数量
		var count int
		countRow := DB.QueryRow(sqlStrCountPost, topic.Topic_id)
		countRow.Scan(&count)
		topic.Topic_sum = count
		storageTopic = append(storageTopic, topic)
	}

	stringBoard_id := strconv.Itoa(board_id)
	boardNews := &BoardNews{
		Topic:        storageTopic,
		LinkBoard_id: stringBoard_id,
	}
	return boardNews, nil
}
