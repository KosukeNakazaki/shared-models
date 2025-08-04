package models

import (
	"time"
)

type Company struct {
	ID               string    `json:"id" gorm:"primaryKey" validate:"required,uuid4"`            // 例：uuid
	Name             string    `json:"name" gorm:"not null" validate:"required"`                  // 例：三菱商事
	CorporateGroupID string    `json:"corporate_group_id" gorm:"index" validate:"required,uuid4"` // 紐づくグループ
	Remarks          string    `json:"remarks,omitempty"`                                         // 任意の備考
	Verified         bool      `json:"verified"`                                                  // メール認証後 true
	CreatedAt        time.Time `json:"created_at" gorm:"autoCreateTime"`                          // 自動作成時刻
	UpdatedAt        time.Time `json:"updated_at" gorm:"autoUpdateTime"`                          // 自動更新時刻
}
