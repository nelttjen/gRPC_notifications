package notifications

import (
	"log"
	v1 "notification_grpc/api"
	cfg "notification_grpc/internal/config"
	"notification_grpc/pkg/database"
)

type UserBookmarks struct {
	UserId   int64
	Category int32
	Name     string
}

func ProcessActionRequest(dbConnection database.Connection, request *v1.NotificationCreateRequest) bool {
	switch request.Action {
	case cfg.ActionTitleNewName:
		return actionTitleNewName(dbConnection, request)
	case cfg.ActionTitleNewChapter:
		return actionTitleNewChapter(dbConnection, request)
	case cfg.ActionSiteNotification:
		return actionSiteNotification(dbConnection, request)
	case cfg.ActionChapterFree:
		return actionTitleChapterFree(dbConnection, request)
	}
	return false
}

func actionTitleNewName(dbConnection database.Connection, request *v1.NotificationCreateRequest) bool {
	connection, err := dbConnection.GetDBConnection()
	if err != nil {
		log.Printf("error getting db connection: %v", err)
		return false
	}

	var items []UserBookmarks

	users := [1]int32{2}
	values := []interface{}{*request.TargetId, users}

	result := connection.Raw(
		"SELECT tb.user_id, tb.category, t.name  FROM title_bookmarks tb LEFT OUTER JOIN titles t ON t.id = tb.title_id WHERE tb.title_id = ? AND tb.user_id IN ?", values...).Scan(&items)
	if result.Error != nil {
		log.Printf("Could not find bookmarks for title %d\n: %v", *request.TargetId, result.Error)
		return false
	}

	for _, item := range items {
		log.Println(item.UserId)
	}
	return true
}

func actionTitleNewChapter(dbConnection database.Connection, request *v1.NotificationCreateRequest) bool {
	return true
}

func actionSiteNotification(dbConnection database.Connection, request *v1.NotificationCreateRequest) bool {
	return true
}

func actionTitleChapterFree(dbConnection database.Connection, request *v1.NotificationCreateRequest) bool {
	return true
}
