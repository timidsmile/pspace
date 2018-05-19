package router

import (
	"github.com/gin-gonic/gin"
	"github.com/timidsmile/pspace/action/passport"
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
	sessionsGroup := router.Group("/session")
	{
		sessionsGroup.POST("/register", session.RegisterAction)
		sessionsGroup.POST("/login", session.LoginAction)
	}

	// passport 模块
	passportGroup := router.Group("/passport")
	{
		passportGroup.POST("/setting", passport.SettingAction)
	}

	return router
}
