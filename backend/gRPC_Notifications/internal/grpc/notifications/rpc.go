package notifications

import (
	"context"
	v1 "notification_grpc/api"
	"notification_grpc/pkg/database"
)

type NotificationService struct {
	v1.UnimplementedCreateNotificationsServer
}

func PrepareTransaction() (conn database.Connection, err error) {
	dbConn := database.NewConnection()
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

	response.IsCreated = true
	return response, nil
}

func (s *NotificationService) CreateNotificationsForUsers(ctx context.Context, req *v1.NotificationCreateRequest) (response *v1.NotificationCreateResponse, err error) {
	response = &v1.NotificationCreateResponse{
		IsCreated: false,
	}

	dbConn, err := PrepareTransaction()

	if err != nil {
		return
	}

	return
}
