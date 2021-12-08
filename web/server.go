package web

import (
	pb "github.com/ISS-Dating/service-analyzer/connections/grpc/go/connections"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"log"
	"net"
)

type server struct{}

func (s *server) Match(ctx context.Context, r *pb.MatchRequest) (*pb.MatchResponse, error) {
	var userA []int32
	var userB []int32

	for i, user := range r.Users {
		if i%2 == 0 {
			userA = append(userA, user)
		} else {
			userB = append(userB, user)
		}
	}

	return &pb.MatchResponse{
		UserA: userA,
		UserB: userB,
	}, nil
}

func Start() {
	ln, err := net.Listen("tcp", ":32001")
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	s := grpc.NewServer()
	pb.RegisterAnalyzerServer(s, &server{})
	s.Serve(ln)
}
