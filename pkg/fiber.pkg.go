package pkg

import (
	"work01/internal/handlers"
	"work01/internal/repositories"
	"work01/internal/usecases"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func RunFiber(gormDatabase *gorm.DB, redisClient *redis.Client) {
	app := fiber.New()

	api := app.Group("/api", TokenValidationMiddleware)
	authService := app.Group("/auth", TokenValidationMiddleware)

	authRepo := repositories.NewAuthorizationRepository(gormDatabase, redisClient)
	authUsecase := usecases.NewAuthorizationUsecase(authRepo)
	authHandler := handlers.NewHttpAuthorizationHandler(authUsecase)

	api.Get("/auths", authHandler.GetAllAuthorizationsHandler)

	userRepo := repositories.NewUserRepository(gormDatabase, redisClient)
	userUsecase := usecases.NewUserUsecase(userRepo)
	userHandler := handlers.NewHttpUserHandler(userUsecase)

	//auth-services
	app.Post("/login", authHandler.LoginHandler)
	authService.Post("/refresh", authHandler.RefreshToken)
	authService.Post("/logout", authHandler.LogoutHandler)

	//users
	api.Get("/users/me", userHandler.GetUserByIdHandler)
	api.Get("/users", userHandler.GetAllUsersHandler)
	api.Post("/users", userHandler.CreateUserHandler)
	api.Put("/users/:id", userHandler.UpdateUserHandler)
	api.Delete("/users/:id", userHandler.DeleteUserHandler)

	roleRepo := repositories.NewRoleRepository(gormDatabase, redisClient)
	roleUsecase := usecases.NewRoleUsecase(roleRepo)
	roleHandler := handlers.NewHttpRoleHandler(roleUsecase)

	//roles
	api.Get("/roles/:id", roleHandler.GetRoleByIdHandler)
	api.Get("/roles", roleHandler.GetAllRolesHandler)
	api.Get("/rolesdropdown", roleHandler.GetAllRolesDropdownHandler)
	api.Post("/roles", roleHandler.CreateRoleHandler)
	api.Put("/roles/:id", roleHandler.UpdateRoleHandler)
	api.Delete("/roles/:id", roleHandler.DeleteRoleHandler)

	featureRepo := repositories.NewFeatureRepository(gormDatabase, redisClient)
	featureUsecase := usecases.NewFeatureUsecase(featureRepo)
	featureHandler := handlers.NewHttpFeatureHandler(featureUsecase)

	//features
	api.Get("/features/:id", featureHandler.GetFeatureByIdHandler)
	api.Get("/features", featureHandler.GetAllFeaturesHandler)
	api.Post("/features", featureHandler.CreateFeatureHandler)
	api.Put("/features/:id", featureHandler.UpdateFeatureHandler)
	api.Delete("/features/:id", featureHandler.DeleteFeatureHandler)

	permissionRepo := repositories.NewPermissionRepository(gormDatabase, redisClient)
	permissionUsecase := usecases.NewPermissionUsecase(permissionRepo)
	permissionHandler := handlers.NewHttpPermissionHandler(permissionUsecase)

	//permissions
	api.Get("/permissions/:id", permissionHandler.GetPermissionByIdHandler)
	api.Get("/permissions", permissionHandler.GetAllPermissionsHandler)
	api.Post("/permissions", permissionHandler.CreatePermissionHandler)
	api.Put("/permissions/:id", permissionHandler.UpdatePermissionHandler)
	api.Delete("/permissions/:id", permissionHandler.DeletePermissionHandler)

	rolePermissionRepo := repositories.NewRolePermissionRepository(gormDatabase, redisClient)
	rolePermissionUsecase := usecases.NewRolePermissionUsecase(rolePermissionRepo)
	rolePermissionHandler := handlers.NewHttpRolePermissionHandler(rolePermissionUsecase)

	//rolePermissions
	api.Get("/rolePermissions/:id", rolePermissionHandler.GetRolePermissionByIdHandler)
	api.Get("/rolePermissions", rolePermissionHandler.GetAllRolePermissionsHandler)
	api.Post("/rolePermissions", rolePermissionHandler.CreateRolePermissionHandler)
	api.Put("/rolePermissions/:id", rolePermissionHandler.UpdateRolePermissionHandler)
	api.Delete("/rolePermissions/:id", rolePermissionHandler.DeleteRolePermissionHandler)

	app.Listen(":8080")
}
