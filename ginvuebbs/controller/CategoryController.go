package controller

import (
	"ginAndVueBBS/dao"
	"ginAndVueBBS/model"
	"ginAndVueBBS/response"
	"ginAndVueBBS/vo"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
}

func (cate CategoryController) PageList(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

//实例化
func NewCategoryController() ICategoryController {
	return CategoryController{}
}

func (cate CategoryController) Create(c *gin.Context) {
	var requestCategory vo.CreateCategoryRequest
	if err := c.ShouldBind(&requestCategory); err != nil {
		response.Fail(c, nil, "数据验证错误，分类名称必填")
		return
	}
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	strNowTime, _ := time.Parse("2006-01-02 15:04:05", nowTime)
	createCategory := &model.Category{
		ID:       0,
		Name:     requestCategory.Name,
		CreateAt: strNowTime,
		UpdateAt: strNowTime,
	}
	err := dao.CreateCategory(createCategory)
	if err != nil {
		panic(err)
	}
	response.Success(c, gin.H{"catagory": createCategory}, "创建成功")
}

func (cate CategoryController) Delete(c *gin.Context) {
	categoryId, _ := strconv.Atoi(c.Params.ByName("id"))

	if err := dao.DeleteCategory(categoryId); err != nil {
		response.Fail(c, nil, "删除失败，请重试")
		return
	}

	response.Success(c, nil, "删除分类成功")
}

func (cate CategoryController) Update(c *gin.Context) {
	var requestCategory vo.CreateCategoryRequest
	if err := c.ShouldBind(&requestCategory); err != nil {
		response.Fail(c, nil, "数据验证错误，分类名称必填")
		return
	}

	categoryId, _ := strconv.Atoi(c.Params.ByName("id"))
	isCategory, _ := dao.GetCategory(categoryId)
	if isCategory.ID == 0 {
		response.Fail(c, nil, "分类不存在")
		return
	}
	isCategory.Name = requestCategory.Name
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	isCategory.UpdateAt, _ = time.Parse("2006-01-02 15:04:05", nowTime)
	err := dao.UpdateCategory(isCategory)
	if err != nil {
		log.Println(err)
		response.Fail(c, nil, "分类修改失败")
		return
	}
	newCategory, _ := dao.GetCategory(categoryId)
	response.Success(c, gin.H{"category": newCategory}, "修改成功")
}

func (cate CategoryController) Show(c *gin.Context) {
	categoryId, _ := strconv.Atoi(c.Params.ByName("id"))

	isCatagory, _ := dao.GetCategory(categoryId)
	if isCatagory.ID == 0 {
		response.Fail(c, nil, "分类不存在")
	}

	response.Success(c, gin.H{"category": isCatagory}, "")
}
