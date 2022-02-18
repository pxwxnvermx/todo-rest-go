package main

import (
	db "github.com/pxwxnvermx/todo-rest-go/db/sqlc"
	"github.com/pxwxnvermx/todo-rest-go/routes"
	"github.com/pxwxnvermx/todo-rest-go/utils"
	logger "github.com/sirupsen/logrus"
)

func main() {
	var err error
	config := utils.LoadConfig()

	database, err := utils.InitDB(config.Database[0])
	if err != nil {
		logger.Error(err)
	}

	s := db.NewStore(database)
	r := routes.NewRouter(s)

	if err = r.Run(":5000"); err != nil {
		logger.Error(err)
	}

}
