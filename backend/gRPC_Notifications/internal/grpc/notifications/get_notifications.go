package notifications

import (
	"encoding/json"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	v1 "notification_grpc/api"
	"notification_grpc/pkg/database"
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

	result := conn.Preload("NotificationText").Where(filters).Limit(int(request.Count)).Offset(int(request.Count * (request.Page - 1))).Find(&notifications)

	response := &v1.UserNotificationsResponse{}

	if result.Error != nil {
		response.Notifications = make([]*structpb.Struct, 0)
		return response, result.Error
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
