package dao

import (
	"ginAndVueBBS/common"
	"ginAndVueBBS/model"
)

func IsTelephoneExist(telephone string) bool {
	sqlStr := "select user_id,name,telephone,password from users where telephone = ?"
	row := common.DB.QueryRow(sqlStr, telephone)
	var user model.User
	row.Scan(&user.UserId, &user.Name, &user.Telephone, &user.Password)
	if user.UserId != 0 {
		return true
	}
	return false
}

func GetTelephoneUser(telephone string) model.User {
	sqlStr := "select user_id,name,telephone,password from users where telephone = ?"
	row := common.DB.QueryRow(sqlStr, telephone)
	var user model.User
	err := row.Scan(&user.UserId, &user.Name, &user.Telephone, &user.Password)
	if err != nil {
		panic("getuser scan failed,err: " + err.Error())
	}
	return user
}

func GetUserIdUser(userId int) model.User {
	sqlStr := "select user_id,name,telephone,password from users where user_id = ?"
	row := common.DB.QueryRow(sqlStr, userId)
	var user model.User
	row.Scan(&user.UserId, &user.Name, &user.Telephone, &user.Password)
	return user
}

//注册用户
func RegisterUser(name, telephone, password string) bool {
	sqlStr := "insert into users(name,telephone,password) values(?,?,?)"
	res, _ := common.DB.Exec(sqlStr, name, telephone, password)
	ids, _ := res.RowsAffected()
	if ids == 0 {
		return false
	}
	return true
}
