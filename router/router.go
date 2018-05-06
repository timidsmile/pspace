package router

import (
	"github.com/timidsmile/pspace/action/test"
	"github.com/timidsmile/pspace/action/session"
	"github.com/gin-gonic/gin"
)

func LoadRouters() *gin.Engine{

	router := gin.New()
	router.Use(gin.Recovery())

	// 测试
	testGroup := router.Group("/test")
	{
		testGroup.GET("/welcome", test.WelcomeAction)
		testGroup.GET("/goodbye", test.GoodbyeAction)
	}

	// session 模块
	sessions := router.Group("/session")
	{
		sessions.GET("/register", session.RegisterAction)
		sessions.GET("/login", session.LoginAction)
	}

	return router;
}
