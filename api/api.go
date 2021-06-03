package main

import (
	"fmt"
	env "github.com/erictmc/url_shortener/api/environment"
	"github.com/erictmc/url_shortener/api/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"io/ioutil"
)

// Entry point for the code
// The code in this project is broken into primary categories:
//  - handlers: used to handle data from incoming web requests, similar to controllers.
//  - models: interfaces for interacting with any data stores.
func main() {
	e := echo.New()
	appEnv := env.CreateAppEnvironment()

	if !appEnv.IsProductionEnv {
		CreateDatabaseTables(&appEnv)
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.File("/", "web/build/index.html")
	e.Static("/assets", "web/build/assets")
	e.Static("/", "web/build/")

	// Specified to prevent name collisions with short urls
	e.File("/robots.txt", "web/build/robots.txt")

	e.POST("/url/new", handlers.CreateShortUrl(appEnv))
	e.GET("/:short_url", handlers.RouteToShortUrl(appEnv))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", appEnv.Port)))
}

// CreateDatabaseTables sets up the database schema when in local development mode.
// A basic SQL file was used here, instead of an ORM, for readability for the reviewer.
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
