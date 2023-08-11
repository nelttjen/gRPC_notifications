package notifications

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/structpb"
	v1 "notification_grpc/api"
	"notification_grpc/internal/config"
	"notification_grpc/pkg/database"
	"notification_grpc/pkg/logger"
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
		logger.LogflnIfExists("error", "Error create notifications action: %v", logrus.ErrorLevel, config.LoggerLevelImportant, err)
		return response, nil
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
		logger.LogflnIfExists("error", "Error create notifications for user: %v", logrus.ErrorLevel, config.LoggerLevelImportant, err)

		return response, nil
	}

	response.IsCreated = ProcessUserRequest(dbConn, req)

	return
}
func (s *NotificationService) GetNotifications(ctx context.Context, req *v1.UserNotificationsRequest) (response *v1.UserNotificationsResponse, err error) {
	dbConn, err := PrepareTransaction()
	defaultResponse := &v1.UserNotificationsResponse{Notifications: make([]*structpb.Struct, 0)}

	if err != nil {
		logger.LogflnIfExists("error", "Failed to get user notifications for user_id %d: %v", logrus.ErrorLevel, config.LoggerLevelImportant, req.UserId, err)
		return defaultResponse, nil
	}

	response, err = GetUserNotifications(dbConn, req)
	if err != nil {
		logger.LogflnIfExists("error", "Failed to get user notifications for user_id %d: %v", logrus.ErrorLevel, config.LoggerLevelImportant, req.UserId, err)
		return defaultResponse, nil
	}
	return
}

func (s *NotificationService) GetMassNotifications(ctx context.Context, req *v1.UserMassNotificationRequest) (response *v1.UserMassNotificationResponse, err error) {
	dbConn, err := PrepareTransaction()
	defaultResponse := &v1.UserMassNotificationResponse{Notifications: make([]*structpb.Struct, 0)}
	if err != nil {
		logger.LogflnIfExists("error", "Failed to get user mass notifications for user_id %d: %v", logrus.ErrorLevel, config.LoggerLevelImportant, req.UserId, err)
		return defaultResponse, nil
	}

	response, err = GetUserMassNotifications(dbConn, req)
	if err != nil {
		logger.LogflnIfExists("error", "Failed to get user mass notifications for user_id %d: %v", logrus.ErrorLevel, config.LoggerLevelImportant, req.UserId, err)
		return defaultResponse, nil
	}

	return
}

func (s *NotificationService) ManageNotifications(ctx context.Context, req *v1.NotificationManageRequest) (response *v1.NotificationManageResponse, err error) {
	response = &v1.NotificationManageResponse{
		Success: false,
	}

	dbConn, err := PrepareTransaction()
	if err != nil {
		logger.LogflnIfExists("error", "Failed to manage notifications for user_id %d: %v", logrus.ErrorLevel, config.LoggerLevelImportant, req.UserId, err)
		return response, nil
	}

	err = ManageNotifications(dbConn, req)
	if err != nil {
		logger.LogfIfExists("error", "Failed to manage notifications for user_id %d: %v", logrus.ErrorLevel, config.LoggerLevelImportant, req.UserId, err)
		return response, nil
	}
	response.Success = true

	return
}
