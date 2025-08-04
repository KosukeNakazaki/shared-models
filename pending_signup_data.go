package models

import (
	"time"
)

type PendingSignupData struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	User           User           `gorm:"embedded;embeddedPrefix:user_" json:"user"`
	CorporateGroup CorporateGroup `gorm:"embedded;embeddedPrefix:group_" json:"corporate_group"`
	Company        Company        `gorm:"embedded;embeddedPrefix:company_" json:"company"`
	Office         Office         `gorm:"embedded;embeddedPrefix:office_" json:"office"`
	Department     Department     `gorm:"embedded;embeddedPrefix:department_" json:"department"`
	Token          string         `gorm:"unique;not null" json:"token"`
	CreatedAt      time.Time      `json:"created_at" gorm:"autoCreateTime"`
}
