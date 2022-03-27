package main

import (
	"os"

	"github.com/amillert/hexa-go-ne/internal/adapters/app/api"
	"github.com/amillert/hexa-go-ne/internal/adapters/core/arithmetic"
	rpc "github.com/amillert/hexa-go-ne/internal/adapters/framework/primary/grpc"
	"github.com/amillert/hexa-go-ne/internal/adapters/framework/secondary/db"
	"github.com/amillert/hexa-go-ne/internal/ports"
)

func main() {
	// ports
	var dbAdapter ports.DbPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	dbDriver := os.Getenv("DB_DRIVER")
	dSourceName := os.Getenv("DS_NAME")
	dbAdapter = db.NewAdapter(dbDriver, dSourceName)

	defer dbAdapter.CloseDbConnection()

	// adapters
	core = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(core, dbAdapter)
	gRPCAdapter = rpc.NewAdapter(appAdapter)
	gRPCAdapter.Run()
}
