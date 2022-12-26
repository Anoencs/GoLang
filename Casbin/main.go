package main

import (
	pb "casbin-golang/proto"
	"casbin-golang/utils"
	"fmt"
	"log"
	"net"

	"github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/server"
)

func main() {
	// Update configure with environment variable
	conf, err := utils.ReadConfig(".")

	// Check port
	log.Println("Listening on", conf.GRPCPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", conf.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create new server
	s := grpc.NewServer()
	pb.RegisterCasbinServer(s, server.NewServer(conf.DbDriverName, conf.DbConnectString, conf.ModelPath))
	reflection.Register(s)
	log.Println("Listening on", conf.GRPCPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
