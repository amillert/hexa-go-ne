package rpc

import (
	"log"
	"net"

	"github.com/amillert/hexa-go-ne/internal/adapters/framework/primary/grpc/pb"
	"github.com/amillert/hexa-go-ne/internal/ports"
	"google.golang.org/grpc"
)

type Adapter struct{ api ports.APIPort }

func NewAdapter(api ports.APIPort) *Adapter { return &Adapter{api: api} }

func (grpca Adapter) Run() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("listening on port 9000 failed: %v", err)
	}

	// TODO: remove?
	arithmeticServiceServer := grpca
	grpcServer := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(grpcServer, arithmeticServiceServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("serving gRPC server on port 9000 failed: %v", err)
	}
}
