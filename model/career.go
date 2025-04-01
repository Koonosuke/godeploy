package model

import "time"

// Career represents a career entry.
type Career struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Period    string    `json:"period"`
    Content   string    `json:"content"`
    CreatedAt time.Time `json:"created_at"`
}
