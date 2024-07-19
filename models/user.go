package models

import "go-auth-service/common"


type UserRole string

const (
	Admin    UserRole = "Admin"
	Employer UserRole = "User"
)


type User struct {
	common.BaseModel `bson:",inline"`
	Username 			 string   `json:"username" bson:"username"`
	Password 			 string   `json:"password" bson:"password"`
	Role     			 UserRole `json:"role" bson:"role"`
}

type UserAction string

const (
	Edit     UserAction = "Edit"
	Delete   UserAction = "Delete"
	Transfer UserAction = "Transfer"
)

type JWTUser struct {
	Username string `json:"username"`
	Role     UserRole `json:"role"`
}
