package server

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/kjasuquo/jobslocation/internal/api"
	"github.com/kjasuquo/jobslocation/internal/repository"
	"log"
	"os"
)

//Run injects all dependencies needed to run the app
func Run(db *gorm.DB, port string) {
	newRepo := repository.NewDB(db)

	Handler := api.NewHTTPHandler(newRepo)
	router := SetupRouter(Handler)

	_ = router.Run(":" + port)
}

//Params is a data model of the data in our environment variable
type Params struct {
	Port       string
	DbUsername string
	DbPassword string
	DbHost     string
	DbName     string
	DbMode     string
}

//InitDBParams gets environment variables needed to run the app
func InitDBParams() Params {
	ginMode := os.Getenv("GIN_MODE")
	log.Println(ginMode)
	if ginMode != "release" {
		errEnv := godotenv.Load()
		if errEnv != nil {
			log.Fatal("Error loading .env file")
		}
	}

	port := os.Getenv("PORT")
	dbUsername := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbMode := os.Getenv("DB_MODE")

	return Params{
		Port:       port,
		DbUsername: dbUsername,
		DbPassword: dbPassword,
		DbHost:     dbHost,
		DbName:     dbName,
		DbMode:     dbMode,
	}
}
