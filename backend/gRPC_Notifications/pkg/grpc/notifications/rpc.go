package notifications

import (
	"context"
	v1 "notification_grpc/api"
	"notification_grpc/pkg/database"
)

type NotificationService struct {
	v1.UnimplementedCreateNotificationsServer
}

func (s *NotificationService) CreateModels(ctx context.Context, req *v1.NotificationCreateRequest) (response *v1.NotificationCreateResponse, err error) {
	response = &v1.NotificationCreateResponse{
		IsCreated: false,
	}

	dbConn := database.NewConnection()
	err = dbConn.MakeConnection()

	if err != nil {
		return
	}

	err = dbConn.NewTransaction()

	if err != nil {
		return
	}

	response.IsCreated = true
	return response, nil
}
