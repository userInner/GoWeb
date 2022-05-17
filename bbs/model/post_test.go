package model

import (
	"fmt"
	"testing"
	"time"
)

func Test_post(t *testing.T) {
	t.Run("test InsertPost", TestInsertPost)
	t.Run("test queryTopic", TestQueryBoardNews)
}

func TestInsertPost(t *testing.T) {
	post := &Post{
		Title:     "demo",
		Content:   "demodemodemodemodedmeodemodemo",
		Last_post: time.Now().Format("2006-01-02 15:04:05"),
		User_id:   1,
		Topic_id:  1,
	}
	InsertPost(post)
}

func TestQueryBoardNews(t *testing.T) {
	boardNews, _ := QueryBoardNews(1)
	for _, v := range boardNews.Topic {
		fmt.Println("postTitle", v.Boardtopic_lastpost.Title)
		fmt.Println("作者", v.Topic_post_username)
		fmt.Println()

	}

	fmt.Println(boardNews.LinkBoard_id)

}
