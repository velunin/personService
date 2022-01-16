package domain

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPerson_ShouldConstructPersonWithExpectedState(t *testing.T) {
	id := uuid.New()
	firstName := "Some"
	lastName := "Person"

	person, err := NewPerson(id, firstName, lastName)

	assert.NoError(t, err)
	assert.Equal(t, id, person.state.Id)
	assert.Equal(t, firstName, person.state.FirstName)
	assert.Equal(t, lastName, person.state.LastName)
	assert.False(t, person.state.IsBlocked)
	assert.NotEmpty(t, person.events)

	event, ok := person.events[0].(PersonCreatedEvent)

	assert.Truef(t, ok, "unexpected event type")
	assert.Equal(t, id, event.Id)
	assert.Equal(t, firstName, event.FirstName)
	assert.Equal(t, lastName, event.LastName)
}

func TestNewPerson_ShouldReturnErrorWhenPassedInvalidArgs(t *testing.T) {
	_, err := NewPerson(uuid.New(), "", "Person")

	assert.Equal(t, ErrFirstNameEmpty, err)

	_, err = NewPerson(uuid.New(), "Some", "")

	assert.Equal(t, ErrLastNameEmpty, err)
}

func TestPersonAggregate_ChangeName_ShouldChangeStateExpectedWay(t *testing.T) {
	person := restore(false)

	expectedFirstName := "FirstName"
	expectedLastName := "LastName"

	err := person.ChangeName(expectedFirstName, expectedLastName)

	assert.NoError(t, err)
	assert.Equal(t, expectedFirstName, person.state.FirstName)
	assert.Equal(t, expectedLastName, person.state.LastName)

	event, ok := person.events[0].(PersonNameChangedEvent)

	assert.Truef(t, ok, "unexpected event type")
	assert.Equal(t, person.state.Id, event.Id)
	assert.Equal(t, expectedFirstName, event.FirstName)
	assert.Equal(t, expectedLastName, event.LastName)
}

func TestPersonAggregate_ChangeName_ShouldReturnErrorWhenPassedInvalidArgs(t *testing.T) {
	person, _ := NewPerson(uuid.New(), "Some", "Person")

	err := person.ChangeName("", "Person")

	assert.Equal(t, ErrFirstNameEmpty, err)

	err = person.ChangeName("Some", "")

	assert.ErrorIs(t, err, ErrLastNameEmpty)
}

func TestPersonAggregate_ChangeName_ShouldReturnErrorWhenPersonIsBlocked(t *testing.T) {
	person, _ := NewPerson(uuid.New(), "Some", "Person")
	_ = person.Block()

	err := person.ChangeName("FirstName", "LastName")

	assert.ErrorIs(t, err, ErrPersonBlocked)
}

func TestPersonAggregate_Block_ShouldChangeStateExpectedWay(t *testing.T) {
	person := restore(false)

	err := person.Block()

	assert.NoError(t, err)
	assert.True(t, person.state.IsBlocked)

	event, ok := person.events[0].(PersonBlockedEvent)

	assert.Truef(t, ok, "unexpected event type")
	assert.Equal(t, person.state.Id, event.Id)
}

func TestPersonAggregate_Block_ShouldReturnErrorWhenPersonAlreadyBlocked(t *testing.T) {
	person, _ := NewPerson(uuid.New(), "Some", "Person")
	_ = person.Block()

	err := person.Block()

	assert.ErrorIs(t, err, ErrAlreadyPersonBlocked)
}

func TestPersonAggregate_Unblock_ShouldChangeStateExpectedWay(t *testing.T) {
	person := restore(true)

	err := person.Unblock()

	assert.NoError(t, err)
	assert.False(t, person.state.IsBlocked)

	event, ok := person.events[0].(PersonUnblockedEvent)

	assert.Truef(t, ok, "unexpected event type")
	assert.Equal(t, person.state.Id, event.Id)
}

func TestPersonAggregate_Unblock_ShouldReturnErrorWhenPersonAlreadyUnblocked(t *testing.T) {
	person, _ := NewPerson(uuid.New(), "Some", "Person")

	err := person.Unblock()

	assert.ErrorIs(t, err, ErrAlreadyPersonUnblocked)
}

func restore(isBlocked bool) *PersonAggregate {
	return RestorePerson(&PersonState{
		Id:        uuid.New(),
		FirstName: "Some",
		LastName:  "Person",
		IsBlocked: isBlocked,
	})
}
