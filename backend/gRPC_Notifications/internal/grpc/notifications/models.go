package notifications

import (
	"time"
)

type Notification struct {
	ID     uint      `gorm:"column:id"`
	UserId int       `gorm:"column:user_id"`
	Image  string    `gorm:"column:image"`
	Text   string    `gorm:"column:text"`
	Link   string    `gorm:"column:link"`
	Type   int       `gorm:"column:type"`
	Action int       `gorm:"column:action"`
	Status bool      `gorm:"column:status"`
	Date   time.Time `gorm:"column:date"`
}
