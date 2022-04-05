package main

import (
	// application
	"hex/internal/adapters/app/api"
	"hex/internal/adapters/core/arithmetic"

	"hex/internal/ports"
	"log"
	"os"

	// adapters
	gRPC "hex/internal/adapters/framework/left/grpc"
	"hex/internal/adapters/framework/right/db"
)

func main() {
	var err error

	// ports
	var dbAdapter ports.DbPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	dbDriver := os.Getenv("DB_DRIVER")
	dbSourceName := os.Getenv("DB_NAME")

	dbAdapter, err = db.NewAdapter(dbDriver, dbSourceName)
	if err != nil {
		log.Fatalf("failed to initiate db connection: %v", err)
	}
	defer dbAdapter.CloseDbConnection()

	core = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(dbAdapter, core)

	gRPCAdapter = gRPC.NewAdapter(appAdapter)
	gRPCAdapter.Run()
}
