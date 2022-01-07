package domain

import (
	"errors"
	"github.com/google/uuid"
)

type Person struct {
	state  PersonState
	events []interface{}
}

type PersonState struct {
	Id        uuid.UUID `db:"id"`
	FirstName string    `db:"first_name""`
	LastName  string    `db:"last_name"`
	IsBlocked bool      `db:"is_blocked"`
}

func (p *Person) State() PersonState {
	return p.state
}

func (p *Person) ChangeName(firstName, lastName string) error {
	if p.state.IsBlocked == true {
		return ErrPersonBlocked
	}

	err := validateNames(firstName, lastName)
	if err != nil {
		return err
	}

	p.state.FirstName = firstName
	p.state.LastName = lastName

	return nil
}

func (p *Person) Block() error {
	if p.state.IsBlocked {
		return ErrAlreadyPersonBlocked
	}

	p.state.IsBlocked = true

	return nil
}

func (p *Person) Unblock() error {
	if p.state.IsBlocked == false {
		return ErrAlreadyPersonUnblocked
	}

	p.state.IsBlocked = false

	return nil
}

func NewPerson(id uuid.UUID, firstName, lastName string) (*Person, error) {
	err := validateNames(firstName, lastName)
	if err != nil {
		return nil, err
	}

	person := &Person{state: PersonState{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
		IsBlocked: false,
	}}

	person.apply(PersonCreated{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
	})

	return person, nil
}

func RestorePerson(state *PersonState) *Person {
	return &Person{state: *state}
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

func (p *Person) apply(event interface{}) {
	p.events = append(p.events, event)
}

func (p *Person) GetEvents() []interface{} {
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

type PersonCreated struct {
	Id        uuid.UUID
	FirstName string
	LastName  string
}
