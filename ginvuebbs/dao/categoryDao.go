package dao

import (
	"ginAndVueBBS/common"
	"ginAndVueBBS/model"
)

func CreateCategory(category *model.Category) error {
	sqlStr := "insert into category(name,create_at,update_at) values(?,?,?)"

	_, err := common.DB.Exec(sqlStr, category.Name, category.CreateAt, category.UpdateAt)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCategory(categoryId int) error {
	sqlStr := "delete from category where id = ?"
	_, err := common.DB.Exec(sqlStr, categoryId)
	if err != nil {
		return err
	}
	return nil
}

func UpdateCategory(category *model.Category) error {
	sqlStr := "update category set name = ?,update_at = ? where id = ?"
	strTime := category.UpdateAt.Format("2006-01-06 15:04:05")
	_, err := common.DB.Exec(sqlStr, category.Name, strTime, category.ID)
	if err != nil {
		return err
	}
	return nil
}

func GetCategory(id int) (*model.Category, error) {
	sqlStr := "select id,name,create_at,update_at from category where id = ?"
	row := common.DB.QueryRow(sqlStr, id)
	var tempCategory model.Category
	err := row.Scan(&tempCategory.ID, &tempCategory.Name, &tempCategory.CreateAt, &tempCategory.UpdateAt)
	if err != nil {
		return nil, err
	}
	return &tempCategory, nil
}
