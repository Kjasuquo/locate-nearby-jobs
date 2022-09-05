package helpers

import (
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/kjasuquo/jobslocation/internal/model"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

//LoadCSV reads the CSV file and returns its contents
func LoadCSV(path string) (csvLine [][]string, err error) {
	csvFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return nil, err
	}
	return csvLines, err
}

//CsvToJob converts the contents of the CSV file to our Job model struct
func CsvToJob(csvLines [][]string) ([]model.Jobs, error) {
	var jobs []model.Jobs

	for i := 1; i < len(csvLines); i++ {
		for j := 0; j < len(csvLines[i]); j++ {

			longitude, err := strconv.ParseFloat(csvLines[i][1], 64)
			if err != nil {
				log.Fatalf("error converting longitude string to float64: %v\n", err)
			}

			latitude, err := strconv.ParseFloat(csvLines[i][2], 64)
			if err != nil {
				log.Fatalf("error converting latitude string to float64: %v\n", err)
			}

			job := model.Jobs{
				Title:     csvLines[i][0],
				Longitude: longitude,
				Latitude:  latitude,
			}

			jobs = append(jobs, job)

			break
		}

	}
	return jobs, nil
}

//FindJob helps to check the database if there is any job already in the DB
func FindJob(db *gorm.DB) []model.Jobs {
	var job []model.Jobs
	err := db.First(&job).Error
	if err != nil {
		log.Fatalf("Error finding jobs in db: %v\n", err)
	}
	return job
}

//CreateJob creates the job in the Database
func CreateJob(data []model.Jobs, db *gorm.DB) {
	for _, d := range data {
		db.Create(&d)
	}
}

//CheckAndPopulate checks the db if if it has already been populated, if not, it populates it
func CheckAndPopulate(db *gorm.DB, path string) {
	jobs := FindJob(db)

	//Checking if there is no job before populating the Database
	if len(jobs) == 0 {
		loadCSV, err := LoadCSV(path)
		if err != nil {
			log.Fatalf("error occured while reading CSV: %v\n", err)
		}

		job, err := CsvToJob(loadCSV)
		if err != nil {
			log.Fatalf("error occured while converting csv files to jobs: %v\n", err)
		}

		CreateJob(job, db)
	}
}

//Response is customized to help return all responses need
func Response(c *gin.Context, message string, status int, data interface{}, errs []string) {
	responsedata := gin.H{
		"message":   message,
		"data":      data,
		"errors":    errs,
		"status":    http.StatusText(status),
		"timestamp": time.Now().Format("2006-01-02 15:04:05"),
	}

	c.JSON(status, responsedata)
}
