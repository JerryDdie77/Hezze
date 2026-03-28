package domain

type User struct {
	ID           int    `json:"id"`
	UserName     string `json:"user_name"`
	FirstName    string `json:"first_name"`
	Surname      string `json:"surname"`
	PasswordHash string `json:"password_hash"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	IsBlocked    bool   `json:"is_blocked"`
}

type UserResponse struct {
	ID        int    `json:"id"`
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}
