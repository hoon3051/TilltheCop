package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/hoon3051/TilltheCop/controller"
	"github.com/hoon3051/TilltheCop/middleware"
)

var OauthController controller.OauthController
var ProfileController controller.ProfileController

func OauthRouter(router *gin.Engine) {
	// 세션 스토어 설정
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("oauth_session", store))

	router.GET("/oauth/google/login", OauthController.GoogleLogin)
	router.GET("/oauth/google/callback", OauthController.GoogleCallback)
	router.POST("/oauth/register", OauthController.Register)
}

func ProfileRouter(router *gin.Engine) {
	router.Use(middleware.AuthToken())
	router.GET("/profile/", ProfileController.GetProfile)
	router.PUT("/profile/", ProfileController.UpdateProfile)
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	OauthRouter(router)
	ProfileRouter(router)

	return router
}
