package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	// buat nama function
	RegisterUser(input RegisteruserInput) (User, error)
	Login(input LoginInput) (User, error)
	EmailAvailable(input CheckEmailInput) (bool, error)
	UpdateAvatar(ID int, fileLocation string) (User, error)
}

type service struct {
	// kita panggil repository krn akan di parsing k function repository
	repository Repository
}

func NewService(repository Repository) *service {
	// ini krn kita butuh akses db yg ada di repository maka kita panggil repository nya
	return &service{repository}
}

func (s *service) RegisterUser(input RegisteruserInput) (User, error) {
	// function ini mengembalikan data yg udah di map service ke function repository save
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) EmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil

}

func (s *service) UpdateAvatar(ID int, fileLocation string) (User, error) {
	user, err := s.repository.FindByID(ID)

	if err != nil {
		return user, err
	}

	user.Photo = fileLocation
	updateUser, err := s.repository.Update(user)
	if err != nil {
		return updateUser, err
	}

	return updateUser, nil

}
