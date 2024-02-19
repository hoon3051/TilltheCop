package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/hoon3051/TilltheCop/server/controller"
	"github.com/hoon3051/TilltheCop/server/middleware"
	"github.com/gin-contrib/cors"
	"time"
)

var OauthController controller.OauthController
var ProfileController controller.ProfileController
var MapController controller.MapController

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

	router.GET("/profile", ProfileController.GetProfile)
	router.PUT("/profile", ProfileController.UpdateProfile)
}

func MapRouter(router *gin.Engine) {
	router.Use(middleware.AuthToken())

	router.POST("/map", MapController.GetMap)
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // 클라이언트의 도메인 명시
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // withCredentials 요청 허용
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(config))

	OauthRouter(router)
	ProfileRouter(router)
	MapRouter(router)

	return router
}
