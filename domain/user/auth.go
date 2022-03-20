package user

import (
	"errors"
	"net/http"

	"blog-gin_golang_v177/domain/user/model"
	"blog-gin_golang_v177/domain/user/repository"
	"blog-gin_golang_v177/lib/constant"
	"blog-gin_golang_v177/lib/encrypt"
)

type AuthServiceInterface interface {
	Register(reqBody *model.ReqBodyRegister) (int, error)
}

type authService struct {
	Repository repository.AuthRepositoryInterface
}

func AuthService(repository repository.AuthRepositoryInterface) AuthServiceInterface {
	return &authService{
		Repository: repository,
	}
}

func (s *authService) Register(reqBody *model.ReqBodyRegister) (int, error) {
	if _, err := s.Repository.FirstByEmail(&reqBody.Email); err == nil {
		return http.StatusBadRequest, errors.New(constant.EmailAlreadyExists)
	}
	if err := encrypt.GenerateFromPassword(&reqBody.Password); err != nil {
		return http.StatusInternalServerError, err
	}

	err := s.Repository.Create(reqBody)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
