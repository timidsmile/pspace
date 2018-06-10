package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robvdl/pongo2gin"
	"github.com/timidsmile/pspace/action/common"
	"github.com/timidsmile/pspace/action/index"
	"github.com/timidsmile/pspace/action/passport"
	"github.com/timidsmile/pspace/action/session"
	"github.com/timidsmile/pspace/action/test"
	"github.com/timidsmile/pspace/middleware"
)

func LoadRouters() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	router.HTMLRender = pongo2gin.New(
		pongo2gin.RenderOptions{
			TemplateDir: "./static/html/",
		})

	router.Static("./static", "./static")

	fmt.Println(router.HTMLRender)

	// 测试
	testGroup := router.Group("/api/test")
	{
		testGroup.GET("/welcome", test.WelcomeAction)
		testGroup.GET("/goodbye", test.GoodbyeAction)
		testGroup.GET("/testdb", test.TestdbAction)
	}

	// 首页
	indexGroup := router.Group("/api")
	{
		indexGroup.GET("/", index.IndexAction)
	}

	// session 模块
	sessionsGroup := router.Group("/api/session")
	{
		sessionsGroup.POST("/register", session.RegisterAction)
		sessionsGroup.POST("/login", session.LoginAction)
		sessionsGroup.POST("/checkLogin", session.CheckLoginAction)
	}

	// passport 模块
	passportGroup := router.Group("/api/passport").Use(middleware.CheckLogin)
	{
		passportGroup.POST("/setting", passport.SettingAction)
		passportGroup.GET("/getUserInfo", passport.GetUserInfoAction)
	}

	// 公共服务 模块
	commonGroup := router.Group("/api/common")
	{
		commonGroup.POST("/uploadFile", common.UploadFileAction)
		commonGroup.POST("/uploadAvatarImage", common.UploadAvartarImageAction)
	}

	return router
}
