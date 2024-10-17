package schemas

import "time"

type UserSignup struct {
	Username    string `json:"username" db:"username"`
	EncodedPass string `json:"encoded_password" db:"password_hash"`
}

type UserLogin struct {
	Username    string `json:"username" db:"username"`
	EncodedPass string `json:"encoded_password" db:"password_hash"`
}

type UserToken struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Role     string `db:"role"`
}

type UserDetails struct {
	ID        int        `json:"id" db:"id"`
	Username  string     `json:"username" db:"username"`
	Role      string     `json:"role" db:"role"`
	CreatedAt *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type CurrentUser struct{}
