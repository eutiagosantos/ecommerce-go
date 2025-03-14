package main

import (
	"api-golang/infra/db"
	"api-golang/infra/repositories"
	"api-golang/internal/controller"
	"api-golang/internal/core/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db.NewConnection()
	if err != nil {
		panic(err)
	}
	defer db.Statement.ReflectValue.Close()

	userRepository := repositories.NewUserRepository(db)
	userUseCases := usecases.NewUserUseCases(userRepository)

	userController := controller.NewUserController(userUseCases)

	router := gin.Default()
	router.POST("/users", userController.Create)
	router.POST("/login", userController.Login)
	router.Run(":3333")
}
