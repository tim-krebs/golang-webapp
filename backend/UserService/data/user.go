package data

import (
	"time"
)

type User struct {
	ID         string    `bson:"_id"`
	Password   []byte    `bson:"password"`
	First_Name string    `bson:"first_name"`
	Last_Name  string    `bson:"last_name"`
	Birthday   time.Time `bson:"birthday"`
	Avatar     string    `bson:"avatar"`
	Email      string    `bson:"email"`
	TelNr      string    `bson:"telNr"`
	Company    string    `bson:"company"`
	Occupation string    `bson:"occupation"`
	School     string    `bson:"school"`
	Subject    string    `bson:"subject"`
	Country    string    `bson:"country"`
	IsActive   bool      `bson:"isActive"`
	Bio        string    `bson:"bio"`
	Role       []int     `bson:"role"`
}

type UserRequest struct {
	ID         string    `json:"id"`
	Password   string    `json:"password"`
	First_Name string    `json:"first_name"`
	Last_Name  string    `json:"last_name"`
	Birthday   time.Time `json:"birthday"`
	Avatar     string    `json:"avatar"`
	Email      string    `json:"email"`
	TelNr      string    `json:"telNr"`
	Company    string    `json:"company"`
	Occupation string    `json:"occupation"`
	School     string    `json:"school"`
	Subject    string    `json:"subject"`
	Country    string    `json:"country"`
	IsActive   bool      `bson:"isActive"`
	Bio        string    `bson:"bio"`
	Role       []int     `bson:"role"`
}

type UpdateUserRequest struct {
	Password   string    `json:"password"`
	First_Name string    `json:"first_name"`
	Last_Name  string    `json:"last_name"`
	Birthday   time.Time `json:"birthday"`
	Avatar     string    `json:"avatar"`
	Email      string    `json:"email"`
	TelNr      string    `json:"telNr"`
	Company    string    `json:"company"`
	Occupation string    `json:"occupation"`
	School     string    `json:"school"`
	Subject    string    `json:"subject"`
	Country    string    `json:"country"`
}
