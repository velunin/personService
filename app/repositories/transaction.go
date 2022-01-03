package repositories

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type transaction struct {
	txKey tranKey
	db    *sqlx.DB
}

type TransactionHandler func(ctx context.Context) error

type tranKey struct{}

func (a *transaction) ExecInTran(baseContext context.Context, th TransactionHandler) (err error) {
	var tx *sqlx.Tx
	tx, err = a.db.Beginx()
	ctx := context.WithValue(baseContext, a.txKey, tx)
	defer func() {
		if r := recover(); r != nil || err != nil {
			tx.Rollback() // nolint:nolintlint,errcheck,gosec // no need
		}
	}()

	if err == nil {
		err = th(ctx)
	}

	if err == nil {
		err = tx.Commit()
	}

	return
}

func (a *transaction) GetDB(ctx context.Context) DB {
	o := ctx.Value(a.txKey)
	if o != nil {
		if tx, ok := o.(*sqlx.Tx); ok {
			return tx
		}
	}

	return a.db
}

func NewTransaction(_db *sqlx.DB) Transaction {
	ts := &transaction{
		db:    _db,
		txKey: tranKey{},
	}

	return ts
}

type Transaction interface {
	ExecInTran(baseContext context.Context, th TransactionHandler) (err error)
	GetDB(ctx context.Context) DB
}

type DB interface {
	sqlx.Execer
	sqlx.Queryer

	Select(dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error
	NamedExec(query string, arg interface{}) (sql.Result, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}
