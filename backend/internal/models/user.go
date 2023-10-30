package models

type User struct {
	ID       int64  `json:"id" db:"id" validate:"omitempty" example:""`
	Username string `json:"username" db:"username" validate:"required" example:"Chat_User_1"`
	Email    string `json:"email" db:"email" validate:"email" example:"admin@example.com"`
	Password string `json:"password" db:"password" validate:"min=8" example:"password12345"`
}

type CreateUserReq struct {
	Username string `json:"username" db:"username" validate:"required" example:"Chat_User_1"`
	Email    string `json:"email" db:"email" validate:"email" example:"admin@example.com"`
	Password string `json:"password" db:"password" validate:"min=8" example:"password12345"`
}

type CreateUserRes struct {
	ID       string `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
}

type LoginUserReq struct {
	Email    string `json:"email" db:"email" validate:"email" example:"admin@example.com"`
	Password string `json:"password" db:"password" validate:"min=8" example:"password12345"`
}

type LoginUserRes struct {
	AccessToken string
	ID          string `json:"id" db:"id"`
	Username    string `json:"username" db:"username"`
}
