package rpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"personService/api/go"
	personqueries "personService/app/queries/persons"
	"personService/app/repositories/person"
)

type PersonServer struct {
	proto.UnimplementedPersonsApiServer
	personsRepository person.Repository
	queryService      personqueries.PersonQueryService
}

func (s *PersonServer) GetPersons(context.Context, *proto.PersonsRequest) (*proto.PersonsResponse, error) {
	persons, err := s.queryService.GetPersons(context.Background(), personqueries.GetPersonsQuery{})
	if err != nil {
		return nil, status.Error(codes.Internal, "get Persons error")
	}

	result := make([]*proto.Person, len(persons))
	for i, p := range persons {
		person := &proto.Person{
			Id:        p.Id.String(),
			FirstName: p.FirstName,
			LastName:  p.LastName,
		}
		result[i] = person
	}

	return &proto.PersonsResponse{Persons: result}, nil
}

func New(r person.Repository, qs personqueries.PersonQueryService) *PersonServer {
	return &PersonServer{personsRepository: r, queryService: qs}
}
