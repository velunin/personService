package main

import (
	"context"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	pb "personService/api/go"
	personcommands "personService/app/commands/persons"
	"personService/app/config"
	personqueries "personService/app/queries/persons"
	"personService/app/repositories"
	personrepo "personService/app/repositories/person"
	"personService/app/rpc"
	"strings"
)

const (
	port = ":50051"
)

func main() {

	app := fx.New(
		fx.Provide(
			config.SetupConfigs,
			repositories.NewTransaction,

			rpc.New,
			personrepo.NewPersonRepository,
			personqueries.NewPersonQueryService,
			personcommands.NewPersonCommandService,
		),

		fx.Invoke(registerHooks))

	app.Run()
}

func registerHooks(
	lifecycle fx.Lifecycle, rpcServer *rpc.PersonServer,
) {
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

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO(tamird): point to merged gRPC code rather than a PR.
		// This is a partial recreation of gRPC's internal checks https://github.com/grpc/grpc-go/pull/514/files#diff-95e9a25b738459a2d3030e1e6fa2a718R61
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}
