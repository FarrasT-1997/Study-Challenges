package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {
	// SIGNUP FOR USER AND ADMIN
	e.POST("/users/signup", controller.UserSignup)
	e.POST("/admin/signup", controller.AdminSignup)

	//LOGIN FOR USER AND ADMIN
	e.POST("/users/login", controller.UserLogin)
	e.POST("/admin/login", controller.AdminLogin)

	//AUTHORIZATION JWT
	eJwt := e.Group("/jwt")
	eJwt.Use(middleware.JWT([]byte(constant.SECRET_JWT)))

	//USER PROFILE
	eJwt.GET("/users/:id", controller.ShowUserProfile)
	eJwt.PUT("/users/:id", controller.EditUserProfile)
	eJwt.DELETE("/users/:id", controller.DeleteUser)
	eJwt.GET("users/peringkat", controller.ShowLeaderboards)

	//ADMIN PROFILE
	eJwt.GET("/admin/:id", controller.ShowAdminProfile)
	eJwt.PUT("/admin/:id", controller.EditAdminProfile)
	eJwt.DELETE("/admin/:id", controller.DeleteAdmin)

	//ADMIN FEATURES INPUT QUESTION
	eJwt.POST("admin/soal", controller.SubmitQuestionAdmin)
	eJwt.PUT("admin/soal/:soalId", controller.EditQuestion)
	eJwt.DELETE("admin/soal/:soalId", controller.DeleteQuestion)
	eJwt.GET("admin/soal/:soalId", controller.GetQuestionById)
	eJwt.GET("admin/soal/:MataPelajaranId", controller.GetQuestionByCategory)

	//ADMIN FEATURE REVIEW SUBMITTED QUESTION FROM USER
	eJwt.GET("admin/submit_soal/:kategori_materi_pelajaran_id", controller.ShowSubmittedQuestion) //SHOW QUESTIONS ARE NOTE REVIEWED BY CATEGORY
	eJwt.PUT("admin/submit_soal/:id", controller.EditSubmitQuestion)                              //CLASSIFICATION THE LEVEL OF QUESTION (EASY, MEDIUM, HARD) AND APPROVAL THE QUESTION (APPROVED OR REJECT)
	eJwt.POST("admin/submit_soal/soal/:id", controller.SendQuestion)                              //SEND SUBMITTED QUESTION TO QUESTION BOX AFTER REVIEW
	eJwt.DELETE("admin/submit_soal/:id", controller.RejectQuestion)

	//USER FEATURE SUBMIT NEW QUESTION
	eJwt.POST("users/submit_soal", controller.SubmitQuestion)

	//USER FEATURE EXERCISE
	eJwt.POST("users/:user_id/soal/:kategori_materi_pelajaran_id/set_soal", controller.GenerateRandomQuestion)
	eJwt.GET("users/:user_id/soal/:set_soal_id", controller.ShowActiveQuestion)
	eJwt.PUT("users/:user_id/soal/:set_soal_id/:nomor1/:nomor2/:nomor3/:nomor4/:nomor5".controller.AnswerQuestion)
	eJwt.GET("users/:user_id/soal/:set_soal_id/score", controller.ShowScore)
	eJwt.GET("users/:user_id/soal/:set_soal_id/solution", controller.ShowSolution)

	//GET CATEGORY
	eJwt.GET("users/:user_id/soal/:kategori_materi_pelajaran", controller.ShowCategory)
}
