package main

import (
	"github.com/kjasuquo/jobslocation/cmd/server"
	"github.com/kjasuquo/jobslocation/internal/repository"
)

func main() {
	//Gets the environment variables
	env := server.InitDBParams()

	//Initializes the database
	db, err := repository.Initialize(env.DbUrl)
	if err != nil {
		return
	}

	//Runs the app
	server.Run(db, env.Port)

}
