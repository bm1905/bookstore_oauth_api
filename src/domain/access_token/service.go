package access_token

import (
	"github.com/bm1905/bookstore_oauth_api/src/utils/errors_utils"
	"strings"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors_utils.RestError)
	Create(AccessToken) *errors_utils.RestError
	UpdateExpirationTime(AccessToken) *errors_utils.RestError
}

type Service interface {
	GetById(string) (*AccessToken, *errors_utils.RestError)
	Create(AccessToken) *errors_utils.RestError
	UpdateExpirationTime(AccessToken) *errors_utils.RestError
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *errors_utils.RestError) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors_utils.NewBadRequestError("invalid access token id")
	}
	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(at AccessToken) *errors_utils.RestError {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.Create(at)
}

func (s *service) UpdateExpirationTime(at AccessToken) *errors_utils.RestError {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(at)
}
