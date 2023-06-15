package models

import "time"

type Action struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}
