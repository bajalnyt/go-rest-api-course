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
	fmt.Println(cmtService.GetComment(context.Background(), "446aca87-32b0-4af5-a715-7ab4a3b13a56"))

	// cmtService.PostComment(context.Background(), comment.Comment{
	// 	Slug:   "test",
	// 	Body:   "bodyy",
	// 	Author: "authorr",
	// })

	cmtService.UpdateComment(context.Background(), "446aca87-32b0-4af5-a715-7ab4a3b13a56")

	cmtService.DeleteComment(context.Background(), "446aca87-32b0-4af5-a715-7ab4a3b13a56")

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
