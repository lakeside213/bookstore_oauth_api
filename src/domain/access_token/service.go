package access_token

import (
	"microservice_tut/bookstore_oauth_api/src/utils/errors"
	"strings"
)

type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}

type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetByID(token string) (*AccessToken, *errors.RestErr) {
	token = strings.TrimSpace(token)

	if len(token) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}

	accessToken, err := s.repository.GetByID(token)
	if err != nil {
		// c.JSON(err.Status, err)
		return nil, err
	}
	// c.JSON(http.StatusOK, accessToken)
	return accessToken, nil
}
