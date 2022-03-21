package repository

import (
	"gorm.io/gorm"

	"blog-gin_golang_v177/domain/user/model"
)

type AuthRepositoryInterface interface {
	FirstByEmail(email *string) (model.User, error)
	Create(reqBody *model.ReqBodyRegister) error
	UpdateToken(model.User, string) error
}

type authRepository struct {
	DB *gorm.DB
}

func AuthRepository(DB *gorm.DB) AuthRepositoryInterface {
	return &authRepository{
		DB: DB,
	}
}

func (r *authRepository) FirstByEmail(email *string) (user model.User, err error) {
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *authRepository) Create(reqBody *model.ReqBodyRegister) error {
	var user model.User
	user.Name = reqBody.Name
	user.Email = reqBody.Email
	user.EncryptedPassword = reqBody.Password
	if reqBody.RoleID == "" {
		user.RoleID = "1ca78238-8e40-4763-a20f-59b5b41791b1" // default superadmin
	}
	return r.DB.Create(&user).Error
}

func (r *authRepository) UpdateToken(user model.User, ss string) error {
	user.Token = ss
	return r.DB.Save(&user).Error
}
