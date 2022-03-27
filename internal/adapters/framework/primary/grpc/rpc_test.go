package rpc

import (
	"context"
	"log"
	"net"
	"os"
	"testing"

	// application
	"github.com/amillert/hexa-go-ne/internal/adapters/app/api"
	"github.com/amillert/hexa-go-ne/internal/adapters/core/arithmetic"

	// adapters
	"github.com/amillert/hexa-go-ne/internal/adapters/framework/primary/grpc/pb"
	"github.com/amillert/hexa-go-ne/internal/adapters/framework/secondary/db"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

var listener *bufconn.Listener

func init() {
	const bufSize = 1024 * 1024

	listener = bufconn.Listen(bufSize)
	grpcServer := grpc.NewServer()

	dbDriver := os.Getenv("DB_DRIVER")
	dSourceName := os.Getenv("DS_NAME")

	dbAdapter := db.NewAdapter(dbDriver, dSourceName)

	// core
	core := arithmetic.NewAdapter()

	applicationAPI := api.NewAdapter(core, dbAdapter)

	gRPCAdapter := NewAdapter(applicationAPI)

	pb.RegisterArithmeticServiceServer(grpcServer, gRPCAdapter)
	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("test server start error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) { return listener.Dial() }

func getGRPCConnection(ctx context.Context, t *testing.T) *grpc.ClientConn {
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial bufnet: %v", err)
	}

	return conn
}

func TestGetAddition(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)

	params := &pb.OperationParameters{A: 1, B: 1}

	answer, err := client.GetAddition(ctx, params)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer.GetValue(), int32(2))
}

func TestGetSubtraction(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)

	params := &pb.OperationParameters{A: 1, B: 1}

	answer, err := client.GetSubtraction(ctx, params)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer.GetValue(), int32(0))
}

func TestGetMultiplication(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)

	params := &pb.OperationParameters{A: 1, B: 1}

	answer, err := client.GetMultiplication(ctx, params)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer.GetValue(), int32(1))
}

func TestGetDivision(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)

	params := &pb.OperationParameters{A: 1, B: 1}

	answer, err := client.GetDivision(ctx, params)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer.GetValue(), int32(1))
}
