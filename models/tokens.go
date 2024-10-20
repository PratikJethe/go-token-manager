package models

import (
    "time"
)

type Token struct {
    ID                 int       `json:"id"`
    Token              string    `json:"token"`
    LastActivationTime *time.Time `json:"last_activation_time"`
    IsDeleted          bool      `json:"is_deleted"`
    CreatedAt          time.Time `json:"created_at"`
    UpdatedAt          time.Time `json:"updated_at"`
}
