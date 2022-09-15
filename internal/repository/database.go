package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kjasuquo/jobslocation/internal/helpers"
	"github.com/kjasuquo/jobslocation/internal/model"
	"log"
)

type Postgres struct {
	DB *gorm.DB
}

//NewDB create/returns a new instance of our Database
func NewDB(DB *gorm.DB) Repository {

	return &Postgres{
		DB: DB,
	}
}

//Initialize opens the database, creates jobs table if not created and populate it if its empty and returns a DB
func Initialize(dbURI string) (*gorm.DB, error) {

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatalf("We can't open a DATABASE: %v\n", err)

	}

	conn.AutoMigrate(&model.Jobs{})

	//Checks and populates table if empty with the given csv file
	helpers.CheckAndPopulate(conn, "location_data_2000.csv")

	return conn, nil
}

//SearchJobsByLocation queries the database using the user's longitude and latitude to get jobs within 5km radius
func (db *Postgres) SearchJobsByLocation(title string, long, lat float64) ([]model.Jobs, error) {
	var job []model.Jobs

	if title != "" {
		statement := `SELECT * FROM jobs WHERE title LIKE ?
	AND ST_DWithin(ST_MakePoint(longitude,latitude)::geography,ST_MakePoint(?, ?)::geography,?)
	LIMIT 2000
	`
		err := db.DB.Raw(statement, "%"+title+"%", long, lat, 5000).Scan(&job).Error
		if err != nil {
			log.Printf("cannot find job: %v\n", err)
			return nil, err
		}

	} else {
		statement := `SELECT * FROM jobs
	WHERE ST_DWithin(ST_MakePoint(longitude,latitude)::geography,ST_MakePoint(?, ?)::geography,?)
	LIMIT 2000
	`
		err := db.DB.Raw(statement, long, lat, 5000).Scan(&job).Error
		if err != nil {
			log.Printf("cannot find job: %v\n", err)
			return nil, err
		}

	}

	return job, nil
}

//SearchJobsByTitle returns all the jobs in the database that matches the searched keyword and returns everything if keyword is left empty
func (db *Postgres) SearchJobsByTitle(title string) ([]model.Jobs, error) {
	var job []model.Jobs

	if title != "" {

		statement := `SELECT * FROM jobs WHERE title LIKE ?`
		err := db.DB.Raw(statement, "%"+title+"%").Scan(&job).Error
		if err != nil {
			log.Printf("cannot find job: %v\n", err)
			return nil, err
		}

	} else {
		statement := `SELECT * FROM jobs`
		err := db.DB.Raw(statement).Scan(&job).Error

		if err != nil {
			log.Printf("cannot find job: %v\n", err)
			return nil, err
		}
	}

	return job, nil
}
