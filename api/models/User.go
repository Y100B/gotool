package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"gotool/api/security"
	"html"
	"strings"
)

// User model
type User struct {
	gorm.Model
	NickName   string    `gorm:"size:20;not null;unique" json:"NickName"`
	Email      string    `gorm:"size:100;not null;unique" json:"Email"`
	UserAnswer string    `gorm:"size:100;not null;unique" json:"UserAnswer"`
	PassWord   string    `gorm:"size:60;not null" json:"PassWord,omitempty"`
}

// BeforeSave hash the user PassWord
func (u *User) BeforeSave() error {
	hashedPassWord, err := security.Hash(u.PassWord)
	if err != nil {
		return err
	}
	u.PassWord = string(hashedPassWord)
	return nil
}

// Prepare cleans the inputs
func (u *User) Prepare() {
	u.ID = 0
	u.NickName = html.EscapeString(strings.TrimSpace(u.NickName))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
}

// Validate validates the inputs
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.NickName == "" {
			return errors.New("NickName is required")
		}

		if u.Email == "" {
			return errors.New("Email is required")
		}

		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid email")
		}
	case "login":
		if u.Email == "" {
			return errors.New("Email is required")
		}

		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid email")
		}

		if u.PassWord == "" {
			return errors.New("PassWord is required")
		}
	default:
		if u.NickName == "" {
			return errors.New("NickName is required")
		}

		if u.PassWord == "" {
			return errors.New("PassWord is required")
		}

		if u.Email == "" {
			return errors.New("Email is required")
		}

		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid email")
		}
	}

	return nil
}
