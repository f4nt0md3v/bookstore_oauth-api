package db

import (
	"github.com/gocql/gocql"

	"github.com/f4nt0md3v/bookstore_oauth-api/src/client/cassandra"
	"github.com/f4nt0md3v/bookstore_oauth-api/src/domain/access_token"
	"github.com/f4nt0md3v/bookstore_oauth-api/src/utils/errors"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token = ?;"
	queryCreateAccessToken = "INSERT INTO access_tokens (access_token, user_id, client_id, expires) VALUES(?, ?, ?, ?);"
	queryUpdateExpires     = "UPDATE access_tokens SET expires = ? WHERE access_token = ?;"
)

func NewRepository() DatabaseRepository {
	return &dbRepository{}
}

type DatabaseRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestError)
	Create(access_token.AccessToken) *errors.RestError
	UpdateExpirationTime(access_token.AccessToken) *errors.RestError
}

type dbRepository struct {}

func (r *dbRepository) Create(token access_token.AccessToken) *errors.RestError {
	if err := cassandra.GetSession().Query(
		queryCreateAccessToken,
		token.AccessToken,
		token.UserId,
		token.ClientId,
		token.Expires,
	).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (r *dbRepository) UpdateExpirationTime(token access_token.AccessToken) *errors.RestError {
	if err := cassandra.GetSession().Query(
		queryUpdateExpires,
		&token.Expires,
		&token.AccessToken,
	).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestError) {
	var token access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&token.AccessToken,
		&token.UserId,
		&token.ClientId,
		&token.Expires,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("no access token found with given id")
		}
		return nil, errors.NewInternalServerError(err.Error())
	}

	return &token, nil
}
