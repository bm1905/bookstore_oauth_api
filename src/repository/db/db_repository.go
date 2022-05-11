package db

import (
	"github.com/bm1905/bookstore_oauth_api/src/clients/cassandra"
	"github.com/bm1905/bookstore_oauth_api/src/domain/access_token"
	"github.com/bm1905/bookstore_oauth_api/src/utils/errors_utils"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES(?, ?, ?, ?);"
	queryUpdateExpires     = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

func NewRepository() Repository {
	return &repository{}
}

type Repository interface {
	GetById(string) (*access_token.AccessToken, *errors_utils.RestError)
	Create(access_token.AccessToken) *errors_utils.RestError
	UpdateExpirationTime(access_token.AccessToken) *errors_utils.RestError
}

type repository struct {
}

func (r *repository) GetById(id string) (*access_token.AccessToken, *errors_utils.RestError) {
	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors_utils.NewInternalServerError("no access token found")
		}
		return nil, errors_utils.NewInternalServerError(err.Error())
	}

	return &result, nil
}

func (r *repository) Create(at access_token.AccessToken) *errors_utils.RestError {
	if err := cassandra.GetSession().Query(queryCreateAccessToken, at.AccessToken, at.UserId, at.ClientId, at.Expires).
		Exec(); err != nil {
		return errors_utils.NewInternalServerError(err.Error())
	}

	return nil
}

func (r *repository) UpdateExpirationTime(at access_token.AccessToken) *errors_utils.RestError {
	if err := cassandra.GetSession().Query(queryUpdateExpires, at.Expires, at.AccessToken).
		Exec(); err != nil {
		return errors_utils.NewInternalServerError(err.Error())
	}

	return nil
}
