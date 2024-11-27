package pkg

import (
	"log"
	"net"
	"work01/internal/repositories"
	"work01/internal/usecases"
	pb "work01/work01/proto"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

const (
	port = ":50051"
)

func NewGRPCServer(gormDatabase *gorm.DB, redisClient *redis.Client) {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	authRepo := repositories.NewAuthorizationRepository(gormDatabase, redisClient)
	authUsecase := usecases.NewAuthorizationUsecase(authRepo)

	userRepo := repositories.NewUserRepository(gormDatabase, redisClient)
	userUsecase := usecases.NewUserUsecase(userRepo)

	pb.RegisterGreeterServer(s, pb.NewGreeterServer())
	pb.RegisterAuthorizationServer(s, pb.NewAuthorizationServer(authUsecase))
	pb.RegisterUserGrpcServiceServer(s, pb.NewUserGrpcServiceServer(userUsecase))

	log.Printf("Server is listening on port %v", port)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
