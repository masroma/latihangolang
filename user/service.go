package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	// buat nama function
	RegisterUser(input RegisteruserInput) (User, error)
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
