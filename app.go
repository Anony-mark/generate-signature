package app

import (
	route "app/route"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	color.Green("Server starting...")

	r := gin.Default()

	route.RouteStatement(r)
	r.Run(":7777")
}
