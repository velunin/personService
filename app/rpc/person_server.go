package rpc

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"personService/api/go"
	personqueries "personService/app/queries/persons"
	repoperson "personService/app/repositories/person"
)

type PersonServer struct {
	proto.UnimplementedPersonsApiServer
	personsRepository repoperson.Repository
	queryService      personqueries.PersonQueryService
}

func (s *PersonServer) GetPersons(ctx context.Context, req *proto.GetPersonsRequest) (*proto.GetPersonsResponse, error) {
	persons, err := s.queryService.GetPersons(ctx, personqueries.GetPersonsQuery{
		Offset: req.Offset,
		Limit:  req.Limit,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
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

	return &proto.GetPersonsResponse{Persons: result}, nil
}

func (s *PersonServer) GetPerson(ctx context.Context, req *proto.GetPersonRequest) (*proto.GetPersonResponse, error) {
	personId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "id must be UUID")
	}

	person, err := s.queryService.GetPerson(ctx, personqueries.GetPersonQuery{Id: personId})
	if err != nil {
		if errors.Is(err, repoperson.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "person not found")
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.GetPersonResponse{Person: &proto.Person{
		Id:        person.Id.String(),
		FirstName: person.FirstName,
		LastName:  person.LastName,
		IsBlocked: person.IsBlocked,
	}}, nil
}

func New(r repoperson.Repository, qs personqueries.PersonQueryService) *PersonServer {
	return &PersonServer{personsRepository: r, queryService: qs}
}
