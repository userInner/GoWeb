package control

import (
	"fmt"
	"gotest/GoWeb/bbs/model"
	"html/template"
	"net/http"
	"strconv"
)

//主页
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	//重定向
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, "")
}

//logo，登录
func ReIndexHeadLoad(w http.ResponseWriter, r *http.Request) {
	flag, _ := IsLogin(r)
	cookie, _ := r.Cookie("user")
	if flag { //已登录
		//fmt.Println(1)
		//找用户
		sess, _ := model.SessionIdGet(cookie.Value)
		user, _ := model.SelectNews(sess.User_id)
		//传head
		t := template.Must(template.ParseFiles("views/index/head.html"))
		t.Execute(w, user.UserName)
	} else {
		t := template.Must(template.ParseFiles("views/index/head.html"))
		t.Execute(w, "")
	}

}

//主页内容
func ReIndexMainLoad(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/index/main.html"))
	t.Execute(w, "")
}

//页脚
func ReIndexFootLoad(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/index/foot.html"))
	t.Execute(w, "")
}

func DetailMainLoad(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/datail/main.html"))
	t.Execute(w, "")
}

func ListMainLoad(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/list/main.html"))
	//page, _ := model.GetPagePosts(1)
	t.Execute(w, "")
}

func PostMainLoad(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/post/main.html"))
	t.Execute(w, "")
}

//展示board
func ShowBoards(w http.ResponseWriter, r *http.Request) {
	boardId := r.FormValue("board_id")
	intBoardId, err := strconv.Atoi(boardId)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	var boardSlice = []string{"人事管理", "机器管理", "销售管理"}
	//主题
	boardNews, _ := model.QueryBoardNews(intBoardId)
	var tempBoardNewsStr string
	if boardNews != nil {
		tempBoardNewsStr += "<tr class='tr3'>" + "<td colspan='4'>" + boardSlice[intBoardId-1] + "</td>" + "</tr>"
		for _, v := range boardNews.Topic {
			topicLink := "/ShowPagePost?topic_id=" + strconv.Itoa(v.Topic_id) + "&pageNo=" + "1"
			tempBoardNewsStr += "<tr class='tr3'>" + "<td width='5%'>" + "&nbsp;" + "</td>"
			tempBoardNewsStr += "<th align='left'>" + "<img src='static/image/board.gif'/>" + "<a href=" + topicLink + ">" + v.Topic_name + "</a>" + "</th>"
			tempBoardNewsStr += "<td align='center'>" + strconv.Itoa(v.Topic_sum) + "</td>"
			tempBoardNewsStr += "<th>" + "<span>" + "<a href='detail.html'>" + v.Boardtopic_lastpost.Title + "</a>" + "</span>" + "<br/>" + "<span>" + v.Topic_post_username + "</span>" + "<span class='gray'>" + v.Boardtopic_lastpost.Last_post + "</span>" + "</th>" + "</tr>"
		}
	}
	intBoardId++
	boardNews, _ = model.QueryBoardNews(intBoardId)
	var tempBoardNewsStr2 string
	if boardNews != nil {
		tempBoardNewsStr2 += "<tr class='tr3'>" + "<td colspan='4'>" + boardSlice[intBoardId-1] + "</td>" + "</tr>"
		for _, v := range boardNews.Topic {
			topicLink := "/ShowPagePost?topic_id=" + strconv.Itoa(v.Topic_id) + "&pageNo=" + "1"
			tempBoardNewsStr2 += "<tr class='tr3'>" + "<td width='5%'>" + "&nbsp;" + "</td>"
			tempBoardNewsStr2 += "<th align='left'>" + "<img src='static/image/board.gif'/>" + "<a href=" + topicLink + ">" + v.Topic_name + "</a>" + "</th>"
			tempBoardNewsStr2 += "<td align='center'>" + strconv.Itoa(v.Topic_sum) + "</td>"
			tempBoardNewsStr2 += "<th>" + "<span>" + "<a href='detail.html'>" + v.Boardtopic_lastpost.Title + "</a>" + "</span>" + "<br/>" + "<span>" + v.Topic_post_username + "</span>" + "<span class='gray'>" + v.Boardtopic_lastpost.Last_post + "</span>" + "</th>" + "</tr>"
		}
	}
	intBoardId++
	boardNews, _ = model.QueryBoardNews(intBoardId)
	var tempBoardNewsStr3 string
	if boardNews != nil {
		tempBoardNewsStr3 += "<tr class='tr3'>" + "<td colspan='4'>" + boardSlice[intBoardId-1] + "</td>" + "</tr>"
		for _, v := range boardNews.Topic {
			topicLink := "/ShowPagePost?topic_id=" + strconv.Itoa(v.Topic_id) + "&pageNo=" + "1"
			tempBoardNewsStr3 += "<tr class='tr3'>" + "<td width='5%'>" + "&nbsp;" + "</td>"
			tempBoardNewsStr3 += "<th align='left'>" + "<img src='static/image/board.gif'/>" + "<a href=" + topicLink + ">" + v.Topic_name + "</a>" + "</th>"
			tempBoardNewsStr3 += "<td align='center'>" + strconv.Itoa(v.Topic_sum) + "</td>"
			tempBoardNewsStr3 += "<th>" + "<span>" + "<a href='detail.html'>" + v.Boardtopic_lastpost.Title + "</a>" + "</span>" + "<br/>" + "<span>" + v.Topic_post_username + "</span>" + "<span class='gray'>" + v.Boardtopic_lastpost.Last_post + "</span>" + "</th>" + "</tr>"
		}
	}
	w.Write([]byte(tempBoardNewsStr + tempBoardNewsStr2 + tempBoardNewsStr3))
}

//展示该主题一页的信息
func ShowPagePost(w http.ResponseWriter, r *http.Request) {
	topic_id := r.FormValue("topic_id") //链接传过来的主题号
	pageNo := r.FormValue("pageNo")     //页码

	intTopicId, _ := strconv.Atoi(topic_id)
	intPageNo, _ := strconv.Atoi(pageNo)
	pageNews, _ := model.GetPagePosts(intPageNo, intTopicId)
	if intTopicId > pageNews.TotalPageNo {
		t := template.Must(template.ParseFiles("views/list/main.html"))
		t.Execute(w, "")
		return
	}
	var tempPagePostNews string

	for _, v := range pageNews.Post {
		userNews, _ := model.SelectNews(v.User_id)
		replys, _ := model.GetReplyMsg(v.Post_id)
		//计算回复数量
		replyCount := 0
		if replys != nil {
			for _, k := range replys {
				if k != nil {
					replyCount++
				}
			}
		}
		tempPagePostNews += "<tr class='tr3'>"
		tempPagePostNews += "<td>" + "<img src='static/image/topic.gif' border=0/>" + "</td>"
		tempPagePostNews += "<TD style='FONT-SIZE: 15px'>" + "<a href=" + strconv.Itoa(v.Post_id) + ">" + v.Title + "</a>" + "</td>"
		tempPagePostNews += "<td align='center'>" + userNews.UserName + "</td>"
		tempPagePostNews += "<TD align='center'>" + strconv.Itoa(replyCount) + "</TD>"
		tempPagePostNews += "</tr>"
	}
	var tempUpNo int
	var tempNextNo int
	if intPageNo-1 < 1 {
		tempUpNo = 1
	} else {
		tempUpNo = intPageNo - 1
	}
	if intPageNo+1 > pageNews.TotalPageNo {
		tempNextNo = pageNews.TotalPageNo
	} else {
		tempNextNo = intPageNo + 1
	}
	tempUpLink := "/ShowPagePost?topic_id=" + strconv.Itoa(intTopicId) + "&pageNo=" + strconv.Itoa(tempUpNo)
	tempNextLink := "/ShowPagePost?topic_id=" + strconv.Itoa(intTopicId) + "&pageNo=" + strconv.Itoa(tempNextNo)
	page := &model.Page{
		Post:          pageNews.Post,
		PageNo:        pageNews.PageNo,
		PageSize:      pageNews.PageSize,
		TotalPageNo:   pageNews.TotalPageNo,
		TotalRecord:   pageNews.TotalPageNo,
		Totalstr:      tempPagePostNews,
		TotalUpHref:   tempUpLink,
		TotalNextHref: tempNextLink,
	}
	t := template.Must(template.ParseFiles("views/list/main.html"))
	t.Execute(w, page)
}
