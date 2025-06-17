package store

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Email    string `json:"email"`
	CreatedAt string `json:"created_at"`
}