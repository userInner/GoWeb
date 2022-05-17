package model

import "fmt"

//获得回复信息
func GetReplyMsg(post_id int) ([]*Reply, error) {
	sqlStr := "select reply_id,reply_user_id,reply_content,reply_time,reply_post_id from reply where reply_post_id = ?"
	rows, err := DB.Query(sqlStr, post_id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var replys []*Reply
	for rows.Next() {
		reply := &Reply{}
		rows.Scan(&reply.Reply_id, &reply.Reply_user_id, &reply.Reply_content, &reply.Reply_time, &reply.Reply_post_id)
		replys = append(replys, reply)
	}
	return replys, nil
}
