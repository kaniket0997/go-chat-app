package main

import (
	"github.com/go-chat-appp/main/server/db"
	"github.com/go-chat-appp/main/server/internal/user"
	"github.com/go-chat-appp/main/server/router"
	"log"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatal("Error while connecting to database")
	}
	userRepo := user.NewUserRepository(dbConn.GetDb())
	userService := user.NewService(userRepo, 0)
	userHandler := user.NewUserHandler(userService)
	router.InitRouter(userHandler)
	err = router.Start("0.0.0.0:8080")
	if err != nil {
		log.Fatal("Error while starting server")
	}
}
