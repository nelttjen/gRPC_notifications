package notifications

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	v1 "notification_grpc/api"
	cfg "notification_grpc/internal/config"
	"notification_grpc/pkg/database"
	"notification_grpc/pkg/logger"
	"time"
)

func ProcessUserRequest(dbConnection database.PostgresConnection, request *v1.NotificationCreateManualRequest) bool {
	tx, err := dbConnection.GetDBTransaction()
	if err != nil {
		logger.LogflnIfExists("error", "Could not get DB transaction: %v", logrus.ErrorLevel, cfg.LoggerLevelImportant, err)
		return false
	}

	if request.UserIds == nil || len(request.UserIds) == 0 {
		logger.LoglnIfExists("info", "get empty list of user IDs or text is empty, request ignored", logrus.WarnLevel, cfg.LoggerLevelImportant)
		return true
	}

	mongoClient := database.NewMongoConnection()
	err = mongoClient.MakeConnection()
	if err != nil {
		logger.LogflnIfExists("error", "Could not connect to MongoDB: %v", logrus.ErrorLevel, cfg.LoggerLevelImportant, err)
		return false
	}

	defer func(cl database.MongoClient) {
		err := cl.CloseConnection()
		if err != nil {
			logger.LogflnIfExists("error", "Could not close MongoDB connection: %v", logrus.ErrorLevel, cfg.LoggerLevelImportant, err)
		}
	}(mongoClient)

	collection, err := mongoClient.GetCollectionFromAdmin("user_notification_settings")
	if err != nil {
		logger.LogflnIfExists("error", "Could not get MongoDB collection user_notification_settings: %v", logrus.ErrorLevel, cfg.LoggerLevelImportant, err)
		return false
	}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	filter := bson.M{
		"user_id":           bson.M{"$in": request.UserIds},
		request.SettingsKey: true,
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		logger.LogflnIfExists("error", "Could not get MongoDB cursor: %v", logrus.ErrorLevel, cfg.LoggerLevelImportant, err)
		return false
	}
	defer func(c *mongo.Cursor, ctx context.Context) {
		err := c.Close(ctx)
		if err != nil {
			logger.LogflnIfExists("error", "Could not close MongoDB cursor: %v", logrus.ErrorLevel, cfg.LoggerLevelImportant, err)
		}
	}(cursor, ctx)

	var notUserIds []int
	for cursor.Next(ctx) {
		user := &UserNotificationConfig{}
		err := cursor.Decode(user)
		if err != nil {
			logger.LogflnIfExists("error", "Could not decode MongoDB cursor: %v", logrus.ErrorLevel, cfg.LoggerLevelImportant, err)
			return false
		}
		notUserIds = append(notUserIds, user.UserID)
	}

	var userNots []*UserNotification

	useTargets := request.TargetId != nil && request.TargetType != nil
	useText, err := getTextModelOrCreate(tx, request.Text, request)

	if err != nil {
		return false
	}

	for _, uid := range notUserIds {
		newNotification := &UserNotification{
			UserId:       uint64(uid),
			Important:    request.Important,
			Confirmation: request.Confirmation,
			Read:         false,
			Date:         time.Now(),
		}

		if useTargets {
			newNotification.TargetId = *request.TargetId
			newNotification.TargetType = *request.TargetType
		}

		if request.Image != nil {
			newNotification.Image = *request.Image
		}

		if request.Link != nil {
			newNotification.Link = *request.Link
		}

		if useText != nil {
			newNotification.TextId = &useText.ID
		} else {
			newNotification.Text = request.Text
		}

		userNots = append(userNots, newNotification)
	}

	subtx := tx.CreateInBatches(userNots, cfg.DatabaseBatchSize)

	if subtx.Error != nil {
		logger.LogflnIfExists("error", "Could not create user notifications: %v", logrus.ErrorLevel, cfg.LoggerLevelImportant, subtx.Error)
		return false
	}

	if err := dbConnection.CommitTransaction(true); err != nil {
		logger.LogflnIfExists("error", "Could not commit DB transaction: %v", logrus.ErrorLevel, cfg.LoggerLevelImportant, err)
		return false
	}

	logger.LogflnIfExists("debug", "Created %d user notifications", logrus.DebugLevel, cfg.LoggerLevelAll, len(userNots))

	return true
}

func getTextModelOrCreate(tx *gorm.DB, text string, request *v1.NotificationCreateManualRequest) (*NotificationText, error) {
	if !request.TextAsModel {
		return nil, nil
	}

	textModel := &NotificationText{}

	tx = tx.Where("text ILIKE ?", fmt.Sprintf("%%%s%%", text)).FirstOrCreate(textModel, &NotificationText{Text: text})

	return textModel, tx.Error
}
