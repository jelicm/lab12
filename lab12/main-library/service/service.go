package service

import (
	"ccmainproject/domain"
	"ccmainproject/store"
	"errors"
)

type UsersService struct {
	store store.UsersMongoDBStore
}

func NewUsersService(store store.UsersMongoDBStore) (UsersService, error) {
	return UsersService{
		store: store,
	}, nil
}

func (service *UsersService) GetAll() ([]*domain.User, error) {
	return service.store.GetAll()
}

func (service *UsersService) Register(user *domain.User) (int, error) {
	user.BooksNum = 0

	_, err := service.store.GetByJmbg(user.Jmbg)
	if err == nil {
		return 0, errors.New("User with provided jmbg already exists!")
	}

	users, err := service.store.GetAll()
	maxUserId := 0

	for _, u := range users {
		if u.UserId > maxUserId {
			maxUserId = u.UserId
		}
	}

	user.UserId = maxUserId + 1
	err = service.store.Insert(user)
	return user.UserId, err
}

func (service *UsersService) Borrow(userId int) (int, error) {

	user, err := service.store.GetByUserId(userId)
	if err != nil {
		return 0, err
	}

	if user.BooksNum >= 3 {
		return 0, errors.New("Already borrowed 3 books")
	}

	user.BooksNum += 1
	err = service.store.UpdateBooksNum(user)
	if err != nil {
		return 0, err
	}

	return user.BooksNum, err
}

func (service *UsersService) Return(userId int) (int, error) {

	user, err := service.store.GetByUserId(userId)
	if err != nil {
		return 0, err
	}

	if user.BooksNum == 0 {
		return 0, errors.New("User has not borrowed any book")
	}

	user.BooksNum -= 1
	err = service.store.UpdateBooksNum(user)
	if err != nil {
		return 0, err
	}

	return user.BooksNum, err
}
