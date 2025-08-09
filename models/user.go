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
	UUID          string     `json:"uuid" gorm:"type:uuid;primaryKey" validate:"required,uuid4"`
	UserType      string     `json:"user_type" gorm:"type:varchar(20)" validate:"required,oneof=root general sub_general"`
	Username      string     `json:"username" gorm:"type:varchar(30);index:idx_users_username,unique" validate:"required,alphanum,min=3,max=30"`
	Password      string     `json:"password" gorm:"type:text" validate:"required,min=8"`
	Email         string     `json:"email" gorm:"index:idx_users_email_nocase,unique,expression:lower(email)" validate:"required,email"`
	Lastname      string     `json:"lastname" gorm:"type:varchar(100)" validate:"required"`
	Firstname     string     `json:"firstname" gorm:"type:varchar(100)" validate:"required"`
	LastnameKana  string     `json:"lastname_kana" validate:"required,katakana"`
	FirstnameKana string     `json:"firstname_kana" validate:"required,katakana"`
	DepartmentID  string     `json:"department_id" gorm:"type:uuid;index"`
	Department    Department `gorm:"foreignKey:DepartmentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	Remarks       string     `json:"remarks,omitempty"`
	ManagedBy     string     `json:"managed_by,omitempty" gorm:"type:uuid;index"`
	Manager       *User      `gorm:"foreignKey:ManagedBy;references:UUID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Verified      bool       `json:"verified"`
	CreatedAt     time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}
