package main

import (
	"fmt"
	env "github.com/erictmc/url_shortener/api/environment"
	"github.com/erictmc/url_shortener/api/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"io/ioutil"
)

func main() {
	e := echo.New()
	appEnv := env.CreateAppEnvironment()

	if !appEnv.IsProductionEnv {
		CreateDatabaseTables(&appEnv)
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.File("/", "web/build/index.html")
	e.File("/robots.txt", "web/build/robots.txt")
	e.Static("/assets", "web/build/assets")
	e.Static("/", "web/build/")

	e.POST("/new", handlers.CreateShortUrl(appEnv))
	e.GET("/:short_url", handlers.RouteToShortUrl(appEnv))


	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", appEnv.Port)))
}

func CreateDatabaseTables(env *env.AppEnvironment) {
	env.Logger.Println("CREATING TABLES FOR DEVELOPMENT DATABASE")
	bytes, err := ioutil.ReadFile("./schema/schema.sql")
	if err != nil {
		env.Logger.Fatal(err)
	}

	query := string(bytes)
	db := env.Db.DB.Exec(query)
	if db.Error != nil {
		env.Logger.Fatal("FATAL", db.Error)
	}

	env.Logger.Println("CREATED TABLES")
}
