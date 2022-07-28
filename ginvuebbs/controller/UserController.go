package controller

import (
	"ginAndVueBBS/common"
	"ginAndVueBBS/dao"
	"ginAndVueBBS/dto"
	"ginAndVueBBS/model"
	"ginAndVueBBS/response"
	"ginAndVueBBS/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Register(c *gin.Context) {
	var requestUser model.User
	c.Bind(&requestUser)
	//获取参数
	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password

	//验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号不能少于11位")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于六位")
		return
	}
	if len(name) == 0 { //无名，则随机给一个字符串
		name = utils.RandomString(10)
	}

	//查看手机号是否存在
	if dao.IsTelephoneExist(telephone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
		return
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Fail(c, nil, "系统错误")
		return
	}
	if dao.RegisterUser(name, telephone, string(hashPassword)) {
		newUser := dao.GetTelephoneUser(telephone)
		log.Printf("%v %v %v", name, telephone, password)
		//返回token
		tokenString, err := common.ReleaseToken(newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "系统异常",
			})
			log.Printf("token generate error: %v", err)
			return
		}
		response.Success(c, gin.H{"token": tokenString}, "注册成功")
	}
}

func Login(c *gin.Context) {
	var requestUser model.User
	c.Bind(&requestUser)
	//获取参数
	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password

	//验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号不能少于11位")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于六位")
		return
	}
	if len(name) == 0 { //无名，则随机给一个字符串
		name = utils.RandomString(10)
	}
	user := dao.GetTelephoneUser(telephone)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "密码错误",
		})
		return
	}
	tokenString, err := common.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "系统异常",
		})
		log.Printf("token generate error: %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": gin.H{
			"token": tokenString,
		},
	})
}

func UserInfo(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"user": dto.ToUserDto(user.(model.User)),
		},
	})
}
