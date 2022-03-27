package ports

import (
	"context"

	"github.com/amillert/hexa-go-ne/internal/adapters/framework/primary/grpc/pb"
)

type GRPCPort interface {
	Run()
	GetAddition(context.Context, *pb.OperationParameters) (*pb.OperationResponse, error)
	GetSubtraction(context.Context, *pb.OperationParameters) (*pb.OperationResponse, error)
	GetMultiplication(context.Context, *pb.OperationParameters) (*pb.OperationResponse, error)
	GetDivision(context.Context, *pb.OperationParameters) (*pb.OperationResponse, error)
}
