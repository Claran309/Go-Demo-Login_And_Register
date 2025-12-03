package model

// memory used
//
//	type User struct {
//		Username string `json:"username"`
//		Password string `json:"password"`
//		Email    string `json:"email"`
//		UserID   int    `json:"user_id"`
//	}

// User mysql-gorm
type User struct {
	UserID   int    `json:"user_id" gorm:"primary_key;AUTO_INCREMENT;column:user_id"`
	Username string `json:"username" gorm:"column:username;uniqueIndex;type:varchar(50)"`
	Email    string `json:"email" gorm:"column:email;uniqueIndex;type:varchar(100)"`
	Password string `json:"-" gorm:"column:password;type:varchar(255)"`
	Role     string `json:"role" gorm:"column:role;type:varchar(50)"`
}
