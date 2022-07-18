package server

import (
	"go-next/controller"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	//ginでは最初にgin.Default()関数で*Engineというインスタンスを生成します。
	//*EngineにはEndpoint、Middleware、その他Webページ用のTemplateやそこで使われるfuncMapなど様々なものを登録しておくことができます。

	r := gin.Default()
	r.GET("/", controller.ShowAllPost)
	r.GET("/show/:id", controller.ShowOnePost)
	r.GET("/create", controller.ShowCreatePost)
	r.POST("/create", controller.CreatePost)
	r.GET("/edit/:id", controller.ShowEditPost)
	r.POST("/edit", controller.EditPost)
	r.GET("/delete/:id", controller.ShowCheckDeletePost)
	r.POST("/delete", controller.DeletePost)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}
