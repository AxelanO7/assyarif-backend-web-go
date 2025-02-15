package main

import (
	"log"
	"sync"

	"assyarif-backend-web-go/assyarif/delivery"
	"assyarif-backend-web-go/assyarif/repository"
	"assyarif-backend-web-go/assyarif/usecase"
	"assyarif-backend-web-go/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	Init()
	// initEnv()
	listenPort := ":4000"
	// appName := os.Getenv("APP_NAME")

	usrRepo := repository.NewPostgreUser(db.GormClient.DB)
	inRepo := repository.NewPostgreIn(db.GormClient.DB)
	outRepo := repository.NewPostgreOut(db.GormClient.DB)
	employeeRepo := repository.NewPostgreEmployee(db.GormClient.DB)
	outletRepo := repository.NewPostgreOutlet(db.GormClient.DB)
	stockRepo := repository.NewPostgreStock(db.GormClient.DB)
	orderRepo := repository.NewPostgreOrder(db.GormClient.DB)
	rtrRepo := repository.NewPostgreRtr(db.GormClient.DB)
	stockOutletRepo := repository.NewPostgreStockOutlet(db.GormClient.DB)
	opnameRepo := repository.NewPostgreOpname(db.GormClient.DB)

	timeoutContext := fiber.Config{}.ReadTimeout

	userUseCase := usecase.NewUserUseCase(usrRepo, timeoutContext)
	inUseCase := usecase.NewInUseCase(inRepo, timeoutContext)
	outUseCase := usecase.NewOutUseCase(outRepo, timeoutContext)
	employeeUseCase := usecase.NewEmployeeUseCase(employeeRepo, timeoutContext)
	outletUseCase := usecase.NewOutletUseCase(outletRepo, timeoutContext)
	stockUseCase := usecase.NewStockUseCase(stockRepo, inRepo, outRepo, rtrRepo, timeoutContext)
	orderUseCase := usecase.NewOrderUseCase(orderRepo, outRepo, timeoutContext)
	rtrUseCase := usecase.NewRtrUseCase(rtrRepo, timeoutContext)
	stockOutletUseCase := usecase.NewStockOutletUseCase(stockOutletRepo, timeoutContext)
	opnameUseCase := usecase.NewOpnameUseCase(opnameRepo, timeoutContext)

	app := fiber.New(fiber.Config{})
	app.Use(logger.New(logger.Config{
		Format:     "${time} | ${green} ${status} ${white} | ${latency} | ${ip} | ${green} ${method} ${white} | ${path} | ${yellow} ${body} ${reset} | ${magenta} ${resBody} ${reset}\n",
		TimeFormat: "02 January 2006 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))
	app.Use(cors.New())

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {

		//call delivery http here
		delivery.NewUserHandler(app, userUseCase)
		delivery.NewInHandler(app, inUseCase)
		delivery.NewOutHandler(app, outUseCase)
		delivery.NewEmployeeHandler(app, employeeUseCase)
		delivery.NewOutletHandler(app, outletUseCase)
		delivery.NewStockHandler(app, stockUseCase)
		delivery.NewOrderHandler(app, orderUseCase)
		delivery.NewRtrHandler(app, rtrUseCase)
		delivery.NewStockOutletHandler(app, stockOutletUseCase)
		delivery.NewOpnameHandler(app, opnameUseCase)
		log.Fatal(app.Listen(listenPort))
		wg.Done()
	}()
	wg.Wait()

}

func Init() {
	InitEnv()
	InitDB()
}

func InitEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Println(".env file not found")
	}
}

func InitDB() {
	db.NewGormClient()
}
