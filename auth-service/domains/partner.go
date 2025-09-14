package domains

import "time"

type Partner struct {
	ID        uint   `gorm:"primaryKey"`
	PartnerID string `gorm:"uniqueIndex;size:50"`
	Name      string `gorm:"size:100"`
	PublicKey string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
