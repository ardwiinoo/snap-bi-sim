package domains

import "time"

type Token struct {
	ID        uint   `gorm:"primaryKey"`
	PartnerID string `gorm:"size:50;index"`
	Token     string `gorm:"type:text;index"`
	ExpiresAt time.Time
	CreatedAt time.Time
}
