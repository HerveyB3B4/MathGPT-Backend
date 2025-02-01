package router

import (
	"MATHB/app/controllers"
	testcontroller "MATHB/app/controllers/testController"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	const pre = "/api"
	const test = "/test"
	const auth = "/auth"
	api := r.Group(pre)
	{
		authGroup := api.Group(auth)
		{
			authGroup.POST("/login", controllers.LoginByIDHandler)
		}
		testGroup := api.Group(test)
		{
			testGroup.GET("/alive", testcontroller.TestHandler)
		}
	}

}
