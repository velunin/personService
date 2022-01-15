package app

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"go.uber.org/fx"
	"personService/internal/database"
	"personService/internal/dispatcher"
	"personService/modules/persons/domain"
)

type Repository interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.PersonAggregate, error)
	Insert(ctx context.Context, person *domain.PersonAggregate) error
	Update(ctx context.Context, person *domain.PersonAggregate) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type personRepository struct {
	RepoParams
}

type RepoParams struct {
	fx.In
	Tx              database.Transaction
	EventDispatcher dispatcher.Dispatcher
}

func (r *personRepository) Insert(ctx context.Context, person *domain.PersonAggregate) error {
	const query = `INSERT INTO persons(id,first_name,last_name,is_blocked) VALUES($1,$2,$3,$4)`

	state := person.State()
	_, err := r.Tx.GetDB(ctx).Exec(query, state.Id, state.FirstName, state.LastName, state.IsBlocked)
	if err != nil {
		return err
	}

	err = r.dispatchEvents(ctx, person)

	return err
}

func (r *personRepository) Update(ctx context.Context, person *domain.PersonAggregate) error {
	const query = `UPDATE persons SET first_name=$2, last_name=$3, is_blocked=$4 WHERE id=$1`

	state := person.State()
	_, err := r.Tx.GetDB(ctx).Exec(query, state.Id, state.FirstName, state.LastName, state.IsBlocked)

	err = r.dispatchEvents(ctx, person)

	return err
}

func (r *personRepository) Delete(ctx context.Context, id uuid.UUID) error {
	const query = `DELETE FROM persons id=$1`

	_, err := r.Tx.GetDB(ctx).Exec(query, id)

	return err
}

func (r *personRepository) Get(ctx context.Context, id uuid.UUID) (*domain.PersonAggregate, error) {
	const query = `SELECT * FROM persons WHERE id = $1 FOR UPDATE`

	state := domain.PersonState{}
	err := r.Tx.GetDB(ctx).Get(&state, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return domain.RestorePerson(&state), nil
}

func (r *personRepository) dispatchEvents(ctx context.Context, person *domain.PersonAggregate) error {
	for _, event := range person.GetEvents() {
		err := r.EventDispatcher.Dispatch(ctx, event)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewRepository(params RepoParams) Repository {
	return &personRepository{params}
}

var (
	ErrNotFound = errors.New("entity not found")
)
