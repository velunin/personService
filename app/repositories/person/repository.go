package person

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"go.uber.org/fx"
	"personService/app/repositories"
	"personService/domain"
)

type Repository interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.Person, error)
	Insert(ctx context.Context, person *domain.Person) error
	Update(ctx context.Context, person *domain.Person) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type personRepository struct {
	RepoParams
}

type RepoParams struct {
	fx.In
	Tx repositories.Transaction
}

func (r *personRepository) Insert(ctx context.Context, person *domain.Person) error {
	const query = `INSERT INTO persons(id,first_name,last_name,is_blocked) VALUES($1,$2,$3,$4)`

	state := person.State()
	_, err := r.Tx.GetDB(ctx).Exec(query, state.Id, state.FirstName, state.LastName, state.IsBlocked)

	return err
}

func (r *personRepository) Update(ctx context.Context, person *domain.Person) error {
	const query = `UPDATE persons SET first_name=$2, last_name=$3, is_blocked=$4 WHERE id=$1`

	state := person.State()
	_, err := r.Tx.GetDB(ctx).Exec(query, state.Id, state.FirstName, state.LastName, state.IsBlocked)

	return err
}

func (r *personRepository) Delete(ctx context.Context, id uuid.UUID) error {
	const query = `DELETE FROM persons id=$1`

	_, err := r.Tx.GetDB(ctx).Exec(query, id)

	return err
}

func (r *personRepository) Get(ctx context.Context, id uuid.UUID) (*domain.Person, error) {
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

func NewPersonRepository(params RepoParams) Repository {
	return &personRepository{params}
}

var (
	ErrNotFound = errors.New("entity not found")
)
