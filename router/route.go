package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hoon3051/TilltheCop/controller"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

var OauthController controller.OauthController

func OauthRouter(router *gin.Engine) {
	// 세션 스토어 설정
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("oauth_session", store))

	router.GET("/oauth/google/login", OauthController.GoogleLogin)
	router.GET("/oauth/google/callback", OauthController.GoogleCallback)
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	OauthRouter(router)

	return router
}
