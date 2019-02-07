package db

import (
	"github.com/hoenn/mcrosvc/proto"

	sq "github.com/Masterminds/squirrel"
)

func createInsertUserQuery(u *proto.User) sq.InsertBuilder {
	return sq.Insert(
		"users",
	).Columns(
		"username",
		"age",
	).Values(
		u.Name,
		u.Age,
	)
}

func createDeleteUserQuery(id int32) sq.DeleteBuilder {
	return sq.Delete(
		"users",
	).Where(
		sq.Eq{
			"id": id,
		},
	)
}

func createSelectUserQuery(id int32) sq.SelectBuilder {
	return sq.Select(
		"id",
		"username",
		"age",
	).From(
		"users",
	).Where(
		sq.Eq{
			"id": id,
		},
	)

}
