package main

import (
	"work01/internal/servers"
	"work01/pkg"
	"work01/pkg/minio"
)

func main() {

	dbServer := servers.NewDBServer()
	redisClient := pkg.NewRedisClient()
	minio.NewMinioClient()

	// app := fiber.New()

	// api := app.Group("/api/v2", pkg.TokenValidationMiddleware)
	// authService := app.Group("/auth", pkg.TokenValidationMiddleware)

	// authRepo := repositories.NewAuthorizationRepository(dbServer, redisClient)
	// authUsecase := usecases.NewAuthorizationUsecase(authRepo)
	// authHandler := handlers.NewHttpAuthorizationHandler(authUsecase)

	// api.Get("/auths", authHandler.GetAllAuthorizationsHandler)

	// userRepo := repositories.NewUserRepository(dbServer, redisClient)
	// userUsecase := usecases.NewUserUsecase(userRepo)
	// userHandler := handlers.NewHttpUserHandler(userUsecase)

	// //auth-services
	// app.Post("/login", authHandler.LoginHandler)
	// app.Post("/token/refresh", authHandler.RefreshToken)
	// authService.Post("/logout", authHandler.LogoutHandler)

	// //users
	// // app.Post("/createusersnoauth", userHandler.CreateUserHandler)
	// api.Get("/users_default", userHandler.GetAllUsersNoPageHandler)
	// api.Get("/users/me", userHandler.GetUserByIdHandler)
	// api.Get("/users/:id", userHandler.GetUserProfileByIdHandler)
	// api.Get("/users", userHandler.GetAllUsersWithPageHandler)
	// api.Post("/users", userHandler.CreateUserHandler)
	// api.Put("/users/changepassword/:id", userHandler.ChangePsswordHandler)
	// api.Put("/users/:id", userHandler.UpdateUserHandler)
	// api.Delete("/users/:id", userHandler.DeleteUserHandler)

	// roleRepo := repositories.NewRoleRepository(dbServer, redisClient)
	// roleUsecase := usecases.NewRoleUsecase(roleRepo)
	// roleHandler := handlers.NewHttpRoleHandler(roleUsecase)

	// //roles
	// api.Get("/roles_default", roleHandler.GetAllRolesDefaultHandler)
	// api.Get("/roles/:id", roleHandler.GetRoleByIdHandler)
	// api.Get("/roles", roleHandler.GetAllRolesModifyHandler)
	// api.Get("/roles_dropdown", roleHandler.GetAllRolesDropdownHandler)
	// api.Post("/roles", roleHandler.CreateRoleHandler)
	// api.Put("/roles/:id", roleHandler.UpdateRoleHandler)
	// api.Delete("/roles/:id", roleHandler.DeleteRoleHandler)

	// featureRepo := repositories.NewFeatureRepository(dbServer, redisClient)
	// featureUsecase := usecases.NewFeatureUsecase(featureRepo)
	// featureHandler := handlers.NewHttpFeatureHandler(featureUsecase)

	// //features
	// api.Get("/features_dropdown", featureHandler.GetRefFeatureHandler)
	// api.Get("/features_default", featureHandler.GetAllFeaturesDefaultHandler)
	// api.Get("/features/:id", featureHandler.GetFeatureByIdHandler)
	// api.Get("/features", featureHandler.GetAllFeaturePermissionsHandler)
	// api.Post("/features", featureHandler.CreateFeatureHandler)
	// api.Put("/features/:id", featureHandler.UpdateFeatureHandler)
	// api.Delete("/features/:id", featureHandler.DeleteFeatureHandler)

	// roleFeatureRepo := repositories.NewRoleFeatureRepository(dbServer, redisClient)
	// roleFeatureUsecase := usecases.NewRoleFeatureUsecase(roleFeatureRepo)
	// roleFeatureHandler := handlers.NewHttpRoleFeatureHandler(roleFeatureUsecase)

	// //roleFeatures
	// api.Get("/role_features/:id", roleFeatureHandler.GetRoleFeatureByIdHandler)
	// api.Get("/role_features", roleFeatureHandler.GetAllRoleFeaturesHandler)
	// api.Post("/role_features", roleFeatureHandler.CreateRoleFeatureHandler)
	// api.Put("/role_features/:id", roleFeatureHandler.UpdateRoleFeatureHandler)
	// api.Delete("/role_features/:id", roleFeatureHandler.DeleteRoleFeatureHandler)

	// app.Listen(":8080")

	// dbServer.Migrator().DropTable(&entities.Role{}, &entities.Feature{}, &entities.User{}, &entities.RoleFeature{}, &entities.Authorization{})
	// dbServer.AutoMigrate(&entities.Role{}, &entities.Feature{}, &entities.User{}, &entities.RoleFeature{}, &entities.Authorization{})

	pkg.NewGRPCServer(dbServer, redisClient)
}
