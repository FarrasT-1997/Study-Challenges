package route

import (
	"SC/constant"
	"SC/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {
	// SIGNUP FOR USER AND ADMIN
	e.POST("/users/signup", controller.UserSignup)  //OK
	e.POST("/admin/signup", controller.AdminSignup) //OK

	//LOGIN FOR USER AND ADMIN
	e.POST("/users/login", controller.UserLogin)  //OK
	e.POST("/admin/login", controller.AdminLogin) //OK

	//AUTHORIZATION JWT
	eJwt := e.Group("")
	eJwt.Use(middleware.JWT([]byte(constant.SECRET_JWT)))

	//LOGOUT FOR USER AND ADMIN
	eJwt.PUT("/users/:userId/logout", controller.UserLogout)   //OK
	eJwt.PUT("/admin/:adminId/logout", controller.AdminLogout) //OK

	//USER PROFILE
	eJwt.GET("/users/:id", controller.ShowUserProfile) //OK
	eJwt.PUT("/users/:id", controller.EditUserProfile) //OK
	eJwt.GET("/users", controller.ShowLeaderboards)    //OK

	//ADMIN PROFILE
	eJwt.GET("/admin/:id", controller.ShowAdminProfile) //OK
	eJwt.PUT("/admin/:id", controller.EditAdminProfile) //OK

	//ADMIN FEATURES INPUT QUESTION
	eJwt.POST("/admin/soal", controller.SubmitQuestionAdmin)                         //OK
	eJwt.PUT("/admin/soal/:soalId", controller.EditQuestion)                         //OK
	eJwt.DELETE("/admin/soal/:soalId", controller.DeleteQuestion)                    //OK
	eJwt.GET("/admin/soal/:soalId", controller.GetQuestionById)                      //OK
	eJwt.GET("/admin/soal/mapel/:MataPelajaranId", controller.GetQuestionByCategory) //OK

	//ADMIN FEATURE REVIEW SUBMITTED QUESTION FROM USER
	eJwt.GET("/admin/submit_soal/:kategori_materi_pelajaran_id", controller.ShowSubmittedQuestion) //OK  //SHOW QUESTIONS ARE NOTE REVIEWED BY CATEGORY
	eJwt.PUT("/admin/submit_soal/approval/:id", controller.EditSubmitQuestion)                     //OK //APPROVAL THE QUESTION (APPROVED OR REJECT)

	//USER FEATURE SUBMIT NEW QUESTION
	eJwt.POST("/users/submit_soal", controller.SubmitQuestion) //OK

	//USER FEATURE EXERCISE
	eJwt.POST("/users/:user_id/soal", controller.GenerateRandomQuestion)            //OK
	eJwt.GET("/users/:user_id/soal/:set_soal_id", controller.ShowActiveQuestion)    //OK
	eJwt.PUT("/users/:user_id/soal/:set_soal_id", controller.AnswerQuestion)        //OK
	eJwt.GET("/users/:user_id/soal/:set_soal_id/solution", controller.ShowSolution) //OK
}
