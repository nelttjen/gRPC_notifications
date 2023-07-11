package notifications

import (
	v1 "notification_grpc/api"
	cfg "notification_grpc/internal/config"
	"notification_grpc/pkg/database"
)

func ProcessActionRequest(dbConnection *database.Connection, request *v1.NotificationCreateRequest) {
	switch request.Action {
	case cfg.ACTION_USER_CREATED:
		text := "Добро пожаловать на сайт!"
		user := database.
	}
}

func actionTitleNewChapterPaid(dbConnection *database.Connection, request *v1.NotificationCreateRequest) {

}
