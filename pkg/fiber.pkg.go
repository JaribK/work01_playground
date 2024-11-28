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
	app.Post("/token/refresh", authHandler.RefreshToken)
	authService.Post("/logout", authHandler.LogoutHandler)

	//users
	app.Post("/createusersnoauth", userHandler.CreateUserHandler)
	api.Get("/users/default", userHandler.GetAllUsersNoPageHandler)
	api.Get("/users/me", userHandler.GetUserByIdHandler)
	api.Get("/users", userHandler.GetAllUsersWithPageHandler)
	api.Post("/users", userHandler.CreateUserHandler)
	api.Put("/users/:id", userHandler.UpdateUserHandler)
	api.Delete("/users/:id", userHandler.DeleteUserHandler)

	roleRepo := repositories.NewRoleRepository(gormDatabase, redisClient)
	roleUsecase := usecases.NewRoleUsecase(roleRepo)
	roleHandler := handlers.NewHttpRoleHandler(roleUsecase)

	//roles
	api.Get("/roles/default", roleHandler.GetAllRolesDefaultHandler)
	api.Get("/roles/:id", roleHandler.GetRoleByIdHandler)
	api.Get("/roles", roleHandler.GetAllRolesModifyHandler)
	api.Get("/rolesdropdown", roleHandler.GetAllRolesDropdownHandler)
	api.Post("/roles", roleHandler.CreateRoleHandler)
	api.Put("/roles/:id", roleHandler.UpdateRoleHandler)
	api.Delete("/roles/:id", roleHandler.DeleteRoleHandler)

	featureRepo := repositories.NewFeatureRepository(gormDatabase, redisClient)
	featureUsecase := usecases.NewFeatureUsecase(featureRepo)
	featureHandler := handlers.NewHttpFeatureHandler(featureUsecase)

	//features
	api.Get("/features/default", featureHandler.GetAllFeaturesDefaultHandler)
	api.Get("/features/:id", featureHandler.GetFeatureByIdHandler)
	api.Get("/features", featureHandler.GetAllFeaturePermissionsHandler)
	api.Post("/features", featureHandler.CreateFeatureHandler)
	api.Put("/features/:id", featureHandler.UpdateFeatureHandler)
	api.Delete("/features/:id", featureHandler.DeleteFeatureHandler)

	roleFeatureRepo := repositories.NewRoleFeatureRepository(gormDatabase, redisClient)
	roleFeatureUsecase := usecases.NewRoleFeatureUsecase(roleFeatureRepo)
	roleFeatureHandler := handlers.NewHttpRoleFeatureHandler(roleFeatureUsecase)

	//roleFeatures
	api.Get("/roleFeatures/:id", roleFeatureHandler.GetRoleFeatureByIdHandler)
	api.Get("/roleFeatures", roleFeatureHandler.GetAllRoleFeaturesHandler)
	api.Post("/roleFeatures", roleFeatureHandler.CreateRoleFeatureHandler)
	api.Put("/roleFeatures/:id", roleFeatureHandler.UpdateRoleFeatureHandler)
	api.Delete("/roleFeatures/:id", roleFeatureHandler.DeleteRoleFeatureHandler)

	app.Listen(":8080")
}
