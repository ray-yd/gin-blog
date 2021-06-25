package routes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ray-yd/gin-blog/api/v1"
	"github.com/ray-yd/gin-blog/middleware"
	"github.com/ray-yd/gin-blog/utils"
)

// InitRouter 初始化 Router
func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.MaxMultipartMemory = 8 << 20

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// UserModel Router接口
		auth.PUT("/user/:id", v1.EditUser)
		auth.DELETE("/user/:id", v1.DeleteUser)

		// 分類 Model
		auth.POST("/addCategory", v1.AddCategory)
		auth.PUT("/Category/:id", v1.EditCategory)
		auth.DELETE("/Category/:id", v1.DeleteCategory)
		// 文章 Model
		auth.POST("/addArticle", v1.AddArticle)
		auth.PUT("/Article/:id", v1.EditArticle)
		auth.DELETE("/Article/:id", v1.DeleteArticle)
	}
	routerV1 := r.Group("api/v1")
	{
		routerV1.POST("/addUser", v1.AddUser)

		routerV1.GET("/userList", v1.GetUsers)
		routerV1.GET("/CategoryList", v1.GetCategoryList)
		routerV1.GET("/ArticleList", v1.GetArticleList)
		routerV1.GET("/CateArticleList/:id", v1.GetCategoryArticle)
		routerV1.GET("/ArticleInfo/:id", v1.GetArticleInfo)
		routerV1.POST("/login", v1.Login)

		// 上傳接口
		routerV1.POST("/upload", v1.Upload)
	}
	r.Run(utils.HttpPort)
}
