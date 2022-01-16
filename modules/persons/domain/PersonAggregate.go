package domain

import (
	"errors"
	"github.com/google/uuid"
)

type PersonAggregate struct {
	state  PersonState
	events []interface{}
}

type PersonState struct {
	Id        uuid.UUID `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	IsBlocked bool      `db:"is_blocked"`
}

func (p *PersonAggregate) State() PersonState {
	return p.state
}

func (p *PersonAggregate) ChangeName(firstName, lastName string) error {
	if p.state.IsBlocked == true {
		return ErrPersonBlocked
	}

	err := validateNames(firstName, lastName)
	if err != nil {
		return err
	}

	p.state.FirstName = firstName
	p.state.LastName = lastName

	p.apply(PersonNameChangedEvent{
		Id:        p.state.Id,
		FirstName: firstName,
		LastName:  lastName})

	return nil
}

func (p *PersonAggregate) Block() error {
	if p.state.IsBlocked {
		return ErrAlreadyPersonBlocked
	}

	p.state.IsBlocked = true

	p.apply(PersonBlockedEvent{Id: p.state.Id})

	return nil
}

func (p *PersonAggregate) Unblock() error {
	if p.state.IsBlocked == false {
		return ErrAlreadyPersonUnblocked
	}

	p.state.IsBlocked = false

	p.apply(PersonUnblockedEvent{Id: p.state.Id})

	return nil
}

func NewPerson(id uuid.UUID, firstName, lastName string) (*PersonAggregate, error) {
	err := validateNames(firstName, lastName)
	if err != nil {
		return nil, err
	}

	person := &PersonAggregate{state: PersonState{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
		IsBlocked: false,
	}}

	person.apply(PersonCreatedEvent{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
	})

	return person, nil
}

func RestorePerson(state *PersonState) *PersonAggregate {
	return &PersonAggregate{state: *state}
}

func validateNames(firstName, lastName string) error {
	if firstName == "" {
		return ErrFirstNameEmpty
	}
	if lastName == "" {
		return ErrLastNameEmpty
	}

	return nil
}

func (p *PersonAggregate) apply(event interface{}) {
	p.events = append(p.events, event)
}

func (p *PersonAggregate) GetEvents() []interface{} {
	return p.events
}

// Domain errors
var (
	ErrPersonBlocked          = errors.New("person blocked")
	ErrAlreadyPersonBlocked   = errors.New("person already blocked")
	ErrAlreadyPersonUnblocked = errors.New("person already unblocked")
	ErrFirstNameEmpty         = errors.New("first name required")
	ErrLastNameEmpty          = errors.New("last name required")
)

type PersonCreatedEvent struct {
	Id        uuid.UUID
	FirstName string
	LastName  string
}

type PersonNameChangedEvent struct {
	Id        uuid.UUID
	FirstName string
	LastName  string
}

type PersonBlockedEvent struct {
	Id uuid.UUID
}

type PersonUnblockedEvent struct {
	Id uuid.UUID
}
