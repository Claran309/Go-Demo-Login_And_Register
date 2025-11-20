package memory

// ================================未适配mysql前采用的内存存储数据==============================

//
//import (
//	"GoGin/internal/model"
//	"GoGin/internal/dao"
//	"errors"
//	"sync"
//)
//
//type memoryUserRepository struct {
//	database    map[string]*model.User // 以username存储用户信息
//	emailToData map[string]string      // 存储email对应username
//	DataSync    sync.Mutex             // 并发安全
//}
//
//func NewMemoryUserRepository() dao.UserRepository {
//	return &memoryUserRepository{
//		database:    make(map[string]*model.User),
//		emailToData: make(map[string]string),
//	}
//}
//
//// AddUser 添加新用户
//func (repo *memoryUserRepository) AddUser(user *model.User) error {
//	repo.DataSync.Lock()
//	defer repo.DataSync.Unlock()
//
//	if _, ok := repo.database[user.Username]; ok {
//		return errors.New("user already exists")
//	}
//
//	if _, ok := repo.emailToData[user.Email]; ok {
//		return errors.New("email already exists")
//	}
//
//	repo.database[user.Username] = user
//	repo.emailToData[user.Email] = user.Username
//
//	return nil
//}
//
//// SelectByUsername 登录时：返回正确密码供检查
//func (repo *memoryUserRepository) SelectByUsername(username string) (*model.User, error) {
//	repo.DataSync.Lock()
//	defer repo.DataSync.Unlock()
//
//	user, ok := repo.database[username]
//	if !ok {
//		return nil, errors.New("user not found")
//	}
//
//	return user, nil
//}
//
//// SelectByEmail 登录时：返回正确密码供检查
//func (repo *memoryUserRepository) SelectByEmail(email string) (*model.User, error) {
//	repo.DataSync.Lock()
//	defer repo.DataSync.Unlock()
//
//	username, ok := repo.emailToData[email]
//	if !ok {
//		return nil, errors.New("email not found")
//	}
//
//	user, ok := repo.database[username]
//	if !ok {
//		return nil, errors.New("user not found")
//	}
//
//	return user, nil
//}
//
//func (repo *memoryUserRepository) Exists(username, email string) bool {
//	repo.DataSync.Lock()
//	defer repo.DataSync.Unlock()
//
//	_, userExist := repo.database[username]
//	_, emailExist := repo.emailToData[email]
//
//	return userExist || emailExist
//}
