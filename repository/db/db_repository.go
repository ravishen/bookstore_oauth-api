package db

import (
	"github.com/ravishen/bookstore_oauth-api/src/domain/access_token"
	"github.com/ravishen/bookstore_oauth-api/src/utils/errors"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}
type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("datbase not connected")
}
func NewRepository() DbRepository {
	return &dbRepository{}
}
