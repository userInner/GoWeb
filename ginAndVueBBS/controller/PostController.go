package controller

import (
	"fmt"
	"ginAndVueBBS/dao"
	"ginAndVueBBS/model"
	"ginAndVueBBS/response"
	"ginAndVueBBS/vo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"strconv"
	"time"
)

type IPostController interface {
	RestController
}

type PostController struct {
}

func (p PostController) PageList(c *gin.Context) {
	pageNum, err := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	if err != nil {
		log.Println(err)
		response.Fail(c, nil, "参数错误")
		return
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	if err != nil {
		log.Println(err)
		response.Fail(c, nil, "参数错误")
		return
	}

	posts, err := dao.GetPageList(pageNum, pageSize)
	if err != nil {
		log.Println(err)
		response.Fail(c, nil, "获取帖子失败")
		return
	}
	total := dao.PostTotal()
	response.Success(c, gin.H{"posts": posts, "total": total}, "获取成功")
}

func (p PostController) Create(c *gin.Context) {
	var requestPost vo.CreatePostRequest
	if err := c.ShouldBind(&requestPost); err != nil {
		log.Println(err.Error())
		response.Fail(c, nil, "数据验证错误，分类名称必填")
		return
	}

	user, _ := c.Get("user")
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	strNowTime, _ := time.Parse("2006-01-02 15:04:05", nowTime)
	post := model.Post{
		ID:         uuid.New(),
		UserId:     user.(model.User).UserId,
		CategoryId: requestPost.CategoryId,
		Title:      requestPost.Title,
		HeadImg:    requestPost.HeadImg,
		Content:    requestPost.Content,
		CreateAt:   strNowTime,
		UpdateAt:   strNowTime,
	}
	err := dao.CreatePost(post)
	if err != nil {
		log.Println(err)
		panic(err)
		return
	}
	response.Success(c, nil, "创建帖子成功")
}

func (p PostController) Delete(c *gin.Context) {
	post_id := c.Params.ByName("id")
	parse, _ := uuid.Parse(post_id)
	post, err := dao.GetPost(parse)
	if post.Title == "" || err != nil {
		response.Fail(c, nil, "文章不存在")
		log.Println(err)
		return
	}

	user, _ := c.Get("user")
	userId := user.(model.User).UserId
	if userId != post.UserId {
		response.Fail(c, nil, "请勿非法操作，该文章不属于你")
		return
	}
	err = dao.DeletePost(post_id)
	if err != nil {
		log.Println(err)
		response.Fail(c, nil, fmt.Sprint(err))
		return
	}

	response.Success(c, gin.H{"post": post}, "删除成功")
}

func (p PostController) Update(c *gin.Context) {
	post_id := c.Params.ByName("id")
	parse, _ := uuid.Parse(post_id)
	post, err := dao.GetPost(parse)
	if post.Title == "" || err != nil {
		response.Fail(c, nil, "文章不存在")
		log.Println(err)
		panic(err)
		return
	}

	user, _ := c.Get("user")
	userId := user.(model.User).UserId
	if userId != post.UserId {
		response.Fail(c, nil, "请勿非法操作，该文章不属于你")
		return
	}

	//传过来的数据
	var requestPost vo.CreatePostRequest
	if err := c.ShouldBind(&requestPost); err != nil {
		response.Fail(c, nil, "数据验证错误")
		log.Println(err)
		return
	}
	post.Title = requestPost.Title
	post.HeadImg = requestPost.HeadImg
	post.Content = requestPost.Content
	post.CategoryId = requestPost.CategoryId
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	strNowTime, _ := time.Parse("2006-01-02 15:04:05", nowTime)
	post.UpdateAt = strNowTime
	err = dao.UpdatePost(post)
	if err != nil {
		log.Println(err)
		panic(err)
		return
	}

	response.Success(c, gin.H{"post": post}, "更新成功")
}

func (p PostController) Show(c *gin.Context) {
	post_id := c.Params.ByName("id")
	parse, _ := uuid.Parse(post_id)
	post, err := dao.GetPost(parse)
	if err != nil {
		log.Println(err)
		response.Fail(c, nil, "文章不存在")
		return
	}
	response.Success(c, gin.H{"post": post}, "获取成功")
}

func NewIPostController() IPostController {
	return PostController{}
}
