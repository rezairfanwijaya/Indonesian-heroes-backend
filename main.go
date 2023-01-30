package main

import (
	"indonesian-heroes/connection"
	"indonesian-heroes/handler"
	"indonesian-heroes/hero"
	"log"
	"os"

	"github.com/gin-gonic/gin"
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

	repo := hero.NewRepository(connection)
	service := hero.NewService(repo)
	handler := handler.NewHeroHandler(service)

	// init server
	r := gin.Default()

	r.GET("/heros", handler.GetAllHero)

	r.Run(":7575")
}
