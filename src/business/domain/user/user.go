package user

import (
	"crud-user/src/business/entity"

	"gorm.io/gorm"
)

type Interface interface {
	Create(user entity.User) (entity.User, error)
	GetList() ([]entity.User, error)
	Get(id uint) (entity.User, error)
	Update(user entity.User) (entity.User, error)
	Delete(id uint) error
}

type user struct {
	db *gorm.DB
}

func Init(db *gorm.DB) Interface {
	u := &user{
		db: db,
	}

	return u
}

func (u *user) Create(user entity.User) (entity.User, error) {
	tx := u.db.Begin()
	defer tx.Rollback()

	if err := tx.Create(&user).Error; err != nil {
		return user, err
	}

	if err := tx.Commit().Error; err != nil {
		return user, err
	}

	return user, nil
}

func (u *user) Get(id uint) (entity.User, error) {
	user := entity.User{}

	if err := u.db.First(&user, id).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (u *user) GetList() ([]entity.User, error) {
	users := []entity.User{}

	if err := u.db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (u *user) Update(user entity.User) (entity.User, error) {
	tx := u.db.Begin()
	defer tx.Rollback()

	if err := tx.Save(&user).Error; err != nil {
		return user, err
	}

	if err := tx.Commit().Error; err != nil {
		return user, err
	}

	return user, nil
}

func (u *user) Delete(id uint) error {
	tx := u.db.Begin()
	defer tx.Rollback()

	if err := tx.Delete(&entity.User{}, id).Error; err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
