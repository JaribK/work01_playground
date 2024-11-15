package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"work01/internal/handlers"
	"work01/internal/repositories"
	"work01/internal/usecases"
	"work01/pkg"

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

	// db.Migrator().DropTable(&entities.Role{}, &entities.Feature{}, &entities.User{}, &entities.Permission{}, &entities.RolePermission{}, &entities.Authorization{})
	//db.AutoMigrate(&entities.Role{}, &entities.Feature{}, &entities.User{}, &entities.Permission{}, &entities.RolePermission{}, &entities.Authorization{})

	app := fiber.New()

	authRepo := repositories.NewAuthorizationRepository(db)
	authUsecase := usecases.NewAuthorizationUsecase(authRepo)
	authHandler := handlers.NewHttpAuthorizationHandler(authUsecase)

	app.Get("/auths", authHandler.GetAllAuthorizationsHandler)

	userRepo := repositories.NewUserRepository(db)
	userUsecase := usecases.NewUserUsecase(userRepo)
	userHandler := handlers.NewHttpUserHandler(userUsecase)

	app.Post("/refresh", authHandler.RefreshToken)
	app.Post("/login", authHandler.LoginHandler)
	app.Use("/logout", pkg.TokenValidationMiddleware)
	app.Post("/logout", authHandler.LogoutHandler)
	app.Use("/users", pkg.TokenValidationMiddleware)
	app.Get("/users/:id", userHandler.GetUserByIdHandler)
	app.Get("/users", userHandler.GetAllUsersHandler)
	app.Post("/users", userHandler.CreateUserHandler)
	app.Put("/users/:id", userHandler.UpdateUserHandler)
	app.Delete("/users/:id", userHandler.DeleteUserHandler)

	roleRepo := repositories.NewRoleRepository(db)
	roleUsecase := usecases.NewRoleUsecase(roleRepo)
	roleHandler := handlers.NewHttpRoleHandler(roleUsecase)

	app.Get("/roles/:id", roleHandler.GetRoleByIdHandler)
	app.Get("/roles", roleHandler.GetAllRolesHandler)
	app.Post("/roles", roleHandler.CreateRoleHandler)
	app.Put("/roles/:id", roleHandler.UpdateRoleHandler)
	app.Delete("/roles/:id", roleHandler.DeleteRoleHandler)

	featureRepo := repositories.NewFeatureRepository(db)
	featureUsecase := usecases.NewFeatureUsecase(featureRepo)
	featureHandler := handlers.NewHttpFeatureHandler(featureUsecase)

	app.Get("/features/:id", featureHandler.GetFeatureByIdHandler)
	app.Get("/features", featureHandler.GetAllFeaturesHandler)
	app.Post("/features", featureHandler.CreateFeatureHandler)
	app.Put("/features/:id", featureHandler.UpdateFeatureHandler)
	app.Delete("/features/:id", featureHandler.DeleteFeatureHandler)

	permissionRepo := repositories.NewPermissionRepository(db)
	permissionUsecase := usecases.NewPermissionUsecase(permissionRepo)
	permissionHandler := handlers.NewHttpPermissionHandler(permissionUsecase)

	app.Get("/permissions/:id", permissionHandler.GetPermissionByIdHandler)
	app.Get("/permissions", permissionHandler.GetAllPermissionsHandler)
	app.Post("/permissions", permissionHandler.CreatePermissionHandler)
	app.Put("/permissions/:id", permissionHandler.UpdatePermissionHandler)
	app.Delete("/permissions/:id", permissionHandler.DeletePermissionHandler)

	rolePermissionRepo := repositories.NewRolePermissionRepository(db)
	rolePermissionUsecase := usecases.NewRolePermissionUsecase(rolePermissionRepo)
	rolePermissionHandler := handlers.NewHttpRolePermissionHandler(rolePermissionUsecase)

	app.Get("/rolePermissions/:id", rolePermissionHandler.GetRolePermissionByIdHandler)
	app.Get("/rolePermissions", rolePermissionHandler.GetAllRolePermissionsHandler)
	app.Post("/rolePermissions", rolePermissionHandler.CreateRolePermissionHandler)
	app.Put("/rolePermissions/:id", rolePermissionHandler.UpdateRolePermissionHandler)
	app.Delete("/rolePermissions/:id", rolePermissionHandler.DeleteRolePermissionHandler)

	app.Listen(":8080")

}
