package server

import (
	"context"
	pb "demo-casbin/proto"
	"log"

	"github.com/casbin/casbin/v2"

	gormadapter "github.com/casbin/gorm-adapter/v3"
)

type Server struct {
	pb.UnimplementedCasbinServer

	enforcer *casbin.Enforcer
}

// Init Casbin server
func NewServer(dbDriverName string, dbConnectString string, modelPath string) *Server {
	s := Server{}

	// Init adapter to database
	a, _ := gormadapter.NewAdapter(dbDriverName, dbConnectString)
	// Init enforcer using model and adapter (to database)
	e, _ := casbin.NewEnforcer(modelPath, a)

	s.enforcer = e

	return &s
}

// Implement CheckPermission API
func (s *Server) CheckPermission(c context.Context, in *pb.AuthorizeRequest) (*pb.BoolReply, error) {
	log.Println(in)

	// Call Enforce() method to check permission
	r, err := s.enforcer.Enforce(in.Sub, in.Obj, in.Act)
	if err != nil {
		return nil, err
	} else {
		if r {
			return &pb.BoolReply{Res: true}, nil
		} else {
			return &pb.BoolReply{Res: false}, nil
		}
	}
}
