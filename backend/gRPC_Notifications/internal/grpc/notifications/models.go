package notifications

import "time"

type MassNotification struct {
	ID         uint64    `gorm:"column:id"`
	TargetId   uint64    `gorm:"column:target_id"`
	TargetType uint32    `gorm:"column:target_type"`
	Image      string    `gorm:"column:image"`
	Text       string    `gorm:"column:text"`
	Link       string    `gorm:"column:link"`
	Important  bool      `gorm:"column:important"`
	Type       int32     `gorm:"column:type"`
	Action     int32     `gorm:"column:action"`
	Date       time.Time `gorm:"column:date"`
}
