package main

import (
	"log"

	"github.com/bllyanos/saved-actions/internals/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	routes.SetupIndexRoute(app)
	routes.SetupActionsRoute(app)

	if err := app.Run(":8080"); err != nil {
		log.Fatal("run error", err.Error())
	}
}
