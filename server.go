package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jessemillar/dunn/controllers"
	"github.com/jessemillar/health"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	log.Println("Configuring server")

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())

	e.GET("/health", echo.WrapHandler(http.HandlerFunc(health.Check)))
	e.PUT("/v1/dunn", controllers.ArchiveCard)
	e.PUT("/v1/dupes", controllers.KillDupes)

	// TODO Fail if $PORT is not set
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
