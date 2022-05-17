package model

import (
	"fmt"
	"testing"
	"time"
)

func TestUser(t *testing.T) {
	//user, _ := CheckUserAndPwd(123456, "123456")
	//fmt.Println("用户信息：", user)
	//fmt.Println("测试。。。。")
	t.Run("测试登录", testLogin)
	//t.Run("测试插入", testInsertUser)
	t.Run("测试数据库", testDB)
	t.Run("测试查询用户信息", testSelectNews)
}
func testDB(t *testing.T) {
	row := DB.QueryRow("select * from user where id = 123456")
	users := &Users{}
	row.Scan(&users.Id, &users.Pwd, &users.Birthday, &users.Sex, &users.HeadImg)
	fmt.Println("成功》》》", users)
}
func testLogin(t *testing.T) {
	user, _ := CheckUserAndPwd("123456", "123456")
	fmt.Println("用户信息：", user)
}
func testInsertUser(t *testing.T) {
	err := InsertUser("888888", "888888", time.Now().Format("2006-01-02 15:04:05"), "男", "static/image/head/1.gif")
	if err != nil {
		fmt.Println(err)
		//return
	}
	fmt.Println("insert over")
}

func testSelectNews(t *testing.T) {
	user, _ := SelectNews(1)
	fmt.Println(user)
}
