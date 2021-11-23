package main

import (
	"github.com/sirupsen/logrus"

	"github.com/pxwxnvermx/todo-rest/routes"
	"github.com/pxwxnvermx/todo-rest/storage"
	"github.com/pxwxnvermx/todo-rest/utils"
)

func main() {
	var logger = logrus.New()
	var err error

	db, err := utils.InitDB()
	if err != nil {
		logger.Error(err)
	}

	s := storage.NewStorage(db)
	r := routes.NewRouter(logger, s)

	if err = r.Run(":3000"); err != nil {
		logger.Error(err)
	}
}
