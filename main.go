package main

import (
	"indonesian-heroes/connection"
	"log"
	"os"
)

func main() {
	// get env
	username := os.Getenv("username")
	password := os.Getenv("password")
	host := os.Getenv("host")
	port := os.Getenv("port")
	database := os.Getenv("database")

	env := map[string]string{
		"username": username,
		"password": password,
		"host":     host,
		"port":     port,
		"database": database,
	}

	// create connection
	connection, err := connection.NewConnection(env)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(connection)
}
