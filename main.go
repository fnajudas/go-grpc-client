package main

import (
	"base-project/config"
	"base-project/handlers"
	student "base-project/proto/students"
	"base-project/routes"
	"base-project/usecase"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		return
	}

	// grpcConn, err := grpc.Dial("localhost:9990", grpc.WithTransportCredentials(insecure.NewCredentials()))
	grpcConn, err := grpc.Dial(
		"localhost:9991",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(1024*1024*64), // Terima pesan hingga 64MB
			grpc.MaxCallSendMsgSize(1024*1024*64), // Kirim pesan hingga 64MB
		),
	)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	routes := setupRoutes(grpcConn)
	routes.Run(cfg.AppPort)
}

func setupRoutes(grpc *grpc.ClientConn) *routes.Routes {
	studentInt := student.NewStudentServiceClient(grpc)
	studentSvc := usecase.NewStudentSvc(studentInt)
	studentHandler := handlers.NewHandler(studentSvc)

	return &routes.Routes{
		Student: studentHandler,
	}
}
