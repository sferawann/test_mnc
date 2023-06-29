package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sferawann/test_mnc/controller"
	"github.com/sferawann/test_mnc/middleware"
)

func NewRouter(userCon *controller.UserCon, accCon *controller.AccountCon, hisCon *controller.HistoryCon, traCon *controller.TransferCon, sesCon *controller.SessionCon, authCon *controller.AuthCon) *gin.Engine {
	r := gin.Default()

	r.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := r.Group("/api")

	authRouter := router.Group("/auth")
	{
		authRouter.POST("/", authCon.Login)
		authRouter.Use(middleware.AuthMiddleware())
		{
			authRouter.POST("/logout", authCon.Logout)
		}
	}

	usersRouter := router.Group("/user")
	{
		usersRouter.GET("/", userCon.FindAll)
		usersRouter.POST("/", userCon.Create)
		usersRouter.GET("/:id", userCon.FindByID)
		usersRouter.GET("/username/:username", userCon.FindByUsername)
		usersRouter.PUT("/:id", userCon.Update)
		usersRouter.DELETE("/:id", userCon.Delete)
		usersRouter.Use(middleware.AuthMiddleware())
		{
			usersRouter.GET("/get", userCon.Get)
		}
	}

	accRouter := router.Group("/account")
	{
		accRouter.GET("/", accCon.FindAll)
		accRouter.POST("/", accCon.Create)
		accRouter.GET("/:id", accCon.FindByID)
		accRouter.PUT("/:id", accCon.Update)
		accRouter.DELETE("/:id", accCon.Delete)
		accRouter.Use(middleware.AuthMiddleware())
		{
			accRouter.GET("/get", accCon.GetByUserID)
		}
	}

	hisRouter := router.Group("/history")
	{
		hisRouter.GET("/", hisCon.FindAll)
		hisRouter.POST("/", hisCon.Create)
		hisRouter.GET("/:id", hisCon.FindByID)
		hisRouter.PUT("/:id", hisCon.Update)
		hisRouter.DELETE("/:id", hisCon.Delete)
	}

	traRouter := router.Group("/transfer")
	{
		traRouter.GET("/", traCon.FindAll)
		traRouter.POST("/", traCon.Create)
		traRouter.GET("/:id", traCon.FindByID)
		traRouter.PUT("/:id", traCon.Update)
		traRouter.DELETE("/:id", traCon.Delete)
	}

	sesRouter := router.Group("/session")
	{
		sesRouter.GET("/", sesCon.FindAll)
		sesRouter.POST("/", sesCon.Create)
		sesRouter.GET("/:id", sesCon.FindByID)
		sesRouter.PUT("/:id", sesCon.Update)
		sesRouter.DELETE("/:id", sesCon.Delete)
	}

	return r
}
