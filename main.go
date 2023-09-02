package main

import (
	"net/http"
	"quizy-api/app"
	"quizy-api/controller"
	"quizy-api/exception"
	"quizy-api/helper"
	repository "quizy-api/repository/role_repository"
	service "quizy-api/service/role_service"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	roleRepository := repository.NewRoleRepository()
	roleService := service.NewRoleService(db, roleRepository, validate)
	roleController := controller.NewServiceController(roleService)

	router := httprouter.New()

	router.POST("/api/v1/role", roleController.Create)
	router.PUT("/api/v1/role/:roleId", roleController.Update)
	router.DELETE("/api/v1/role/:roleId", roleController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
