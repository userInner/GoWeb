package control

import (
	"fmt"
	"gotest/GoWeb/bbs/model"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

//去发送post页面
func GoPost(w http.ResponseWriter, r *http.Request) {
	cookieVal, _ := r.Cookie("user")
	sess, _ := model.SessionIdGet(cookieVal.Value)
	fmt.Println(sess.User_id)
	userNews, _ := model.SelectNews(sess.User_id)

	if userNews != nil {
		t := template.Must(template.ParseFiles("./views/post.html"))
		t.Execute(w, userNews)
	}
}

//发表帖子
func PublishPost(w http.ResponseWriter, r *http.Request) {
	titleVal := r.PostFormValue("title")
	contentVal := r.PostFormValue("content")
	userIdVal := r.PostFormValue("userId")
	newUserIdVal, _ := strconv.Atoi(userIdVal)
	last_postTime := time.Now().Format("2006-01-02 15:04:05")
	topicIdVal := r.PostFormValue("topicId")
	newtopicIdVal, _ := strconv.Atoi(topicIdVal)
	//user_id := r.PostFormValue("topicId")
	post := &model.Post{
		Title:     titleVal,
		Content:   contentVal,
		Last_post: last_postTime,
		User_id:   newUserIdVal,
		Topic_id:  newtopicIdVal,
	}
	err := model.InsertPost(post)

	//flag,userName := IsLogin(r)
	//重定向
	if err != nil {
		t := template.Must(template.ParseFiles("./views/post.html"))
		t.Execute(w, false) //创建失败
	} else {
		t := template.Must(template.ParseFiles("./views/list.html"))

		t.Execute(w, "")
	}
	//model.InsertPost()
}
