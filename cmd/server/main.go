package main

import (
	"fmt"
	"github.com/Nameer-kp/go-rest-api/internal/comment"
	"github.com/Nameer-kp/go-rest-api/internal/database"
	transportHttp "github.com/Nameer-kp/go-rest-api/internal/transport/http"
)

func Run() error {
	fmt.Println("starting up application")
	db, err := database.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)

	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}
	cmtService := comment.NewService(db)
	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}
	return nil
}
func main() {
	fmt.Println("Go REST API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
