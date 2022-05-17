package control

import (
	"github.com/google/uuid"
	"gotest/GoWeb/bbs/model"
	"html/template"
	"net/http"
	"time"
)

//处理用户登录
func Login(w http.ResponseWriter, r *http.Request) {
	uName := r.PostFormValue("uName")
	uPass := r.PostFormValue("uPass")

	user, _ := model.CheckUserAndPwd(uName, uPass)

	//flag, CookieUserName := IsLogin(r)

	if user != nil {
		flag, _ := IsLogin(r) //查看是否有过登录

		if flag == false { //无登录
			uuid := uuid.New().String()
			session := &model.Sessions{
				uuid,
				user.UserName,
				user.Id,
			}
			model.AddSession(session)
			//建立
			cookie := http.Cookie{
				Name:     "user",
				Value:    uuid,
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)
		}
		t := template.Must(template.ParseFiles("views/index.html"))
		t.Execute(w, user)
	} else {
		t := template.Must(template.ParseFiles("views/login.html"))
		t.Execute(w, nil)
	}

}

//处理注册
func Reg(w http.ResponseWriter, r *http.Request) {
	uName := r.PostFormValue("uName")
	uPass := r.PostFormValue("uPass")
	//uPass1 := r.PostFormValue("uPass1")
	uBirthday := time.Now().Format("2006-01-02 15:04:05")
	gender := r.PostFormValue("gender")
	head := r.PostFormValue("head")

	err := model.InsertUser(uName, uPass, uBirthday, gender, head)
	if err != nil {
		t := template.Must(template.ParseFiles("views/reg.html"))
		t.Execute(w, false)
	} else {
		t := template.Must(template.ParseFiles("views/login.html"))
		t.Execute(w, true)
	}
}

//处理登录用户是否存在
func CheckUserIsOk(w http.ResponseWriter, r *http.Request) {
	uName := r.PostFormValue("userName")
	flag, _ := model.CheckUser(uName)
	if flag { //存在
		w.Write([]byte("<font style='color:red'>用户已存在</font>"))
	} else {
		w.Write([]byte("<font style='color:green'>用户可以注册</font>"))
	}
}

//查看是否登录
func IsLogin(r *http.Request) (bool, string) {
	cookie, _ := r.Cookie("user")
	if cookie != nil {

		cookieValue := cookie.Value
		//查看数据库是否有对应的session
		sess, _ := model.QuerySession(cookieValue)
		if sess.User_id > 0 {
			return true, sess.Username
		}
	}
	return false, ""
}

//注销
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		//获取user值
		cookieValue := cookie.Value
		//删除对应的session
		model.DelSession(cookieValue)
		cookie.MaxAge = -1
	}

	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, "")
}
