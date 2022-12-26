package main

import (
	"fmt"
	"log"
	"net"

	server "demo-casbin/server"
	utils "demo-casbin/utils"

	pb "demo-casbin/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type (
	Config struct {
		GRPCPort string `env:"GRPC_PORT" default:"50051"`

		DbDriverName    string `env:"DB_DRIVER_NAME" default:"postgres"`
		DbConnectString string `env:"DB_CONN_STRING" default:"user=postgres password=admin host=127.0.0.1 port=5432 sslmode=disable"`

		ModelPath string `env:"MODEL_PATH" default:"models/rbac_model.conf"`

		CacheConnectString string `env:"CACHE_CONN_STRING" default:"127.0.0.1:6379"`
		CachePassword      string `env:"CACHE_PASSWORD" default:""`
		CacheChannel       string `env:"CACHE_CHANNEL" default:"casbin"`
	}
)

var (
	conf Config
)

func main() {
	// Update configure with environment variable
	utils.ReadConfig(&conf)

	// Check port
	log.Println("Listening on", conf.GRPCPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", conf.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create new server
	s := grpc.NewServer()
	pb.RegisterCasbinServer(s, server.NewServer(conf.DbDriverName, conf.DbConnectString, conf.ModelPath, conf.CacheConnectString, conf.CachePassword, conf.CacheChannel))
	reflection.Register(s)
	log.Println("Listening on", conf.GRPCPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
