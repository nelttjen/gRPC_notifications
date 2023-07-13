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

func (m *MassNotification) TableName() string {
	return "mass_notifications"
}

type UserNotification struct {
	ID               uint64           `gorm:"column:id"`
	UserId           uint64           `gorm:"column:user_id"`
	TargetId         uint64           `gorm:"column:target_id"`
	TargetType       uint32           `gorm:"column:target_type"`
	Image            string           `gorm:"column:image"`
	Text             string           `gorm:"column:text"`
	TextId           *uint64          `gorm:"column:text_id"`
	NotificationText NotificationText `gorm:"foreignkey:TextId"`
	Link             string           `gorm:"column:link"`
	Important        bool             `gorm:"column:important"`
	Confirmation     bool             `gorm:"column:confirmation"`
	Read             bool             `gorm:"column:read"`
	Date             time.Time        `gorm:"column:date"`
}

func (u *UserNotification) TableName() string {
	return "user_notifications"
}

type UserMassNotification struct {
	ID             uint64           `gorm:"column:id"`
	UserId         uint64           `gorm:"column:user_id"`
	NotificationId *uint64          `gorm:"column:notification_id"`
	Notification   MassNotification `gorm:"foreignkey:NotificationId"`
	Read           bool             `gorm:"column:read"`
}

func (m UserMassNotification) TableName() string {
	return "user_mass_notifications"
}

type NotificationText struct {
	ID   uint64 `gorm:"column:id"`
	Text string `gorm:"column:text"`
}

func (NotificationText) TableName() string {
	return "notification_texts"
}

type Title struct {
	ID   uint64 `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

func (t Title) TableName() string {
	return "titles"
}

type Chapter struct {
	ID      uint64  `gorm:"column:id"`
	TitleId *uint64 `gorm:"column:title_id"`
	Title   Title   `gorm:"foreignkey:TitleId"`
	IsPaid  bool    `gorm:"column:is_paid"`
	Index   int32   `gorm:"column:index"`
}

func (c Chapter) TableName() string {
	return "chapters"
}

type Billing struct {
	ID     uint64  `gorm:"column:id"`
	UserId uint64  `gorm:"column:user_id"`
	Sum    float64 `gorm:"column:sum"`
}

func (b Billing) TableName() string {
	return "billings"
}

type SpecialOffer struct {
	ID      uint64  `gorm:"column:id"`
	UserId  uint64  `gorm:"column:user_id"`
	NeedSum float64 `gorm:"column:need_sum"`
	Reward  float64 `gorm:"column:reward"`
}

func (s SpecialOffer) TableName() string {
	return "special_offers"
}

type Badge struct {
	ID     uint64 `gorm:"column:id"`
	UserId uint64 `gorm:"column:user_id"`
	Name   string `gorm:"column:name"`
}

func (b Badge) TableName() string {
	return "badges"
}

type TitleBookmark struct {
	ID       uint64 `gorm:"column:id"`
	UserId   uint64 `gorm:"column:user_id"`
	TitleId  uint64 `gorm:"column:title_id"`
	Category int32  `gorm:"column:category"`
}

func (t TitleBookmark) TableName() string {
	return "title_bookmarks"
}
