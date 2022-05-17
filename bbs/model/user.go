package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//校验账户和密码
func CheckUserAndPwd(userName string, pwd string) (*Users, error) {
	sqlStr := "select * from user where userName = ? and pwd = ?"
	row := DB.QueryRow(sqlStr, userName, pwd)
	users := &Users{}
	err := row.Scan(&users.Id, &users.UserName, &users.Pwd, &users.Birthday, &users.Sex, &users.HeadImg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return users, nil
}

//插入数据信息
func InsertUser(userName string, pwd, birthday, sex, headImg string) error {
	sqlStr := "insert into user(userName,pwd,birthday,sex,headImg) values(?,?,?,?,?)"

	_, err := DB.Exec(sqlStr, userName, pwd, birthday, sex, headImg)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//校验用户是否存在
func CheckUser(userName string) (bool, error) {
	sqlStr := "select * from user where userName = ?"
	row := DB.QueryRow(sqlStr, userName)
	users := &Users{}
	err := row.Scan(&users.UserName, &users.Pwd, &users.Birthday, &users.Sex, &users.HeadImg)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}

//查询用户
func SelectNews(userId int) (*Users, error) {
	sqlStr := "select id,userName,pwd,birthday,sex,headImg from user where id = ?"
	row := DB.QueryRow(sqlStr, userId)
	if row != nil {
		user := &Users{
			Id:       0,
			UserName: "",
			Pwd:      "",
			Birthday: "",
			Sex:      "",
			HeadImg:  "",
		}
		row.Scan(&user.Id, &user.UserName, &user.Pwd, &user.Birthday, &user.Sex, &user.HeadImg)
		return user, nil
	}
	fmt.Println(err)
	return nil, err

}
