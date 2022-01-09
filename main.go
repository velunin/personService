package main

import (
	"context"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "personService/api/go"
	personcommands "personService/app/commands/persons"
	"personService/app/config"
	"personService/app/database"
	"personService/app/dispatcher"
	"personService/app/eventhandlers/person/whenPersonCreated"
	personqueries "personService/app/queries/persons"
	personrepo "personService/app/repositories/person"
	"personService/app/rpc"
)

const (
	port = ":50051"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.SetupConfigs,

			rpc.New,
			database.NewTransaction,
			personrepo.NewPersonRepository,
			personqueries.NewPersonQueryService,
			personcommands.NewPersonCommandService,

			dispatcher.New,
			whenPersonCreated.NewCreateOutboxMessageHandler,
			whenPersonCreated.NewDoSomeActionsHandler,
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
