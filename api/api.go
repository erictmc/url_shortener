package main

import (
	"fmt"
	env "github.com/erictmc/url_shortener/api/environment"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"io/ioutil"
	"net/http"
)

func main() {
	e := echo.New()
	appEnv := env.CreateAppEnvironment()

	if !appEnv.IsProductionEnv {
		CreateTables(&appEnv)
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})


	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", appEnv.Port)))
}

func CreateTables(env *env.AppEnvironment) {
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
