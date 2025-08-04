package models

import (
	"time"
)

type Department struct {
	ID        string    `json:"id" gorm:"primaryKey" validate:"required,uuid4"`      // 部署の一意ID（UUID）
	Name      string    `json:"name" gorm:"not null" validate:"required"`            // 部署名（例：営業部）
	OfficeID  string    `json:"office_id" gorm:"not null" validate:"required,uuid4"` // 所属営業所のUUID
	Remarks   string    `json:"remarks,omitempty"`                                   // 任意の備考
	Verified  bool      `json:"verified"`                                            // メール認証等の状態
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`                    // 作成日時
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`                    // 更新日時
}
