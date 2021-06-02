package environment

import (
	"fmt"
	"github.com/erictmc/url_shortener/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

const (
	PostEnvironmentVariable  = "PORT"
	DBUrlEnvironmentVariable = "DATABASE_URL"
	AllowedOriginsEnvVariable = "ALLOWED_ORIGINS"
)

type AppEnvironment struct {
	Db     models.AppDB
	Logger *log.Logger
	Port   string
	IsProductionEnv bool
}

func (ae *AppEnvironment) BuildShortUrl(shortUrl string) string {
	if !ae.IsProductionEnv{
		return fmt.Sprintf("http://localhost:%s/%s", ae.Port, shortUrl)
	} else {
		return fmt.Sprintf("https://foobar.com/%s/%s", ae.Port, shortUrl)
	}
}

func CreateAppEnvironment() AppEnvironment {
	osEnvironmentVariableCheck()
	isProductionEnv := IsProductionEnv()
	logger := log.New(os.Stdout, "[app] ", log.LstdFlags)
	logger.SetFlags(log.Lshortfile)

	if !isProductionEnv{
		logger.Println("STARTING APP :: LOCAL DEV ENVIRONMENT")
	} else {
		logger.Println("STARTING APP :: PRODUCTION ENVIRONMENT")
	}


	dbString := os.Getenv(DBUrlEnvironmentVariable)
	var DB *gorm.DB
	var err error

	logger.Println("API initializing connection to database")
	// repeatedly ping DB until connection is established
	for i := 0; i < 10; i++ {
		DB, err = gorm.Open(postgres.Open(dbString), &gorm.Config{})
		if err == nil {
			break
		}
		logger.Printf("Failed to connect to db on %s , retrying...", dbString)
		time.Sleep(6 * time.Second)
	}

	if err != nil {
		logger.Fatal(err)
	}

	logger.Println("API connected to database")

	port := os.Getenv(PostEnvironmentVariable)

	return AppEnvironment{
		Db: models.AppDB { DB: DB},
		Logger: logger,
		Port: port,
		IsProductionEnv: isProductionEnv,
	}

}

func IsProductionEnv() bool {
	return os.Getenv("APP_ENV") != "local_development"
}

func osEnvironmentVariableCheck() {
	envVars := []string{
		DBUrlEnvironmentVariable,
		PostEnvironmentVariable,
		AllowedOriginsEnvVariable,
	}

	for _, v := range envVars {
		if os.Getenv(v) == "" {
			log.Fatalf("FATAL: %s variable not found", v)
		}
	}
}
