package main

import (
	"context"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "personService/api/go"
	"personService/internal/config"
	"personService/internal/database"
	"personService/internal/dispatcher"
	personsapp "personService/modules/persons/app"
	personscommands "personService/modules/persons/app/commands"
	personsoutbox "personService/modules/persons/app/outbox"
	personsqueries "personService/modules/persons/app/queries"
	personsrpc "personService/modules/persons/app/rpc"
)

const (
	port = ":50051"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.SetupConfigs,

			dispatcher.New,
			database.NewTransaction,

			personsqueries.New,
			personscommands.New,
			personsrpc.New,
			personsoutbox.NewWhenPersonCreated,
			personsapp.NewRepository,
		),

		fx.Invoke(registerHooks))

	app.Run()
}

func registerHooks(lifecycle fx.Lifecycle, rpcServer *personsrpc.Server) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {

				lis, err := net.Listen("tcp", port)
				if err != nil {
					log.Fatalf("failed to listen: %v", err)
				}

				s := grpc.NewServer()
				pb.RegisterPersonsApiServer(s, rpcServer)
				log.Printf("server listening at %v", lis.Addr())

				go s.Serve(lis)

				return nil
			},
			OnStop: func(context.Context) error {
				return nil
			},
		},
	)
}
