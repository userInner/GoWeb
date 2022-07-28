package dao

import (
	"ginAndVueBBS/common"
	"ginAndVueBBS/model"
	"github.com/google/uuid"
)

func CreatePost(post model.Post) error {
	sqlStr := "insert into posts(id,user_id,category_id,title,head_img,content,create_at,update_at) values(?,?,?,?,?,?,?,?)"
	_, err := common.DB.Exec(sqlStr, post.ID, post.UserId, post.CategoryId, post.Title, post.HeadImg, post.Content, post.CreateAt, post.UpdateAt)

	if err != nil {
		return err
	}
	return nil
}

func DeletePost(post_id string) error {
	sqlStr := "delete from posts where id = ?"
	_, err := common.DB.Exec(sqlStr, post_id)

	if err != nil {
		return err
	}
	return nil
}

func UpdatePost(post *model.Post) error {
	sqlStr := "update posts set user_id = ?,category_id = ?,title = ?,head_img = ?,content = ?,update_at= ? where id = ?"
	_, err := common.DB.Exec(sqlStr, post.UserId, post.CategoryId, post.Title, post.HeadImg, post.Content, post.UpdateAt, post.ID)
	if err != nil {
		return err
	}
	category, err := GetCategory(int(post.CategoryId))
	if err != nil {
		return err
	}
	post.Category = category
	return nil
}

func GetPost(postId uuid.UUID) (*model.Post, error) {
	sqlStr := "select id,user_id,category_id,title,head_img,content,create_at,update_at from posts where id = ?"
	row := common.DB.QueryRow(sqlStr, postId)

	var post model.Post

	err := row.Scan(&post.ID, &post.UserId, &post.CategoryId, &post.Title, &post.HeadImg, &post.Content, &post.CreateAt, &post.UpdateAt)
	if err != nil {
		return nil, err
	}
	post.Category, err = GetCategory(int(post.CategoryId))
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func GetPageList(pageNum, pageSize int) (*[]model.Post, error) {
	sqlStr := "select id,user_id,category_id,title,head_img,content,create_at,update_at from posts order by create_at desc limit ?,?"
	rows, _ := common.DB.Query(sqlStr, (pageNum-1)*pageSize, pageSize)

	var posts []model.Post
	for rows.Next() {
		var post model.Post
		err := rows.Scan(&post.ID, &post.UserId, &post.CategoryId, &post.Title, &post.HeadImg, &post.Content, &post.CreateAt, &post.UpdateAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return &posts, nil
}

func PostTotal() int {
	sqlStr := "select count(*) from posts"
	row := common.DB.QueryRow(sqlStr)
	var count = 0
	row.Scan(&count)
	return count
}
