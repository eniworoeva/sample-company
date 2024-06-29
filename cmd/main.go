package main

import (
	"github.com/eniworoeva/sample-company/cmd/server"
	"github.com/eniworoeva/sample-company/internal/repository"
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
