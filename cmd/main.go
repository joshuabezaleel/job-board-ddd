package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/joshuabezaleel/job-board/database"
	"github.com/joshuabezaleel/job-board/pkg/core/user"
	"github.com/joshuabezaleel/job-board/server"
)

func main() {
	port := ":8080"
	connusername := "postgres"
	connpassword := "postgres"
	dbname := "test-db-14-9"
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", connusername, connpassword, dbname)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	var userRepo user.Repository
	userRepo = database.NewRepository(db)

	var userService user.Service
	userService = user.NewService(userRepo)

	srv := server.New(userService)

	err = http.ListenAndServe(port, srv.Router)
	if err != nil {
		log.Fatal(err)
	}
}
