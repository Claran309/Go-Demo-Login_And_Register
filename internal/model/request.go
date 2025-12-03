package model

// RegisterRequest "/register"
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

// LoginRequest "/login"
type LoginRequest struct {
	LoginKey string `json:"login_key" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RefreshTokenRequest "/refresh"
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// PickRequest "/pick"
type PickRequest struct {
	CourseID int `json:"course_id" binding:"required"`
}

// DropRequest "/drop"
type DropRequest struct {
	CourseID int `json:"course_id" binding:"required"`
}

// AddCourseRequest "/add/course"
type AddCourseRequest struct {
	Name    string `json:"name" binding:"required"`
	Capital int    `json:"capital" binding:"required"`
}

// CreateTodoRequest "/to-do/create"
type CreateTodoRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

// FinishTodoRequest "/to-do/finish"
type FinishTodoRequest struct {
	TodoID int `json:"todo_id" binding:"required"`
}

// DeleteTodoRequest "/to-do/delete"
type DeleteTodoRequest struct {
	TodoID int `json:"todo_id" binding:"required"`
}
