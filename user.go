package main

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID         string    `gorm:"column:id;primary_key"`
	Username   string    `gorm:"column:username;uniqueIndex"`
	Password   string    `gorm:"column:password"`
	Email      string    `gorm:"column:email;uniqueIndex"`
	UserRoleID string    `gorm:"column:user_role_id"`
	PersonID   string    `gorm:"column:person_id"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	id := uuid.NewV4()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	tx.Statement.SetColumn("id", id.String())
	tx.Statement.SetColumn("password", string(hashedPassword))
	return nil
}