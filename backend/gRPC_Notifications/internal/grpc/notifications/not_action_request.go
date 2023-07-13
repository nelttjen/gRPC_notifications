package notifications

import (
	"context"
	"fmt"
	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson"
	"gorm.io/gorm"
	"log"
	v1 "notification_grpc/api"
	cfg "notification_grpc/internal/config"
	"notification_grpc/pkg/database"
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
		log.Printf("INFO: Action %s - No users to notify\n", request.Action)
		return true
	}
	if request.TargetId == nil || request.TargetType == nil {
		log.Printf("ERROR: RPC got invalid targets: TargetId: %v, TargetType%v \n", *request.TargetId, *request.TargetType)
		return false
	}

	title := &Title{}
	postgresConnection.Where(&Title{ID: *request.TargetId}).First(title)

	if title.ID == 0 {
		log.Printf("ERROR: Title not found: %d\n", *request.TargetId)
		return false
	}

	var items []*UserIds
	values := []interface{}{*request.TargetId, userIds}

	result := postgresConnection.Raw(
		"SELECT tb.user_id FROM title_bookmarks tb WHERE tb.title_id = ? AND tb.user_id IN ?", values...).Scan(&items)

	if result.Error != nil {
		log.Printf("ERROR: Could not find bookmarks for title %d\n: %v", *request.TargetId, result.Error)
		return false
	}

	if len(items) == 0 {
		log.Printf("INFO: Action %s - No users to notify\n", request.Action)
		return true
	}

	action, err := cfg.GetActionByKey(request.Action)

	if err != nil {
		log.Printf("ERROR: RPC got invalid action: %v\n", err)
		return false
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
		Action:     action,
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
		log.Printf("ERROR: Could not get DB transaction: %v\n", err)
		return false
	}

	if request.Text == nil {
		log.Printf("ERROR: RPC got invalid text for action %s: %s\n", request.Action, *request.Text)
		return false
	}

	intAction, err := cfg.GetActionByKey(request.Action)
	if err != nil {
		log.Printf("ERROR: RPC got invalid action: %v\n", err)
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
		Action:    intAction,
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
		log.Printf("ERROR: Connecting to mongo: %v\n", err)
		return nil, err
	}
	defer func(mongoConnection database.MongoClient) {
		err := mongoConnection.CloseConnection()
		if err != nil {
			log.Printf("ERROR: Closing mongo connection: %v\n", err)
		}
	}(mongoConnection)

	collection, err := mongoConnection.GetCollectionFromAdmin("user_notification_settings")

	if err != nil {
		log.Printf("ERROR: Getting collection: %v\n", err)
		return nil, err
	}

	cursor, err := collection.Find(context.TODO(), bson.D{
		{Key: fieldName, Value: true},
	})
	if err != nil {
		log.Printf("ERROR: Getting cursor: %v\n", err)
		return nil, err
	}

	var userIds []int

	for cursor.Next(context.Background()) {
		var config UserNotificationConfig
		err := cursor.Decode(&config)
		if err != nil {
			log.Printf("ERROR: Decoding cursor: %v\n", err)
			return nil, err
		}
		userIds = append(userIds, config.UserID)
	}
	return userIds, nil
}

func getDefaultsToNotifyPostgres(dbConnection database.PostgresConnection, settingsField string) (postgresConnection *gorm.DB, postgresTransaction *gorm.DB, userIds []int, err error) {
	postgresConnection, err = dbConnection.GetDBConnection()
	if err != nil {
		log.Printf("ERROR: Getting db connection: %v\n", err)
		return
	}

	postgresTransaction, err = dbConnection.GetDBTransaction()
	if err != nil {
		log.Printf("ERROR: Getting db transaction: %v\n", err)
		return
	}

	userIds, err = getListOfUsersToNotify(settingsField)
	if err != nil {
		return
	}

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
		log.Printf("ERROR: Sending notifications: %v\n", err)
		return false
	case <-countChan:
		counter++
		if counter == countOperations {
			break
		}
	}

	_ = dbConnection.CommitTransaction(true)

	if request.TargetId != nil {
		log.Printf("INFO: Added %d notifications to database for action: %s and target_id: %d\n", len(items), request.Action, *request.TargetId)
	} else {
		log.Printf("INFO: Added %d notifications to database for action: %s\n", len(items), request.Action)
	}
	return true
}

func chapterNotification(dbConnection database.PostgresConnection, request *v1.NotificationCreateRequest) bool {
	postgresConnection, postgresTransaction, userIds, err := getDefaultsToNotifyPostgres(dbConnection, "new_chapters")

	if err != nil {
		return false
	}

	if len(userIds) == 0 {
		log.Printf("INFO: Action %s - No users to notify\n", request.Action)
		return true
	}

	if request.TargetId == nil || request.TargetType == nil {
		log.Printf("ERROR: RPC got invalid targets: TargetId: %v, TargetType%v \n", *request.TargetId, *request.TargetType)
		return false
	}

	chapter := &Chapter{}
	chapterResult := postgresConnection.Preload("Title").Where(&Chapter{ID: *request.TargetId}).First(chapter)

	if chapter.ID == 0 {
		log.Printf("ERROR: Chapter not found: %d\n", *request.TargetId)
		return false
	}

	if chapterResult.Error != nil {
		log.Printf("ERROR: Could not find chapter %d\n: %v", *request.TargetId, chapterResult.Error)
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
		log.Printf("ERROR: Could not find bookmarks for chapter %d\n: %v", *request.TargetId, result.Error)
		return false
	}
	if len(items) == 0 {
		log.Printf("INFO: Action %s - No users to notify\n", request.Action)
		return true
	}

	intAction, err := cfg.GetActionByKey(request.Action)
	if err != nil {
		log.Printf("ERROR: RPC got invalid action: %v\n", err)
		return false
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
		Action:     intAction,
		Date:       time.Time{},
	}
	postgresTransaction.Create(newNotification)

	return createNotifications(dbConnection, items, newNotification, request)
}
