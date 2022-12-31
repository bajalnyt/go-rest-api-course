package main

import (
	"fmt"

	"github.com/bajalnyt/go-rest-api-course/internal/comment"
	"github.com/bajalnyt/go-rest-api-course/internal/db"
	transportHttp "github.com/bajalnyt/go-rest-api-course/internal/transport/http"
)

// Run instantiates and starts the app
func Run() error {
	fmt.Println("Starting up application")

	mydb, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to database")
		return err
	}

	if err := mydb.MigrateDB(); err != nil {
		fmt.Println("Failed to migrate database")
		return err
	}

	cmtService := comment.NewService(mydb)

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
