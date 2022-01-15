package rpc

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"personService/api/go"
	"personService/modules/persons/app"
	"personService/modules/persons/app/commands"
	"personService/modules/persons/app/queries"
)

type PersonServer struct {
	proto.UnimplementedPersonsApiServer
	personsRepository app.Repository
	queryService      queries.PersonQueryService
	commandService    commands.PersonCommandService
}

func (s *PersonServer) GetPersons(ctx context.Context, req *proto.GetPersonsRequest) (*proto.GetPersonsResponse, error) {
	persons, err := s.queryService.GetPersons(ctx, queries.GetPersonsQuery{
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
			IsBlocked: p.IsBlocked,
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

	person, err := s.queryService.GetPerson(ctx, queries.GetPersonQuery{Id: personId})
	if err != nil {
		if errors.Is(err, app.ErrNotFound) {
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

func (s *PersonServer) CreatePerson(ctx context.Context,
	req *proto.CreatePersonRequest) (*proto.CreatePersonResponse, error) {
	//
	id, err := s.commandService.CreatePerson(ctx, commands.CreatePersonCommand{
		FirstName: req.FirstName,
		LastName:  req.LastName,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.CreatePersonResponse{Id: id.String()}, nil
}

func (s *PersonServer) RenamePerson(ctx context.Context,
	req *proto.RenamePersonRequest) (*proto.RenamePersonResponse, error) {
	//
	personId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "id must be UUID")
	}

	err = s.commandService.RenamePerson(ctx, commands.RenamePersonCommand{
		Id:        personId,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.RenamePersonResponse{}, nil
}

func (s *PersonServer) BlockPerson(ctx context.Context,
	req *proto.BlockPersonRequest) (*proto.BlockPersonResponse, error) {
	//
	personId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "id must be UUID")
	}

	err = s.commandService.BlockPerson(ctx, commands.BlockPersonCommand{
		Id: personId,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.BlockPersonResponse{}, nil
}

func (s *PersonServer) UnblockPerson(ctx context.Context,
	req *proto.UnblockPersonRequest) (*proto.UnblockPersonResponse, error) {
	//
	personId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "id must be UUID")
	}

	err = s.commandService.UnblockPerson(ctx, commands.UnblockPersonCommand{
		Id: personId,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.UnblockPersonResponse{}, nil
}

func New(r app.Repository,
	qs queries.PersonQueryService,
	cs commands.PersonCommandService) *PersonServer {
	//
	return &PersonServer{personsRepository: r, queryService: qs, commandService: cs}
}
