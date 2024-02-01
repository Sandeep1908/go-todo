package models

import "time"

type Todo struct {
    ID          string    `json:"id"`
    UserID      int       `json:"user_id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Status      string    `json:"status"`
    Created     time.Time `json:"created"`
    Updated     time.Time `json:"updated"`
}