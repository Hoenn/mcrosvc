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
	tx, err := d.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return -1, err
	}
	var id int64
	res, err := createInsertUserQuery(u).RunWith(tx).ExecContext(ctx)
	if err != nil {
		return -1, errors.Wrap(err, "could not insert user")
	}
	id, err = res.LastInsertId()
	if err != nil {
		return -1, errors.Wrap(err, "could not get id from insert")
	}
	err = tx.Commit()
	if err != nil {
		return -1, errors.Wrap(err, "could not user")
	}

	return id, err
}

func (d *UserAPI) DeleteUser(ctx context.Context, u int32) error {
	tx, err := d.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	_, err = createDeleteUserQuery(u).RunWith(tx).ExecContext(ctx)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		return errors.Wrap(err, "unable to delete user")
	}
	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "unable to delete user")
	}

	return nil
}

func (d *UserAPI) GetUser(ctx context.Context, u int32) (*proto.User, error) {
	var usr proto.User
	tx, err := d.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	row := createSelectUserQuery(u).RunWith(tx).QueryRowContext(ctx)
	err = row.Scan(
		&usr.UserNum,
		&usr.Name,
		&usr.Age,
	)
	if err != nil {
		return nil, errors.Wrap(err, "unable to scan user")
	}
	err = tx.Commit()
	if err != nil {
		return nil, errors.Wrap(err, "unable to scan user")
	}
	return &usr, nil
}
