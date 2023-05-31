package main

import (
	"github.com/go-chat-app/main/server/db"
	"github.com/go-chat-app/main/server/internal/user"
	"github.com/go-chat-app/main/server/router"
	"log"
	"time"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatal("Error while connecting to database")
	}
	userRepo := user.NewUserRepository(dbConn.GetDb())
	userService := user.NewService(userRepo, time.Duration(20)*time.Second)
	userHandler := user.NewUserHandler(userService)
	router.InitRouter(userHandler)
	err = router.Start("0.0.0.0:8080")
	if err != nil {
		log.Fatal("Error while starting server")
	}
}
