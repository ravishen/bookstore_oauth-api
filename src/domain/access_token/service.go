package access_token

import (
	"github.com/ravishen/bookstore_oauth-api/src/utils/errors"
)

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}
type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}
type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *errors.RestErr) {
	accessToken, err := s.repository.GetById(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("invalid access token")
	}
	if err != nil {
		return nil, err
	}
	accessToken, err = s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(accessToken AccessToken) *errors.RestErr {
	err := accessToken.Validate()
	if err != nil {
		return err
	}
	return s.repository.Create(accessToken)

	//	if error != nil {
	//		return errors.NewInternalServerError(error.Error)
	//	}
	//
	//	return nil
}

func (s *service) UpdateExpirationTime(at AccessToken) *errors.RestErr {
	err := at.Validate()
	if err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(at)
	//	if err != nil {
	//		return errors.NewInternalServerError("unable to create access token")
	//	}
	//	return nil
}
