package main

import (
	"gotest/GoWeb/bbs/control"
	"net/http"
)

func main() {
	//static资源，外链
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	http.Handle("/static/image/", http.StripPrefix("/static/image/", http.FileServer(http.Dir("views/static/image"))))
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./views/"))))

	//主页重定向
	http.HandleFunc("/hello", control.HandleIndex)

	//登录注册注销
	http.HandleFunc("/login", control.Login)
	http.HandleFunc("/reg", control.Reg)
	http.HandleFunc("/checkUser", control.CheckUserIsOk)
	http.HandleFunc("/logout", control.Logout)

	http.HandleFunc("/Demo", control.Demo)

	//发布帖子
	http.HandleFunc("/list", control.PublishPost)

	//主页主题展示
	http.HandleFunc("/ShowBoards", control.ShowBoards)

	//帖子展示
	http.HandleFunc("/ShowPagePost", control.ShowPagePost)

	//去post
	http.HandleFunc("/post", control.GoPost)

	//分布式
	http.HandleFunc("/ReIndexHeadLoad", control.ReIndexHeadLoad)
	http.HandleFunc("/ReIndexMainLoad", control.ReIndexMainLoad)
	http.HandleFunc("/ReIndexFootLoad", control.ReIndexFootLoad)
	http.HandleFunc("/DetailMainLoad", control.DetailMainLoad)
	http.HandleFunc("/ListMainLoad", control.ListMainLoad)
	http.HandleFunc("/PostMainLoad", control.PostMainLoad)
	//http.HandleFunc("/demo", control.IsLogin)
	http.ListenAndServe(":8080", nil)
}
