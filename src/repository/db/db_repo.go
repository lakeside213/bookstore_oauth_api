package db

import (
	"microservice_tut/bookstore_oauth_api/src/domain/access_token"
	"microservice_tut/bookstore_oauth_api/src/utils/errors"
)

func NewRepo() DbRepository {
	return &dbRepo{}
}

type DbRepository interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepo struct {
}

func (db *dbRepo) GetByID(a string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("database connection failed")
	// return nil, nil
}
