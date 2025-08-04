package models

import (
	"time"
)

type CorporateGroup struct {
	ID        string    `json:"id" gorm:"primaryKey" validate:"required,uuid4"` // 例：uuid
	Name      string    `json:"name" gorm:"not null" validate:"required"`       // 例：三菱グループ
	Remarks   string    `json:"remarks,omitempty"`                              // 任意の備考
	Verified  bool      `json:"verified"`                                       // メール認証済みか
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`               // 自動作成タイムスタンプ
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`               // 自動更新タイムスタンプ
}
