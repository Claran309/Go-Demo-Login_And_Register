package main

import (
	"GoGin/api/dao/mysql"
	handlers2 "GoGin/api/handlers"
	"GoGin/api/services"
	"GoGin/internal/config"
	"GoGin/internal/middleware"
	"GoGin/internal/util/jwt_util"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	//========================初始化===========================
	// 数据层依赖
	// userRepo := memory.NewMemoryUserRepository() 内存记忆储存信息
	db, err := mysql.InitMysql(cfg)
	if err != nil {
		log.Fatal(err)
	}
	userRepo := mysql.NewMysqlUserRepo(db)
	courseRepo := mysql.NewMysqlCourseRepo(db)
	// JWT工具
	jwtUtil := jwt_util.NewJWTUtil(cfg)
	// 业务逻辑层依赖
	userService := services.NewUserService(userRepo, jwtUtil)
	courseService := services.NewCourseService(courseRepo)
	// 处理器层依赖
	userHandler := handlers2.NewUserHandler(userService)
	courseHandler := handlers2.NewCourseHandler(courseService)
	//创建中间件
	jwtMiddleware := middleware.NewJWTMiddleware(jwtUtil)

	r := gin.Default()

	//=====================注册和登录路由========================
	user := r.Group("/user")
	user.POST("/register", userHandler.Register)
	user.POST("/login", userHandler.Login)
	user.POST("/refresh", userHandler.Refresh)
	user.GET("/info", jwtMiddleware.JWTAuthentication(), userHandler.InfoHandler)

	//=====================课程相关路由==========================
	course := r.Group("/course")
	course.Use(jwtMiddleware.JWTAuthentication())
	//获取课程列表
	course.GET("/info", courseHandler.Info)
	//获取已选课程列表
	course.GET("/enrollment", courseHandler.EnrollmentInfo)
	//选课
	course.POST("/pick", courseHandler.PickCourse)
	//退课
	course.POST("/drop", courseHandler.DropCourse)
	//新增课程 (admin)
	course.POST("/add/course", jwtMiddleware.JWTAuthentication(), jwtMiddleware.JWTAuthorization(), courseHandler.AddCourse)

	err = r.Run()
	if err != nil {
		panic("Failed to start Gin server: " + err.Error())
	}
}

/*
各路由请求体应有数据：

===================="/user"======================
"/register":
	Body:
		username
		password
		email
		role (admin/user)

"/login":
	Body:
		login_key
		password

"/refresh":
	Body:
		refresh_token

"/info":
	Header:
		Authorization : Bearer <Token>

===================="/course"=====================
"/pick"
	Header:
		Authorization : Bearer <Token>
	Body:
		course_id

"/drop":
	Header:
		Authorization : Bearer <Token>
	Body:
		course_id

"/add/course":
	Header:
		Authorization : Bearer <Token>
	Body:
		name
		capital

"/info":
	nil

“/enrollment”:
	Header:
		Authorization : Bearer <Token>
*/
