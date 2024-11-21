package route

import (
	"app/controller"

	"github.com/gin-gonic/gin"
)

// MemberRoute create route
func RouteStatement(r *gin.Engine) {
	api := r.Group("/api")

	//Expense

	api.POST("/generate-signature", controller.GenerateSignature())

}
