package main

import (
	"work01/internal/servers"
	"work01/pkg"
)

func main() {

	dbServer := servers.NewDBServer()
	redisClient := pkg.NewRedisClient()

	// dbServer.Migrator().DropTable(&entities.Role{}, &entities.Feature{}, &entities.User{}, &entities.RoleFeature{}, &entities.Authorization{})
	// dbServer.AutoMigrate(&entities.Role{}, &entities.Feature{}, &entities.User{}, &entities.RoleFeature{}, &entities.Authorization{})

	// pkg.NewGRPCServer(dbServer, redisClient)
	pkg.RunFiber(dbServer, redisClient)

}
