package router

import (
	"github.com/gin-gonic/gin"
	"github.com/timidsmile/pspace/action/session"
	"github.com/timidsmile/pspace/action/test"
)

func LoadRouters() *gin.Engine {

	router := gin.New()
	router.Use(gin.Recovery())

	// 测试
	testGroup := router.Group("/test")
	{
		testGroup.GET("/welcome", test.WelcomeAction)
		testGroup.GET("/goodbye", test.GoodbyeAction)
		testGroup.GET("/testdb", test.TestdbAction)
	}

	// session 模块
	sessions := router.Group("/session")
	{
		sessions.POST("/register", session.RegisterAction)
		sessions.POST("/login", session.LoginAction)
	}

	return router
}
