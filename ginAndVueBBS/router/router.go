package router

import (
	"ginAndVueBBS/controller"
	"ginAndVueBBS/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.UserInfo)

	categoryRouters := r.Group("/categorys")
	{
		var iCategoryController controller.ICategoryController
		iCategoryController = controller.CategoryController{}
		categoryRouters.POST("", iCategoryController.Create)
		categoryRouters.GET("/:id", iCategoryController.Show)
		categoryRouters.PUT("/:id", iCategoryController.Update)
		categoryRouters.DELETE("/:id", iCategoryController.Delete)
	}

	postRouters := r.Group("/posts", middleware.AuthMiddleware())
	{
		var iPostController controller.ICategoryController
		iPostController = controller.PostController{}
		postRouters.POST("", iPostController.Create)
		postRouters.GET("/:id", iPostController.Show)
		postRouters.PUT("/:id", iPostController.Update)
		postRouters.DELETE("/:id", iPostController.Delete)
		postRouters.POST("/page/list", iPostController.PageList)
	}
	return r
}
