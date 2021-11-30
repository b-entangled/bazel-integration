package main

import (
	"flag"
	"log"
	"net"

	"github.com/b-entangled/bazel-integration/pkg/company"
	"github.com/b-entangled/bazel-integration/pkg/design"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Server running :8080")
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	design.RegisterCompanyServiceServer(grpcServer, company.NewCompanyServer())
	grpcServer.Serve(lis)
}
