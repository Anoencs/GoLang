package server

import (
	"context"
	pb "demo-casbin/proto"
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
	gormadapter "github.com/casbin/gorm-adapter/v3"

	// rediswatcher "github.com/casbin/redis-watcher/v2"
	rediswatcher "demo-casbin/watcher"

	"github.com/go-redis/redis/v8"
)

// Server is used to implement proto.CasbinServer.
type Server struct {
	pb.UnimplementedCasbinServer

	enforcer *casbin.Enforcer
	watcher  persist.WatcherEx
}

type EnforceRequest struct {
	sub string
	obj string
	act string
}

func NewServer(dbDriverName string, dbConnectString string, modelPath string, cacheConnectString string, cachePassword string, cacheChannel string) *Server {
	s := Server{}

	a, _ := gormadapter.NewAdapter(dbDriverName, dbConnectString)
	e, _ := casbin.NewEnforcer(modelPath, a)

	w, _ := rediswatcher.NewWatcher(cacheConnectString, rediswatcher.WatcherOptions{
		Options: redis.Options{
			Network:  "tcp",
			Password: cachePassword,
		},
		Channel: cacheChannel,
		// Only exists in test, generally be true
		IgnoreSelf: false,
	})
	_ = e.SetWatcher(w)
	_ = w.SetUpdateCallback(func(msg string) {
		log.Println(msg)
		e.LoadPolicy()
	})

	s.enforcer = e
	s.watcher = w

	return &s
}

func (s *Server) CheckPermission(c context.Context, in *pb.AuthorizeRequest) (*pb.BoolReply, error) {
	log.Println(in)

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

func (s *Server) AddPolicy(c context.Context, in *pb.AddPolicyRequest) (*pb.EmptyReply, error) {
	s.enforcer.AddPolicy(in.Sub, in.Obj, in.Act)
	s.enforcer.SavePolicy()
	return &pb.EmptyReply{}, nil
}
