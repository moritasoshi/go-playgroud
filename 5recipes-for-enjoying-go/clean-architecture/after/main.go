package main

import (
	"log"

	"after/api"
	"after/config"
	"after/infrastructure/db"
	"after/infrastructure/server"
)

func main() {
	conn, err := db.NewDB()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	defer conn.Close()

	srv := server.NewServer(api.BuildRouter(conn))

	log.Printf("Serving on localhost:%v\n", config.Config.ServerPort)
	log.Fatal(srv.ListenAndServe())
}
