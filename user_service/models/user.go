package models

type User struct {
	ID       uint   `gorm:"primaryKey;auto_increment" db:"id"`
	Username string `gorm:"type:varchar(256);uniqueIndex;not_null" db:"username"`
	Password string `gorm:"type:varchar(256)" db:"password"`
}
type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserIdRequest struct {
	Id string `json:"id"`
}

type UsernameRequest struct {
	Username string `json:"username"`
}

type UserResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

type JWT struct {
	Token string `json:"token"`
}
