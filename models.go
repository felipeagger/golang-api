package main

import "time"

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username" binding:"required"`
	Age       int       `json:"age"`
	Bio       string    `json:"bio"`
	Link      string    `json:"link"`
	Avatar    string    `json:"avatar"`
	Score     float64   `json:"score"`
	CreatedAt time.Time `json:"created_at"`
}
