package main

import (
	"golang-restful-api/app"
	"golang-restful-api/controller"
	"golang-restful-api/repository"
	"golang-restful-api/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db:=app.NewDB()

	validate:= validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categorySerive := service.NewCategoryService(categoryRepository, db, validate)
	categoryController:= controller.NewCategoryController(categorySerive)

	router:= httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("api/categories/:categoryId", categoryController.FindById)
	router.POST("api/categories", categoryController.Create)
	router.PUT("api/categories/:categoryId", categoryController.Update)
	router.DELETE("api/categories/:categoryId", categoryController.Delete)

}