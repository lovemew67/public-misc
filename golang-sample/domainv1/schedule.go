package domainv1

import (
	"time"
)

type Schedule struct {
	ID            string `json:"id" gorm:"column:id;primary_key"`
	TimeInHourUTC int    `json:"timeInHourUTC" gorm:"column:timeInHourUTC"`
	Enable        bool   `json:"enable" gorm:"column:enable"`

	Type JobType `json:"type" gorm:"column:type"`

	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at;index:updated_at"`
}
