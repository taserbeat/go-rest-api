package models

import "time"

// ユーザーモデル
type User struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// ユーザー情報のレスポンス
type UserResponse struct {
	Id    uint   `json:"id" gorm:"primaryKey"`
	Email string `json:"email" gorm:"unique"`
}
