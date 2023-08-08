package notifications

import (
	"encoding/json"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	"gorm.io/gorm"
	v1 "notification_grpc/api"
	"notification_grpc/pkg/database"
	"strings"
)

func GetUserNotifications(dbConnection database.PostgresConnection, request *v1.UserNotificationsRequest) (*v1.UserNotificationsResponse, error) {
	conn, _ := dbConnection.GetDBConnection()

	var notifications []*UserNotification

	filters := UserNotification{
		UserId: uint64(request.UserId),
	}

	if request.Read != nil {
		filters.Read = *request.Read
	}
	if *request.OnlyImportant {
		filters.Important = true
	}

	result := conn.Joins("NotificationText").
		Where(filters).Limit(int(request.Count)).
		Offset(int(request.Count * (request.Page - 1))).
		Find(&notifications)

	response := &v1.UserNotificationsResponse{}
	response.Notifications = make([]*structpb.Struct, 0)

	if result.Error != nil {
		return response, result.Error
	}

	err := processNotificationTargets(notifications, conn)
	if err != nil {
		return response, err
	}

	var structs []*structpb.Struct
	for _, notification := range notifications {
		newStruct := &structpb.Struct{}
		jsn, _ := json.Marshal(notification)
		_ = protojson.Unmarshal(jsn, newStruct)
		structs = append(structs, newStruct)
	}

	response.Notifications = structs
	return response, nil
}

func GetUserMassNotifications(dbConnection database.PostgresConnection, request *v1.UserMassNotificationRequest) (*v1.UserMassNotificationResponse, error) {
	conn, _ := dbConnection.GetDBConnection()

	var notifications []*UserMassNotification
	var andClauses []string
	var andArgs []interface{}

	filters := &UserMassNotification{
		UserId: uint64(request.UserId),
	}

	if request.Read != nil {
		filters.Read = *request.Read
	}
	if *request.OnlyImportant {
		andClauses = append(andClauses, "\"Notification\".\"important\" = ?")
		andArgs = append(andArgs, true)
	}

	if request.Action != nil {
		andClauses = append(andClauses, "\"Notification\".\"action\" = ?")
		andArgs = append(andArgs, *request.Action)
	}

	andClauses = append(andClauses, "\"Notification\".\"type\" = ?")
	andArgs = append(andArgs, request.Type)

	result := conn.Joins("Notification").
		Where(filters).
		Where(strings.Join(andClauses, " AND "), andArgs...).
		Limit(int(request.Count)).Offset(int(request.Count * (request.Page - 1))).
		Find(&notifications)

	response := &v1.UserMassNotificationResponse{}
	response.Notifications = make([]*structpb.Struct, 0)

	if result.Error != nil {
		return response, result.Error
	}

	err := processMassNotificationTargets(notifications, conn)
	if err != nil {
		return response, err
	}

	var structs []*structpb.Struct
	for _, notification := range notifications {
		newStruct := &structpb.Struct{}
		jsn, _ := json.Marshal(notification)
		_ = protojson.Unmarshal(jsn, newStruct)
		structs = append(structs, newStruct)
	}

	response.Notifications = structs
	return response, nil
}

func processNotificationTargets(notifications []*UserNotification, conn *gorm.DB) error {
	queries := map[int][]int{}

	for _, notification := range notifications {
		_, ok := queries[int(notification.TargetType)]
		if !ok {
			queries[int(notification.TargetType)] = []int{}
		}
		queries[int(notification.TargetType)] = append(queries[int(notification.TargetType)], int(notification.TargetId))
	}

	results, err := loadResults(queries, conn)

	if err != nil {
		return err
	}

	for _, notification := range notifications {
		_, ok := results[int(notification.TargetType)][notification.TargetId]
		if !ok {
			continue
		}
		switch int(notification.TargetType) {
		case 1:
			notification.Target = results[1][notification.TargetId].(*Title)
		case 2:
			notification.Target = results[2][notification.TargetId].(*Chapter)
		case 3:
			notification.Target = results[3][notification.TargetId].(*Comment)
		case 4:
			notification.Target = results[4][notification.TargetId].(*Billing)
		case 5:
			notification.Target = results[5][notification.TargetId].(*SpecialOffer)
		case 6:
			notification.Target = results[6][notification.TargetId].(*Badge)
		}
	}

	return nil
}

func processMassNotificationTargets(notifications []*UserMassNotification, conn *gorm.DB) error {
	queries := map[int][]int{}

	for _, notification := range notifications {
		targetType := int(notification.Notification.TargetType)
		_, ok := queries[targetType]
		if !ok {
			queries[targetType] = []int{}
		}
		queries[targetType] = append(queries[targetType], int(notification.Notification.TargetId))
	}

	results, err := loadResults(queries, conn)
	if err != nil {
		return err
	}

	for _, notification := range notifications {
		_, ok := results[int(notification.Notification.TargetType)][notification.Notification.TargetId]
		if !ok {
			continue
		}
		switch int(notification.Notification.TargetType) {
		case 1:
			notification.Notification.Target = results[1][notification.Notification.TargetId].(*Title)
		case 2:
			notification.Notification.Target = results[2][notification.Notification.TargetId].(*Chapter)
		case 3:
			notification.Notification.Target = results[3][notification.Notification.TargetId].(*Comment)
		case 4:
			notification.Notification.Target = results[4][notification.Notification.TargetId].(*Billing)
		case 5:
			notification.Notification.Target = results[5][notification.Notification.TargetId].(*SpecialOffer)
		case 6:
			notification.Notification.Target = results[6][notification.Notification.TargetId].(*Badge)
		}
	}
	return nil
}

func loadResults(queries map[int][]int, conn *gorm.DB) (map[int]map[uint64]interface{}, error) {
	results := map[int]map[uint64]interface{}{}

	for i := 1; i < 7; i++ {
		results[i] = make(map[uint64]interface{})
		ids, ok := queries[i]
		if ok {
			switch i {
			case 1:
				var titles []*Title
				result := conn.Where("id in (?)", ids).Find(&titles)
				if result.Error != nil {
					return nil, result.Error
				}
				for _, title := range titles {
					results[1][title.ID] = title
				}
			case 2:
				var chapters []*Chapter
				result := conn.Joins("Title").Where("\"chapters\".\"id\" in (?)", ids).Find(&chapters)
				if result.Error != nil {
					return nil, result.Error
				}
				for _, chapter := range chapters {
					results[2][chapter.ID] = chapter
				}
			case 3:
				var comments []*Comment
				result := conn.Preload("ReplyTo").Where("id in (?)", ids).Find(&comments)
				if result.Error != nil {
					return nil, result.Error
				}
				for _, comment := range comments {
					results[3][comment.ID] = comment
				}
			case 4:
				var billings []*Billing
				result := conn.Where("id in (?)", ids).Find(&billings)
				if result.Error != nil {
					return nil, result.Error
				}
				for _, billing := range billings {
					results[4][billing.ID] = billing
				}
			case 5:
				var specialOffers []*SpecialOffer
				result := conn.Where("id in (?)", ids).Find(&specialOffers)
				if result.Error != nil {
					return nil, result.Error
				}
				for _, specialOffer := range specialOffers {
					results[5][specialOffer.ID] = specialOffer
				}
			case 6:
				var badges []*Badge
				result := conn.Where("id in (?)", ids).Find(&badges)
				if result.Error != nil {
					return nil, result.Error
				}
				for _, badge := range badges {
					results[6][badge.ID] = badge
				}
			}

		}
	}
	return results, nil
}
