package model

import "time"

type User struct {
	ID             int64     `json:"id"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"hashedPassword"`
	CreatedAt      time.Time `json:"createdAt"`
	IsSubscribed   bool      `json:"isSubscribed"`
}

type UserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
