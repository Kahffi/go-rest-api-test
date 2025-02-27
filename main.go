package main

import (
	"github.com/aronipurwanto/go-restful-api/app"
	"github.com/aronipurwanto/go-restful-api/controller"
	"github.com/aronipurwanto/go-restful-api/helper"
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"github.com/aronipurwanto/go-restful-api/repository"
	"github.com/aronipurwanto/go-restful-api/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	server := fiber.New()

	// Initialize Database
	db := app.NewDB()

	// Run Auto Migration (Opsional, bisa dihapus jika tidak diperlukan)
	err := db.AutoMigrate(&domain.Category{})
	err = db.AutoMigrate(&domain.Product{})
	err = db.AutoMigrate(&domain.Employee{})
	err = db.AutoMigrate(&domain.Customer{})
	helper.PanicIfError(err)

	// Initialize Validator
	validate := validator.New()

	// Initialize Repository, Service, and Controller
	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository, validate)
	categoryController := controller.NewCategoryController(categoryService)

	employeeRepository := repository.NewEmployeeRepository(db)
	employeeService := service.NewEmployeeService(employeeRepository, validate)
	employeeController := controller.NewEmployeeController(employeeService)

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository, validate)
	productController := controller.NewProductController(productService)

	customerRepository := repository.NewCustomerRepository(db)
	customerService := service.NewCustomerService(customerRepository, validate)
	customerController := controller.NewCustomerController(customerService)

	// Setup Routes
	app.NewRouter(server, categoryController, customerController, productController, employeeController)

	// Start Server
	log.Println("Server running on port 8081")
	err = server.Listen(":8081")
	helper.PanicIfError(err)
}
