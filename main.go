package main

import (
	"work01/internal/handlers"
	"work01/internal/repositories"
	"work01/internal/servers"
	"work01/internal/usecases"
	"work01/pkg"

	"github.com/gofiber/fiber/v2"
)

func main() {

	dbServer := servers.NewDBServer()
	redisClient := pkg.NewRedisClient()

	// dbServer.Migrator().DropTable(&entities.Role{}, &entities.Feature{}, &entities.User{}, &entities.Permission{}, &entities.RolePermission{}, &entities.Authorization{})
	//dbServer.AutoMigrate(&entities.Role{}, &entities.Feature{}, &entities.User{}, &entities.Permission{}, &entities.RolePermission{}, &entities.Authorization{})

	app := fiber.New()

	authRepo := repositories.NewAuthorizationRepository(dbServer, redisClient)
	authUsecase := usecases.NewAuthorizationUsecase(authRepo)
	authHandler := handlers.NewHttpAuthorizationHandler(authUsecase)

	app.Get("/auths", authHandler.GetAllAuthorizationsHandler)

	userRepo := repositories.NewUserRepository(dbServer, redisClient)
	userUsecase := usecases.NewUserUsecase(userRepo)
	userHandler := handlers.NewHttpUserHandler(userUsecase)

	app.Post("/auth/refresh", authHandler.RefreshToken)
	app.Post("/auth/login", authHandler.LoginHandler)
	app.Post("/auth/logout", authHandler.LogoutHandler)
	// app.Use("/users", pkg.TokenValidationMiddleware)
	app.Get("/users/:id", userHandler.GetUserByIdHandler)
	app.Get("/users", userHandler.GetAllUsersHandler)
	app.Post("/users", userHandler.CreateUserHandler)
	app.Put("/users/:id", userHandler.UpdateUserHandler)
	app.Delete("/users/:id", userHandler.DeleteUserHandler)

	roleRepo := repositories.NewRoleRepository(dbServer)
	roleUsecase := usecases.NewRoleUsecase(roleRepo)
	roleHandler := handlers.NewHttpRoleHandler(roleUsecase)

	app.Use("/roles", pkg.TokenValidationMiddleware)
	app.Get("/roles/:id", roleHandler.GetRoleByIdHandler)
	app.Get("/roles", roleHandler.GetAllRolesHandler)
	app.Get("/rolesdropdown", roleHandler.GetAllRolesDropdownHandler)
	app.Post("/roles", roleHandler.CreateRoleHandler)
	app.Put("/roles/:id", roleHandler.UpdateRoleHandler)
	app.Delete("/roles/:id", roleHandler.DeleteRoleHandler)

	featureRepo := repositories.NewFeatureRepository(dbServer)
	featureUsecase := usecases.NewFeatureUsecase(featureRepo)
	featureHandler := handlers.NewHttpFeatureHandler(featureUsecase)

	app.Get("/features/:id", featureHandler.GetFeatureByIdHandler)
	app.Get("/features", featureHandler.GetAllFeaturesHandler)
	app.Post("/features", featureHandler.CreateFeatureHandler)
	app.Put("/features/:id", featureHandler.UpdateFeatureHandler)
	app.Delete("/features/:id", featureHandler.DeleteFeatureHandler)

	permissionRepo := repositories.NewPermissionRepository(dbServer)
	permissionUsecase := usecases.NewPermissionUsecase(permissionRepo)
	permissionHandler := handlers.NewHttpPermissionHandler(permissionUsecase)

	app.Get("/permissions/:id", permissionHandler.GetPermissionByIdHandler)
	app.Get("/permissions", permissionHandler.GetAllPermissionsHandler)
	app.Post("/permissions", permissionHandler.CreatePermissionHandler)
	app.Put("/permissions/:id", permissionHandler.UpdatePermissionHandler)
	app.Delete("/permissions/:id", permissionHandler.DeletePermissionHandler)

	rolePermissionRepo := repositories.NewRolePermissionRepository(dbServer)
	rolePermissionUsecase := usecases.NewRolePermissionUsecase(rolePermissionRepo)
	rolePermissionHandler := handlers.NewHttpRolePermissionHandler(rolePermissionUsecase)

	app.Get("/rolePermissions/:id", rolePermissionHandler.GetRolePermissionByIdHandler)
	app.Get("/rolePermissions", rolePermissionHandler.GetAllRolePermissionsHandler)
	app.Post("/rolePermissions", rolePermissionHandler.CreateRolePermissionHandler)
	app.Put("/rolePermissions/:id", rolePermissionHandler.UpdateRolePermissionHandler)
	app.Delete("/rolePermissions/:id", rolePermissionHandler.DeleteRolePermissionHandler)

	app.Listen(":8080")

}
