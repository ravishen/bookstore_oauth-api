package db

import (
	"fmt"

	"github.com/gocql/gocql"
	"github.com/ravishen/bookstore_oauth-api/src/clients/cassandra"
	"github.com/ravishen/bookstore_oauth-api/src/domain/access_token"
	"github.com/ravishen/bookstore_oauth-api/src/utils/errors"
)

const (
	queryGetAccessToken       = "select access_token, user_id, client_id, expires from access_tokens where access_token=?;"
	qureyCreateAccessToken    = "insert into access_tokens(access_token, user_id, client_id, expires ) values (?,?,?,?);"
	queryUpdateExpirationTime = "update access_tokens set expires=? where access_token=?; "
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}
type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer session.Close()
	var result access_token.AccessToken
	if err := session.Query(queryGetAccessToken, id).Scan(&result.UserId, &result.ClientId, &result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError(fmt.Sprintf("no access token found with id %s", id))
		}
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &result, nil
}
func NewRepository() DbRepository {
	return &dbRepository{}
}
func (r *dbRepository) Create(at access_token.AccessToken) *errors.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	if er := session.Query(qureyCreateAccessToken, at.AccessToken, at.UserId, at.ClientId, at.Expires).Exec(); er != nil {
		return errors.NewInternalServerError(er.Error())
	}
	defer session.Close()
	return nil
}
func (r *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	if err := session.Query(queryUpdateExpirationTime, at.Expires, at.AccessToken).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer session.Close()
	return nil
}
