package main

import (
	"net/http"
	"quizy-api/app"
	"quizy-api/controller"
	"quizy-api/controller/role_controller"
	"quizy-api/controller/user_controller"
	"quizy-api/helper"
	"quizy-api/repository/role_repository"
	"quizy-api/repository/user_repository"
	"quizy-api/service/role_service"
	"quizy-api/service/user_service"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	roleRepository := role_repository.NewRoleRepository()
	roleService := role_service.NewRoleService(db, roleRepository, validate)
	roleController := role_controller.NewServiceController(roleService)

	userRepository := user_repository.NewUserRepository()
	userService := user_service.NewUserService(db, userRepository, validate)
	userController := user_controller.NewUserController(userService)

	router := httprouter.New()

	router.POST("/api/v1/role", roleController.Create)
	router.PUT("/api/v1/role/:roleId", roleController.Update)
	router.DELETE("/api/v1/role/:roleId", roleController.Delete)

	router.POST("/api/v1/user", userController.Create)
	router.PUT("/api/v1/user/:userId", userController.Update)
	router.DELETE("/api/v1/user/:userId", userController.Delete)

	router.PanicHandler = controller.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
