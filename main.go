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
	personsoutbox "personService/internal/outboxes/persons"
	personsrepo "personService/modules/persons/app"
	personscommands "personService/modules/persons/app/commands"
	personsqueries "personService/modules/persons/app/queries"
	"personService/modules/persons/app/rpc"
)

const (
	port = ":50051"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.SetupConfigs,

			rpc.New,
			dispatcher.New,
			database.NewTransaction,
			personsoutbox.NewCreateOutboxMessageHandler,

			personsrepo.NewPersonRepository,
			personsqueries.NewPersonQueryService,
			personscommands.NewPersonCommandService,
		),

		fx.Invoke(registerHooks))

	app.Run()
}

func registerHooks(lifecycle fx.Lifecycle, rpcServer *rpc.PersonServer) {
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
