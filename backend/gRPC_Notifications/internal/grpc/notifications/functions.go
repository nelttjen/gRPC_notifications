package notifications

import (
	"context"
	"fmt"
	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	v1 "notification_grpc/api"
	cfg "notification_grpc/internal/config"
	"notification_grpc/pkg/database"
	"time"
)

type UserBookmarks struct {
	UserId   int64
	Category int32
	Name     string
}

type UserNotificationConfig struct {
	UserID                  int  `bson:"user_id"`
	NewChapters             bool `bson:"new_chapters"`
	SpecialOffers           bool `bson:"special_offers"`
	CommentAnswer           bool `bson:"comment_answer"`
	AuthorPosts             bool `bson:"author_posts"`
	NewTitleStatus          bool `bson:"new_title_status"`
	NewAchievements         bool `bson:"new_achievements"`
	BattlepassNewLevel      bool `bson:"battlepass_new_level"`
	PersonalRecommendations bool `bson:"personal_recommendations"`
	NewMessages             bool `bson:"new_messages"`
}

func ProcessActionRequest(dbConnection database.PostgresConnection, request *v1.NotificationCreateRequest) bool {
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

func actionTitleNewName(dbConnection database.PostgresConnection, request *v1.NotificationCreateRequest) bool {
	postgresConnection, err := dbConnection.GetDBConnection()
	if err != nil {
		log.Printf("ERROR: getting db connection: %v\n", err)
		return false
	}

	postgresTransaction, err := dbConnection.GetDBTransaction()
	if err != nil {
		log.Printf("ERROR: getting db transaction: %v\n", err)
		return false
	}

	mongoConnection := database.NewMongoConnection()
	if err := mongoConnection.MakeConnection(); err != nil {
		log.Printf("ERROR: connecting to mongo: %v\n", err)
		return false
	}
	defer func(mongoConnection database.MongoClient) {
		err := mongoConnection.CloseConnection()
		if err != nil {
			log.Printf("ERROR: closing mongo connection: %v\n", err)
		}
	}(mongoConnection)

	collection, err := mongoConnection.GetCollectionFromAdmin("user_notification_settings")

	if err != nil {
		log.Printf("ERROR: getting collection: %v\n", err)
		return false
	}

	cursor, err := collection.Find(context.TODO(), bson.D{
		{Key: "new_title_status", Value: true},
	})
	if err != nil {
		log.Printf("ERROR: getting cursor: %v\n", err)
		return false
	}

	var userIds []int

	for cursor.Next(context.Background()) {
		var config UserNotificationConfig
		err := cursor.Decode(&config)
		if err != nil {
			log.Printf("ERROR: decoding cursor: %v\n", err)
			return false
		}
		userIds = append(userIds, config.UserID)
	}

	var items []UserBookmarks

	values := []interface{}{*request.TargetId, userIds}

	result := postgresConnection.Raw(
		"SELECT tb.user_id, tb.category, t.name FROM title_bookmarks tb LEFT OUTER JOIN titles t ON t.id = tb.title_id WHERE tb.title_id = ? AND tb.user_id IN ?", values...).Scan(&items)

	if result.Error != nil {
		log.Printf("ERROR: Could not find bookmarks for title %d\n: %v", *request.TargetId, result.Error)
		return false
	}

	action, err := cfg.GetActionByKey(request.Action)

	fmt.Printf("action: %d, request action: %s, target type: %d\n", action, request.Action, *request.TargetType)

	if err != nil {
		log.Printf("ERROR: RPC got invalid action: %v\n", err)
		return false
	}

	newNotification := &MassNotification{
		TargetId:   *request.TargetId,
		TargetType: *request.TargetType,
		Important:  request.Important,
		Type:       request.Type,
		Action:     action,
		Date:       time.Now(),
	}

	postgresTransaction.Create(newNotification)

	chunkList := funk.Chunk(items, cfg.DatabaseBatchSize)
	errChan := make(chan error)
	countChan := make(chan int)
	countOperations := 0
	counter := 0

	for _, chunk := range chunkList.([][]*UserBookmarks) {
		go massNotificationsAddToDatabase(dbConnection, chunk, int(newNotification.ID), countChan, errChan)
		countOperations++
	}

	select {
	case err := <-errChan:
		log.Printf("error sending notifications: %v\n", err)
		return false
	case <-countChan:
		counter++
		if counter == countOperations {
			break
		}
	}

	return true
}

func actionTitleNewChapter(dbConnection database.PostgresConnection, request *v1.NotificationCreateRequest) bool {
	return true
}

func actionSiteNotification(dbConnection database.PostgresConnection, request *v1.NotificationCreateRequest) bool {
	return true
}

func actionTitleChapterFree(dbConnection database.PostgresConnection, request *v1.NotificationCreateRequest) bool {
	return true
}

func massNotificationsAddToDatabase(connection database.PostgresConnection, chunk []*UserBookmarks, massNotId int, intChan chan int, errChan chan error) {

}
