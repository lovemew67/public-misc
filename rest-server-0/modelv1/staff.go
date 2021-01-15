package modelv1

import (
	"fmt"
	"time"
)

const (
	TableName = "staff"
)

type Staff struct {
	ID        int       `json:"id" gorm:"column:id;primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;index:updated_at"`
	Name      string    `json:"name,omitempty" gorm:"column:name"`
	Email     string    `json:"email,omitempty" gorm:"column:email"`
	AvatarUrl string    `json:"avatar_url,omitempty" gorm:"column:avatar_url"`
}

func (t *Staff) String() string {
	return fmt.Sprintf("%+v", *t)
}
