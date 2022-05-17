package model

import (
	"fmt"
)

func AddSession(session *Sessions) error {
	sqlStr := "insert into sessions(sessions_id,username,user_id) values(?,?,?)"
	_, err := DB.Exec(sqlStr, session.Sessions_id, session.Username, session.User_id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func DelSession(sess_id string) error {
	sqlStr := "delete from sessions where sessions_id = ?"
	_, err := DB.Exec(sqlStr, sess_id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func QuerySession(sess_id string) (*Sessions, error) {
	sqlStr := "select sessions_id,username,user_id from sessions where sessions_id = ?"
	preSql, err := DB.Prepare(sqlStr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	row := preSql.QueryRow(sess_id)
	sess := &Sessions{}
	row.Scan(&sess.Sessions_id, &sess.Username, &sess.User_id)
	return sess, nil
}

//根据uuid查询所有
func SessionIdGet(uuidStr string) (*Sessions, error) {
	sqlStr := "select sessions_id,username,user_id from sessions where sessions_id = ?"
	row := DB.QueryRow(sqlStr, uuidStr)

	session := &Sessions{
		Sessions_id: "",
		Username:    "",
		User_id:     0,
	}

	if row != nil {
		row.Scan(&session.Sessions_id, &session.Username, &session.User_id)
		return session, nil
	}
	fmt.Println(err)
	return nil, err
}
