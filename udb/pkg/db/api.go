package db

import (
	"context"
	"database/sql"

	//mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"

	"github.com/hoenn/mcrosvc/proto"
)

type UserAPI struct {
	db *sql.DB
}

func NewUserAPI(database *sql.DB) *UserAPI {
	return &UserAPI{
		db: database,
	}
}

func (d *UserAPI) Close() error {
	return d.db.Close()
}

func (d *UserAPI) CreateUser(ctx context.Context, u *proto.User) (int64, error) {
	var id int64
	_, err := WithAutomaticCommit(ctx, d.db, func(tx *sql.Tx) (sql.Result, error) {
		res, err := createInsertUserQuery(u).RunWith(tx).ExecContext(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "could not insert user")
		}
		id, err = res.LastInsertId()
		return nil, nil
	})
	if err != nil {
		return -1, errors.Wrap(err, "could not get id from insert")
	}

	return id, err
}

func (d *UserAPI) DeleteUser(ctx context.Context, u int32) error {
	_, err := WithAutomaticCommit(ctx, d.db, func(tx *sql.Tx) (sql.Result, error) {
		_, err := createDeleteUserQuery(u).RunWith(tx).ExecContext(ctx)
		if err == sql.ErrNoRows {
			return nil, nil
		}
		if err != nil {
			return nil, errors.Wrap(err, "unable to delete user")
		}
		return nil, nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (d *UserAPI) GetUser(ctx context.Context, u int32) (*proto.User, error) {
	var usr proto.User
	_, err := WithAutomaticCommit(ctx, d.db, func(tx *sql.Tx) (sql.Result, error) {
		row := createSelectUserQuery(u).RunWith(tx).QueryRowContext(ctx)
		err := row.Scan(
			&usr.UserNum,
			&usr.Name,
			&usr.Age,
		)
		if err != nil {
			return nil, errors.Wrap(err, "unable to scan user")
		}
		return nil, err
	})
	if err != nil {
		return nil, err
	}
	return &usr, nil
}

// WithAutomaticCommit automatically commits after given func if no errors were returned.
func WithAutomaticCommit(ctx context.Context, db *sql.DB, fn func(*sql.Tx) (sql.Result, error)) (sql.Result, error) {

	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	result, err := fn(tx)

	if err != nil {

		err2 := tx.Rollback()
		if err2 != nil {
		}
	} else {
		err2 := tx.Commit()
		if err2 != nil {
		}
	}

	return result, err
}
