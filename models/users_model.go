package models

import (
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Email     string `json:"email" binding:"max=64"`
	Password  string `json:"password" binding:"gte=8,max=64"`
	Firstname string `json:"firstname" binding:"max=64"`
	Lastname  string `json:"lastname" binding:"max=64"`
	RolesID   int    `json:"rolesId" gorm:"column:roles_id"`
}

type Roles struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Title string `json:"title" binding:"max=64"`
}

type LoginReq struct {
	Email    string `json:"email" binding:"max=64"`
	Password string `json:"password" binding:"gte=8,max=64"`
}

type LoginRes struct {
	Token   string `json:"token"`
	RolesID int    `json:"rolesId"`
}

type TokenClaims struct {
	jwt.RegisteredClaims
	UserID    int   `json:"userID"`
	ExpiresAt int64 `json:"exp"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes), err
}

func ComparePassword(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}
