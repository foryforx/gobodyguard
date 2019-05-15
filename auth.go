package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karuppaiah/gobodyguard/app"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	// HealthCheck
	// TODO: implement proper healthcheck with DB and redis(if applicable)
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return r

}

func main() {
	r := setupRouter()
	db := app.GetDBInstance()
	// Dependency injection and route setting
	app.NewAuthHandler(r, db.DB)
	r.Run(":5422")
}
