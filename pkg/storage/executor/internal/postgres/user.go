package postgres

import (
	"context"
	"database/sql"

	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/kamilsk/form-api/pkg/errors"
	"github.com/kamilsk/form-api/pkg/storage/types"
)

// NewUserContext TODO issue#173
func NewUserContext(ctx context.Context, conn *sql.Conn) userScope {
	return userScope{ctx, conn}
}

type userScope struct {
	ctx  context.Context
	conn *sql.Conn
}

// Token TODO issue#173
func (scope userScope) Token(id domain.ID) (*types.Token, error) {
	var (
		token   = types.Token{ID: id}
		user    = types.User{}
		account = types.Account{}
	)
	q := `SELECT "t"."user_id", "t"."expired_at", "t"."created_at",
	             "u"."account_id", "u"."name", "u"."created_at", "u"."updated_at",
	             "a"."name", "a"."created_at", "a"."updated_at"
	        FROM "token" "t"
	  INNER JOIN "user" "u" ON "u"."id" = "t"."user_id"
	  INNER JOIN "account" "a" ON "a"."id" = "u"."account_id"
	       WHERE "t"."id" = $1 AND ("t"."expired_at" IS NULL OR "t"."expired_at" > now())
	         AND "u"."deleted_at" IS NULL
	         AND "a"."deleted_at" IS NULL`
	row := scope.conn.QueryRowContext(scope.ctx, q, token.ID)
	if err := row.Scan(
		&token.UserID, &token.ExpiredAt, &token.CreatedAt,
		&user.AccountID, &user.Name, &user.CreatedAt, &user.UpdatedAt,
		&account.Name, &account.CreatedAt, &account.UpdatedAt,
	); err != nil {
		return nil, errors.Database(errors.ServerErrorMessage, err, "trying to fetch credentials by the token %q", id)
	}
	user.ID, account.ID = token.UserID, user.AccountID
	token.User, user.Account = &user, &account
	user.Tokens, account.Users = append(user.Tokens, &token), append(account.Users, &user)
	return &token, nil
}
