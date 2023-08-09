package notifications

import (
	"fmt"
	"log"
	v1 "notification_grpc/api"
	"notification_grpc/internal/config"
	"notification_grpc/pkg/database"
)

func ManageNotifications(connection database.PostgresConnection, request *v1.NotificationManageRequest) error {
	conn, _ := connection.GetDBConnection()

	var table string
	var sqlStr string
	var args []interface{}

	switch request.NotificationType {
	case config.NotificationTypeMass:
		table = "user_mass_notifications"
	case config.NotificationTypeUser:
		table = "user_notifications"
	default:
		return fmt.Errorf("unknown notification type: %s", request.NotificationType)
	}

	switch request.Action {
	case config.NotificationActionRead:
		sqlStr = fmt.Sprintf("UPDATE %s SET read = ? WHERE user_id = ? AND id IN ?", table)
		args = []interface{}{true, request.UserId, request.NotificationIds}
	case config.NotificationActionUnread:
		sqlStr = fmt.Sprintf("UPDATE %s SET read = ? WHERE user_id = ? AND id IN ?", table)
		args = []interface{}{false, request.UserId, request.NotificationIds}
	case config.NotificationActionDelete:
		sqlStr = fmt.Sprintf("DELETE FROM %s WHERE user_id = ? AND id IN ?", table)
		args = []interface{}{request.UserId, request.NotificationIds}
	default:
		log.Printf("[ERROR]: unknown notification action: %s", request.Action)
		return fmt.Errorf("invalid action: %s", request.Action)
	}

	result := conn.Exec(sqlStr, args...)
	if result.Error != nil {
		log.Printf("[ERROR]: failed to execute query: %s", sqlStr)
		return result.Error
	}

	return nil
}
