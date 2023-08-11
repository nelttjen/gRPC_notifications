package notifications

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson"
	"gorm.io/gorm"
	v1 "notification_grpc/api"
	cfg "notification_grpc/internal/config"
	"notification_grpc/pkg/database"
	"notification_grpc/pkg/logger"
	"strings"
	"time"
)

type UserIds struct {
	UserId int64
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
	postgresConnection, postgresTransaction, userIds, err := getDefaultsToNotifyPostgres(dbConnection, "title_new_status")

	if err != nil {
		return false
	}

	if len(userIds) == 0 {
		logger.LogflnIfExists("debug", "Action %s - No users to notify", logrus.DebugLevel, cfg.LoggerLevelAll, request.Action)
		return true
	}
	if request.TargetId == nil || request.TargetType == nil {
		logger.LogflnIfExists("error", "RPC got invalid targets: TargetId: %v, TargetType%v \n", logrus.ErrorLevel, cfg.LoggerLevelImportant, *request.TargetId, *request.TargetType)
		return false
	}

	title := &Title{}
	postgresConnection.Where(&Title{ID: *request.TargetId}).First(title)

	if title.ID == 0 {
		logger.LogflnIfExists("error", "Title not found: %d\n", logrus.ErrorLevel, cfg.LoggerLevelImportant, *request.TargetId)
		return false
	}

	var items []*UserIds
	values := []interface{}{*request.TargetId, userIds}

	result := postgresConnection.Raw(
		"SELECT tb.user_id FROM title_bookmarks tb WHERE tb.title_id = ? AND tb.user_id IN ?", values...).Scan(&items)

	if result.Error != nil {
		logger.LogflnIfExists("error", "Could not find bookmarks for title %d: %v", logrus.ErrorLevel, cfg.LoggerLevelImportant, *request.TargetId, result.Error)
		return false
	}

	if len(items) == 0 {
		logger.LogflnIfExists("debug", "Action %s - No users to notify", logrus.DebugLevel, cfg.LoggerLevelAll, request.Action)
		return true
	}

	prevName := "Неизвестно"
	if *request.Text != "" {
		prevName = *request.Text
	}

	newNotification := &MassNotification{
		TargetId:   *request.TargetId,
		TargetType: *request.TargetType,
		Important:  request.Important,
		Type:       request.Type,
		Action:     request.Action,
		Text:       fmt.Sprintf("У тайтла в ваших закладках изменилось название. Прежнее название: %s, Новое название: %s", prevName, title.Name),
		Link:       fmt.Sprintf("/titles/%d", *request.TargetId),
		Date:       time.Now(),
	}
	postgresTransaction.Create(newNotification)

	return createNotifications(dbConnection, items, newNotification, request)
}

func actionTitleNewChapter(dbConnection database.PostgresConnection, request *v1.NotificationCreateRequest) bool {
	return chapterNotification(dbConnection, request)
}

func actionSiteNotification(dbConnection database.PostgresConnection, request *v1.NotificationCreateRequest) bool {
	transaction, err := dbConnection.GetDBTransaction()
	if err != nil {
		logger.LogflnIfExists("error", "Could not get DB transaction: %v\n", logrus.ErrorLevel, cfg.LoggerLevelImportant, err)
		return false
	}

	if request.Text == nil {
		logger.LogflnIfExists("error", "RPC got invalid text for action %s: %s\n", logrus.ErrorLevel, cfg.LoggerLevelImportant, request.Action, *request.Text)
		return false
	}

	var userIds []*UserIds
	transaction.Raw("SELECT a.id as user_id FROM auth_user a").Scan(&userIds)

	newNotification := &MassNotification{
		Text:      *request.Text,
		Link:      *request.Link,
		Image:     *request.Image,
		Important: request.Important,
		Type:      request.Type,
		Action:    request.Action,
		Date:      time.Now(),
	}

	c, _ := dbConnection.GetDBConnection()
	c.Create(newNotification)

	return createNotifications(dbConnection, userIds, newNotification, request)
}

func actionTitleChapterFree(dbConnection database.PostgresConnection, request *v1.NotificationCreateRequest) bool {
	return chapterNotification(dbConnection, request)
}

func massNotificationsAddToDatabase(connection database.PostgresConnection, chunk []*UserIds, massNotId int, intChan chan int, errChan chan error) {
	var strs []string
	var valueArgs []interface{}

	for _, bookmark := range chunk {
		strs = append(strs, "(?, ?, ?)")
		valueArgs = append(valueArgs, false)
		valueArgs = append(valueArgs, massNotId)
		valueArgs = append(valueArgs, bookmark.UserId)
	}

	err := connection.ExecTransaction(fmt.Sprintf("INSERT INTO user_mass_notifications(read, notification_id, user_id) VALUES %s", strings.Join(strs, ",")), valueArgs)
	if err != nil {
		errChan <- err
		return
	}
	intChan <- 1
}

func getListOfUsersToNotify(fieldName string) ([]int, error) {
	mongoConnection := database.NewMongoConnection()
	if err := mongoConnection.MakeConnection(); err != nil {
		logger.LogflnIfExists("error", "Could not connect to mongo: %v", logrus.ErrorLevel, cfg.LoggerLevelImportant, err)
		return nil, err
	}
	defer func(mongoConnection database.MongoClient) {
		err := mongoConnection.CloseConnection()
		if err != nil {
			logger.LogflnIfExists("error", "Could not close mongo connection: %v", logrus.ErrorLevel, cfg.LoggerLevelImportant, err)
		}
	}(mongoConnection)

	collection, err := mongoConnection.GetCollectionFromAdmin("user_notification_settings")

	if err != nil {
		logger.LogflnIfExists("error", "Could not get collection: %v", logrus.ErrorLevel, cfg.LoggerLevelImportant, err)
		return nil, err
	}

	cursor, err := collection.Find(context.TODO(), bson.D{
		{Key: fieldName, Value: true},
	})
	if err != nil {
		logger.LogflnIfExists("error", "Could not get cursor: %v", logrus.ErrorLevel, cfg.LoggerLevelImportant, err)
		return nil, err
	}

	var userIds []int

	for cursor.Next(context.Background()) {
		var config UserNotificationConfig
		err := cursor.Decode(&config)
		if err != nil {
			logger.LogflnIfExists("error", "Could not decode cursor: %v", logrus.ErrorLevel, cfg.LoggerLevelImportant, err)
			return nil, err
		}
		userIds = append(userIds, config.UserID)
	}
	return userIds, nil
}

func getDefaultsToNotifyPostgres(dbConnection database.PostgresConnection, settingsField string) (postgresConnection *gorm.DB, postgresTransaction *gorm.DB, userIds []int, err error) {
	postgresConnection, err = dbConnection.GetDBConnection()
	if err != nil {
		logger.LogflnIfExists("error", "Could not get db connection: %v", logrus.ErrorLevel, cfg.LoggerLevelImportant, err)
		return
	}

	postgresTransaction, err = dbConnection.GetDBTransaction()
	if err != nil {
		logger.LogflnIfExists("error", "Could not get db transaction: %v", logrus.ErrorLevel, cfg.LoggerLevelImportant, err)
		return
	}

	userIds, err = getListOfUsersToNotify(settingsField)
	return
}

func createNotifications(dbConnection database.PostgresConnection, items []*UserIds, newNotification *MassNotification, request *v1.NotificationCreateRequest) bool {
	chunkList := funk.Chunk(items, cfg.DatabaseBatchSize)
	errChan := make(chan error)
	countChan := make(chan int)
	countOperations := 0
	counter := 0

	for _, chunk := range chunkList.([][]*UserIds) {
		go massNotificationsAddToDatabase(dbConnection, chunk, int(newNotification.ID), countChan, errChan)
		countOperations++
	}

	select {
	case err := <-errChan:
		logger.LogflnIfExists("error", "Could not send notifications: %v", logrus.ErrorLevel, cfg.LoggerLevelImportant, err)
		return false
	case <-countChan:
		counter++
		if counter == countOperations {
			break
		}
	}

	_ = dbConnection.CommitTransaction(true)

	if request.TargetId != nil {
		logger.LogflnIfExists("debug", "Added %d notifications to database for action: %s and target_id: %d", logrus.DebugLevel, cfg.LoggerLevelAll, len(items), request.Action, *request.TargetId)
	} else {
		logger.LogflnIfExists("debug", "Added %d notifications to database for action: %s", logrus.DebugLevel, cfg.LoggerLevelAll, len(items), request.Action)
	}
	return true
}

func chapterNotification(dbConnection database.PostgresConnection, request *v1.NotificationCreateRequest) bool {
	postgresConnection, postgresTransaction, userIds, err := getDefaultsToNotifyPostgres(dbConnection, "new_chapters")

	if err != nil {
		return false
	}

	if len(userIds) == 0 {
		logger.LogflnIfExists("debug", "Action %s - No users to notify", logrus.DebugLevel, cfg.LoggerLevelAll, request.Action)
		return true
	}

	if request.TargetId == nil || request.TargetType == nil {
		logger.LogflnIfExists("error", "RPC got invalid targets: TargetId: %v, TargetType: %v", logrus.ErrorLevel, cfg.LoggerLevelImportant, *request.TargetId, *request.TargetType)
		return false
	}

	chapter := &Chapter{}
	chapterResult := postgresConnection.Joins("Title").Where(&Chapter{ID: *request.TargetId}).First(chapter)

	if chapter.ID == 0 {
		logger.LogflnIfExists("error", "Could not find chapter with id: %d", logrus.ErrorLevel, cfg.LoggerLevelImportant, *request.TargetId)
		return false
	}

	if chapterResult.Error != nil {
		logger.LogflnIfExists("error", "Could not find chapter with id%d - %v", logrus.ErrorLevel, cfg.LoggerLevelImportant, *request.TargetId, chapterResult.Error)
		return false
	}

	var items []*UserIds
	values := []interface{}{*request.TargetId, userIds}
	result := postgresConnection.Raw(`
	 SELECT tb.user_id FROM title_bookmarks tb 
     LEFT OUTER JOIN titles t ON t.id = tb.title_id 
	 LEFT OUTER JOIN chapters c ON c.title_id = t.id 
	 WHERE c.id = ? AND tb.user_id IN ?`, values...).Scan(&items)

	if result.Error != nil {
		logger.LogflnIfExists("error", "Could not find bookmarks for chapter %d: %v", logrus.ErrorLevel, cfg.LoggerLevelImportant, *request.TargetId, result.Error)
		return false
	}
	if len(items) == 0 {
		logger.LogflnIfExists("debug", "Action %s - No users to notify", logrus.DebugLevel, cfg.LoggerLevelAll, request.Action)
		return true
	}

	var text string

	if request.Action == cfg.ActionChapterFree {
		text = fmt.Sprintf("\"%s\" - Глава %d вышла в бесплатном доступе", chapter.Title.Name, chapter.Index)
	} else {
		access := "бесплатном"
		if chapter.IsPaid {
			access = "платном"
		}
		text = fmt.Sprintf("\"%s\" - Добавлена глава %d в %s доступе", chapter.Title.Name, chapter.Index, access)
	}

	newNotification := &MassNotification{
		TargetId:   *request.TargetId,
		TargetType: *request.TargetType,
		Text:       text,
		Link:       fmt.Sprintf("/titles/%d/chapters/%d", *chapter.TitleId, chapter.ID),
		Important:  request.Important,
		Type:       request.Type,
		Action:     request.Action,
		Date:       time.Time{},
	}
	postgresTransaction.Create(newNotification)

	return createNotifications(dbConnection, items, newNotification, request)
}
