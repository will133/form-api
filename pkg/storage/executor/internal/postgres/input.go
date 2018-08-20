package postgres

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"

	"github.com/kamilsk/form-api/pkg/domain"
	"github.com/kamilsk/form-api/pkg/errors"
	"github.com/kamilsk/form-api/pkg/storage/query"
)

// NewInputContext TODO
func NewInputContext(ctx context.Context, conn *sql.Conn) inputScope {
	return inputScope{ctx, conn}
}

type inputScope struct {
	ctx  context.Context
	conn *sql.Conn
}

// Write TODO
func (scope inputScope) Write(data query.WriteInput) (query.Input, error) {
	var entity = query.Input{SchemaID: data.SchemaID, Data: data.VerifiedData}
	encoded, encodeErr := json.Marshal(data.VerifiedData)
	if encodeErr != nil {
		return entity, errors.Serialization(errors.ServerErrorMessage, encodeErr,
			"trying to marshal data `%#v` for the schema %q into JSON",
			data.VerifiedData, data.SchemaID)
	}
	q := `INSERT INTO "input" ("schema_id", "data") VALUES ($1, $2) RETURNING "id", "created_at"`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.SchemaID, encoded)
	if scanErr := row.Scan(&entity.ID, &entity.CreatedAt); scanErr != nil {
		return entity, errors.Database(errors.ServerErrorMessage, scanErr,
			"trying to insert input `%s` for the schema %q",
			encoded, entity.SchemaID)
	}
	return entity, nil
}

// ReadByID TODO
// TODO check access
func (scope inputScope) ReadByID(token *query.Token, id domain.ID) (query.Input, error) {
	var entity, encoded = query.Input{ID: id}, []byte(nil)
	q := `SELECT "schema_id", "data", "created_at" FROM "input" WHERE "id" = $1`
	row := scope.conn.QueryRowContext(scope.ctx, q, entity.ID)
	if err := row.Scan(&entity.SchemaID, &encoded, &entity.CreatedAt); err != nil {
		return entity, errors.Database(errors.ServerErrorMessage, err,
			"user %q of account %q tried to read the input %q",
			token.UserID, token.User.AccountID, id)
	}
	if err := json.Unmarshal(encoded, &entity.Data); err != nil {
		return entity, errors.Serialization(errors.ServerErrorMessage, err,
			"user %q of account %q tried to unmarshal JSON `%s` of the input %q",
			token.UserID, token.User.AccountID, encoded, id)
	}
	return entity, nil
}

// ReadByFilter TODO
// TODO check access
func (scope inputScope) ReadByFilter(token *query.Token, filter query.InputFilter) ([]query.Input, error) {
	args := append(make([]interface{}, 0, 3), filter.SchemaID)
	// TODO go 1.10 builder := strings.Builder{}
	builder := bytes.NewBuffer(make([]byte, 0, 82+35))
	_, _ = builder.WriteString(`SELECT "id", "data", "created_at" FROM "input" WHERE "schema_id" = $1`)
	switch {
	case !filter.From.IsZero() && !filter.To.IsZero():
		_, _ = builder.WriteString(` AND "created_at" BETWEEN $2 AND $3`)
		args = append(args, filter.From, filter.To)
	case !filter.From.IsZero():
		_, _ = builder.WriteString(` AND "created_at" >= $2`)
		args = append(args, filter.From)
	case !filter.To.IsZero():
		_, _ = builder.WriteString(` AND "created_at" <= $2`)
		args = append(args, filter.To)
	}
	var entities = make([]query.Input, 0, 8)
	rows, dbErr := scope.conn.QueryContext(scope.ctx, builder.String(), args...)
	if dbErr != nil {
		return nil, errors.Database(errors.ServerErrorMessage, dbErr,
			"user %q of account %q tried to read inputs by criteria %+v",
			token.UserID, token.User.AccountID, filter)
	}
	for rows.Next() {
		var entity, encoded = query.Input{SchemaID: filter.SchemaID}, []byte(nil)
		if scanErr := rows.Scan(&entity.ID, &encoded, &entity.CreatedAt); scanErr != nil {
			return nil, errors.Database(errors.ServerErrorMessage, scanErr,
				"user %q of account %q tried to read inputs by criteria %+v",
				token.UserID, token.User.AccountID, filter)
		}
		if decodeErr := json.Unmarshal(encoded, &entity.Data); decodeErr != nil {
			return nil, errors.Serialization(errors.ServerErrorMessage, decodeErr,
				"user %q of account %q tried to unmarshal JSON `%s` of the input %q",
				token.UserID, token.User.AccountID, encoded, entity.ID)
		}
		entities = append(entities, entity)
	}
	return entities, nil
}
