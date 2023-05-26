package grpc

import (
	"context"
	"fmt"
	v1 "notification_grpc/api"
)

type NotificationService struct {
	v1.UnimplementedCreateNotificationsServer
}

func (s *NotificationService) CreateModels(ctx context.Context, req *v1.NotificationCreateRequest) (*v1.NotificationCreateResponse, error) {

	fmt.Print(req.Text)
	response := &v1.NotificationCreateResponse{
		IsCreated: true,
	}

	return response, nil
}
