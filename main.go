package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"work01/internal/handlers"
	"work01/internal/repositories"
	"work01/internal/usecases"

	"github.com/gofiber/fiber/v2"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "myuser"
	password = "mypassword"
	dbname   = "mydatabase"
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	var err error

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
		DryRun: false,
	})
	if err != nil {
		panic("failed connect to database")
	}

	// db.AutoMigrate(&entities.Role{}, &entities.Feature{}, &entities.User{}, &entities.Permission{})

	app := fiber.New()

	roleRepo := repositories.NewRoleRepository(db)
	roleUsecase := usecases.NewRoleService(roleRepo)
	roleHandler := handlers.NewHttpRoleHandler(roleUsecase)

	app.Get("/roles/:id", roleHandler.GetRoleByIdHandler)
	app.Get("/roles", roleHandler.GetAllRoleHandler)
	app.Post("/roles", roleHandler.CreateRoleHandler)
	app.Put("/roles/:id", roleHandler.UpdateRoleHandler)
	app.Delete("/roles/:id", roleHandler.DeleteRoleHandler)

	app.Listen(":8080")

}
