package main

import (
	"GoGin/internal/api/handlers"
	"GoGin/internal/config"
	"GoGin/internal/middleware"
	"GoGin/internal/repository/mysql"
	"GoGin/internal/services"
	"GoGin/internal/util/jwt_util"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	// 初始化依赖
	// 数据层依赖
	// userRepo := memory.NewMemoryUserRepository() 内存记忆储存信息
	db, err := mysql.InitMysql(cfg)
	if err != nil {
		log.Fatal(err)
	}
	userRepo := mysql.NewMysqlUserRepo(db)
	// JWT工具
	jwtUtil := jwt_util.NewJWTUtil(cfg)
	// 业务逻辑层依赖
	userService := services.NewUserService(userRepo, jwtUtil)
	// 处理器层依赖
	userHandler := handlers.NewUserHandler(userService)

	r := gin.Default()

	//注册和登录路由
	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	//受保护路由：JWT中间件判断访问权限
	protected := r.Group("/protected")
	//创建中间件
	jwtMiddleware := middleware.NewJWTMiddleware(jwtUtil)

	protected.Use(jwtMiddleware.JWTAuthorizationMiddleware())
	{
		protected.GET("/info", userHandler.InfoHandler)
	}

	err = r.Run()
	if err != nil {
		panic("Failed to start Gin server: " + err.Error())
	}
}

/*
注册时：
前端JSON:
username
password
email

登录时：
前端JSON：
loginkey
password
*/
