package models

import (
	"time"
)

// ユーザー種別
type UserType string

const (
	RootUser       UserType = "root"
	GeneralUser    UserType = "general"
	SubGeneralUser UserType = "sub_general"
)

type User struct {
	UUID          string    `json:"uuid" gorm:"primaryKey" validate:"required,uuid4"`
	UserType      UserType  `json:"user_type" gorm:"type:varchar(20)" validate:"required,oneof=root general sub_general"`
	Username      string    `json:"username" validate:"required,alphanum,min=3,max=30" gorm:"uniqueIndex"`
	Password      string    `json:"password" validate:"required,min=8" gorm:"type:text"`
	Email         string    `json:"email" validate:"required,email" gorm:"uniqueIndex"`
	Lastname      string    `json:"lastname" validate:"required"`
	Firstname     string    `json:"firstname" validate:"required"`
	LastnameKana  string    `json:"lastname_kana" validate:"required,katakana"`
	FirstnameKana string    `json:"firstname_kana" validate:"required,katakana"`
	DepartmentID  string    `json:"department_id" gorm:"index"`
	Remarks       string    `json:"remarks,omitempty"`
	ManagedBy     string    `json:"managed_by,omitempty" gorm:"index"`
	Verified      bool      `json:"verified"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
