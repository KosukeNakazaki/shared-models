package models

import (
	"time"
)

type Office struct {
	ID        string    `json:"id" gorm:"primaryKey" validate:"required,uuid4"`       // 事業所ID（UUID）
	Name      string    `json:"name" gorm:"not null" validate:"required"`             // 事業所名（例：東京本社）
	CompanyID string    `json:"company_id" gorm:"not null" validate:"required,uuid4"` // 紐づく会社ID
	Remarks   string    `json:"remarks,omitempty"`                                    // 任意の備考
	Verified  bool      `json:"verified"`                                             // 認証済みかどうか
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`                     // 作成日時
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`                     // 更新日時
}
