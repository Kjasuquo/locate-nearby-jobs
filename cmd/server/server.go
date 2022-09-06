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
	Port  string
	DbUrl string
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
	dbURL := os.Getenv("DATABASE_URL")

	return Params{
		Port:  port,
		DbUrl: dbURL,
	}
}
