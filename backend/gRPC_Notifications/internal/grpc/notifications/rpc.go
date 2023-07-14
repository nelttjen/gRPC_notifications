package notifications

import (
	"context"
	"log"
	v1 "notification_grpc/api"
	"notification_grpc/pkg/database"
)

type NotificationService struct {
	v1.UnimplementedCreateNotificationsServer
}

func PrepareTransaction() (conn database.PostgresConnection, err error) {
	dbConn := database.NewPostgresConnection()
	err = dbConn.MakeConnection()

	if err != nil {
		return
	}

	err = dbConn.NewTransaction()

	if err != nil {
		return
	}
	return dbConn, nil
}

func (s *NotificationService) CreateNotificationsAction(ctx context.Context, req *v1.NotificationCreateRequest) (response *v1.NotificationCreateResponse, err error) {
	response = &v1.NotificationCreateResponse{
		IsCreated: false,
	}
	dbConn, err := PrepareTransaction()

	if err != nil {
		return
	}

	response.IsCreated = ProcessActionRequest(dbConn, req)
	return response, nil
}

func (s *NotificationService) CreateNotificationForUsers(ctx context.Context, req *v1.NotificationCreateManualRequest) (response *v1.NotificationCreateResponse, err error) {
	response = &v1.NotificationCreateResponse{
		IsCreated: false,
	}

	dbConn, err := PrepareTransaction()

	if err != nil {
		return
	}

	response.IsCreated = ProcessUserRequest(dbConn, req)

	return
}
func (s *NotificationService) GetNotifications(ctx context.Context, req *v1.UserNotificationsRequest) (response *v1.UserNotificationsResponse, err error) {
	dbConn, err := PrepareTransaction()
	if err != nil {
		return &v1.UserNotificationsResponse{}, err
	}

	response, err = GetUserNotifications(dbConn, req)
	if err != nil {
		log.Printf("Failed to get user notifications for user_id %d: %v", req.UserId, err)
	}
	return
}
