package main

import (
	"fmt"
	"log"
	"net"

	"github.com/SaidovZohid/medium_post_service/config"
	pb "github.com/SaidovZohid/medium_post_service/genproto/post_service"
	"github.com/SaidovZohid/medium_post_service/service"
	"github.com/SaidovZohid/medium_post_service/storage"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load(".")

	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)

	psqlConn, err := sqlx.Connect("postgres", psqlUrl)

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	strg := storage.NewStoragePg(psqlConn)

	postService := service.NewPostService(&strg)
	categoryService := service.NewCategoryService(&strg)

	listen, err := net.Listen("tcp", cfg.GrpcPort)

	s := grpc.NewServer()
	pb.RegisterPostServiceServer(s, postService)
	pb.RegisterCategoryServiceServer(s, categoryService)
	reflection.Register(s)

	log.Println("gRPC server started port in: ", cfg.GrpcPort)
	if s.Serve(listen); err != nil {
		log.Fatalf("Error while listening: %v", err)
	}
}
