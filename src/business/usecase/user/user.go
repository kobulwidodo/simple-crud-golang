package user

import (
	userDom "crud-user/src/business/domain/user"
	"crud-user/src/business/entity"
)

type Interface interface {
	Create(user entity.User) (entity.User, error)
	GetList() ([]entity.User, error)
	Update(id uint, updateParam entity.UserUpdateParam) (entity.User, error)
	Delete(id uint) error
}

type user struct {
	user userDom.Interface
}

func Init(ud userDom.Interface) Interface {
	u := &user{
		user: ud,
	}

	return u
}

func (u *user) Create(user entity.User) (entity.User, error) {
	newPond, err := u.user.Create(user)
	if err != nil {
		return newPond, err
	}

	return newPond, nil
}

func (u *user) GetList() ([]entity.User, error) {
	users, err := u.user.GetList()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (u *user) Update(id uint, updateParam entity.UserUpdateParam) (entity.User, error) {
	user, err := u.user.Get(id)
	if err != nil {
		return user, err
	}

	user.Username = updateParam.Username
	user.Password = updateParam.Password

	newUser, err := u.user.Update(user)
	if err != nil {
		return user, err
	}

	return newUser, nil
}

func (u *user) Delete(id uint) error {
	if err := u.user.Delete(id); err != nil {
		return err
	}

	return nil
}
