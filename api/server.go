package api

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gotool/api/auth"
	"gotool/api/controllers"
	"gotool/auto"
	"gotool/config"
	"time"
)


func init() {
	auto.Load()
}

func Run() {

	gin.DisableConsoleColor()

	// Logging to a file.
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	//gin.SetMode(gin.ReleaseMode)

	//系统初始化
	r := gin.Default()
	r.Static("/statics", "./statics")
	r.LoadHTMLGlob("templates/*")

	r.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowOrigins:     []string{"http://localhost:8005"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//登录
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "未登录用户",
		})
	})


	//用户
	user := r.Group("/v1/api/user")
	user.POST("/register", controllers.Register)
	user.POST("/login", controllers.Login)
	user.POST("/data", controllers.Data)
	user.GET("/data", controllers.Data)
	authorize := r.Group("/user", auth.JWTAuth())
	{
		authorize.GET("ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "登录用户",
			})
		})
	}
	r.Run(fmt.Sprintf(":%d", config.PORT))
}
