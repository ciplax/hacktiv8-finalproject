package users

//User store user
type User struct {
	UserID    string `json:"user_id"  db:"user_id"`
	Username  string `json:"username" db:"username"`
	Name      string `json:"name" db:"name"`
	Password  string `json:"password" db:"password"`
	LastLogin string `json:"last_login" db:"last_login"`
	BirthDate string `json:"birth_date" db:"birth_date"`
	Address   string `json:"address" db:"address"`
	Gender    int    `json:"gender" db:"gender"`
	RoleID    int    `json:"role_id" db:"role_id"`
}
