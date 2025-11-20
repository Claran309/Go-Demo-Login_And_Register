package mysql

import (
	"GoGin/api/dao"
	"GoGin/internal/model"
	"errors"
	"log"

	"gorm.io/gorm"
)

type mysqlUserRepo struct {
	db *gorm.DB
}

func NewMysqlUserRepo(db *gorm.DB) dao.UserRepository {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("Failed to migrate user table:", err)
	}

	return &mysqlUserRepo{db: db}
}

func (repo *mysqlUserRepo) AddUser(user *model.User) error {
	//检查用户名是否存在
	var existsUsernameCount int64
	repo.db.Model(&model.User{}).
		Where("username = ?", user.Username).
		Count(&existsUsernameCount)
	if existsUsernameCount > 0 {
		return errors.New("user already exists")
	}

	//检查邮箱是否存在
	var existsEmailCount int64
	repo.db.Model(&model.User{}).
		Where("email = ?", user.Email).
		Count(&existsEmailCount)
	if existsEmailCount > 0 {
		return errors.New("email already exists")
	}

	err := repo.db.Create(user)
	if err.Error != nil {
		return err.Error
	}

	return nil
}

func (repo *mysqlUserRepo) SelectByUsername(username string) (*model.User, error) {
	var user model.User

	err := repo.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, errors.New("username select failed")
	}

	return &user, nil
}

func (repo *mysqlUserRepo) SelectByEmail(email string) (*model.User, error) {
	var user model.User

	err := repo.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, errors.New("email select failed")
	}

	return &user, nil
}

func (repo *mysqlUserRepo) Exists(username, email string) bool {
	var count int64
	repo.db.Where("username = ? AND email = ?", username, email).Count(&count)
	return count > 0
}

func (repo *mysqlUserRepo) GetRole(user *model.User) (string, error) {
	return user.Username, nil
}
