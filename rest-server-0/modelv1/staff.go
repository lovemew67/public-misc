package modelv1

import (
	"fmt"
	"time"
)

const (
	TableName = "staff"
)

type Staff struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ID        string    `json:"id"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	AvatarUrl string    `json:"avatar_url,omitempty"`
}

func (t *Staff) String() string {
	return fmt.Sprintf("%+v", *t)
}
