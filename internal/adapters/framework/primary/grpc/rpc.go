package rpc

import (
	"context"

	"github.com/amillert/hexa-go-ne/internal/adapters/framework/primary/grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (grpca Adapter) GetAddition(ctx context.Context, req *pb.OperationParameters) (*pb.OperationResponse, error) {
	emptyResponse := &pb.OperationResponse{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return emptyResponse, status.Error(codes.InvalidArgument, "missing required values")
	}

	result, err := grpca.api.GetAddition(req.GetA(), req.GetB())
	if err != nil {
		return emptyResponse, status.Error(codes.Internal, "unepected error")
	}

	return &pb.OperationResponse{Value: result}, nil
}

func (grpca Adapter) GetSubtraction(ctx context.Context, req *pb.OperationParameters) (*pb.OperationResponse, error) {
	emptyResponse := &pb.OperationResponse{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return emptyResponse, status.Error(codes.InvalidArgument, "missing required values")
	}

	result, err := grpca.api.GetSubtraction(req.GetA(), req.GetB())
	if err != nil {
		return emptyResponse, status.Error(codes.Internal, "unepected error")
	}

	return &pb.OperationResponse{Value: result}, nil
}

func (grpca Adapter) GetMultiplication(ctx context.Context, req *pb.OperationParameters) (*pb.OperationResponse, error) {
	emptyResponse := &pb.OperationResponse{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return emptyResponse, status.Error(codes.InvalidArgument, "missing required values")
	}

	result, err := grpca.api.GetMultiplication(req.GetA(), req.GetB())
	if err != nil {
		return emptyResponse, status.Error(codes.Internal, "unepected error")
	}

	return &pb.OperationResponse{Value: result}, nil
}

func (grpca Adapter) GetDivision(ctx context.Context, req *pb.OperationParameters) (*pb.OperationResponse, error) {
	emptyResponse := &pb.OperationResponse{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return emptyResponse, status.Error(codes.InvalidArgument, "missing required values")
	}

	result, err := grpca.api.GetDivision(req.GetA(), req.GetB())
	if err != nil {
		return emptyResponse, status.Error(codes.Internal, "unepected error")
	}

	return &pb.OperationResponse{Value: result}, nil
}
