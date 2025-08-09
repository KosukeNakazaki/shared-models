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

// User はDBの users テーブルに対応
// - UUID/ManagedBy/DepartmentID は DB 側の型を uuid に揃える前提
// - 関連(Department, Manager)は持たせず、IDだけにして依存を最小化
type User struct {
	UUID     string   `json:"uuid" gorm:"type:uuid;primaryKey" validate:"required,uuid4"`
	UserType UserType `json:"user_type" gorm:"type:varchar(20)" validate:"required,oneof=root general sub_general"`

	Username string `json:"username" gorm:"type:varchar(30);uniqueIndex" validate:"required,alphanum,min=3,max=30"`
	Password string `json:"password" gorm:"type:text" validate:"required,min=8"`
	Email    string `json:"email" gorm:"uniqueIndex" validate:"required,email"`

	Lastname      string `json:"lastname" gorm:"type:varchar(100)" validate:"required"`
	Firstname     string `json:"firstname" gorm:"type:varchar(100)" validate:"required"`
	LastnameKana  string `json:"lastname_kana" validate:"required,katakana"`
	FirstnameKana string `json:"firstname_kana" validate:"required,katakana"`

	// 外部参照は“IDだけ”にして依存を減らす（NULL許容が運用ラク）
	DepartmentID *string `json:"department_id" gorm:"type:uuid;index"`
	ManagedBy    *string `json:"managed_by,omitempty" gorm:"type:uuid;index"`

	Remarks string `json:"remarks,omitempty"`

	Verified  bool      `json:"verified"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
