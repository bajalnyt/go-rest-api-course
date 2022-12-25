package main

import (
	"context"
	"fmt"

	"github.com/bajalnyt/go-rest-api-course/internal/comment"
	"github.com/bajalnyt/go-rest-api-course/internal/db"
)

// Run instantiates and starts the app
func Run() error {
	fmt.Println("Starting up application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to database")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("Failed to migrate database")
		return err
	}

	cmtService := comment.NewService(db)
	fmt.Println(cmtService.GetComment(context.Background(), "71c5d074-b6cf-11ec-b909-0242ac120002"))
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
