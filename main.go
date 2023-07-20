package main

import (
	"gofiber-ewallet/internal/api"
	"gofiber-ewallet/internal/component"
	"gofiber-ewallet/internal/config"
	"gofiber-ewallet/internal/middleware"
	"gofiber-ewallet/internal/repository"
	"gofiber-ewallet/internal/service"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	config := config.Get()
	db := component.ConnectDB(config)
	cacheConnection := component.GetCacheConnnection()
	userRepository := repository.NewUser(db)
	userService := service.NewUser(userRepository, cacheConnection)
	authMid := middleware.Authenticate(userService)

	app := fiber.New()
	api.NewAuth(app, userService, authMid)

	logger := logrus.New()

	file, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logger.Printf("Failed to Create Log File %s", err.Error())
	}
	logger.SetOutput(file)

	logger.Println("Server Running on Port 8080")

	_ = app.Listen(config.Server.Host + ":" + config.Server.Port)
}
