package notifications

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"log"
	v1 "notification_grpc/api"
	cfg "notification_grpc/internal/config"
	"notification_grpc/pkg/database"
	"time"
)

func ProcessUserRequest(dbConnection database.PostgresConnection, request *v1.NotificationCreateManualRequest) bool {
	tx, err := dbConnection.GetDBTransaction()
	if err != nil {
		log.Printf("ERROR: getting DB transaction: %s", err)
		return false
	}

	if request.UserIds == nil || len(request.UserIds) == 0 || request.Text == nil {
		log.Printf("WARNING: get empty list of user IDs or text is empty, request ignored")
		return true
	}

	mongoClient := database.NewMongoConnection()
	err = mongoClient.MakeConnection()
	if err != nil {
		log.Printf("ERROR: connecting to MongoDB: %s", err)
		return false
	}

	defer func(cl database.MongoClient) {
		err := cl.CloseConnection()
		if err != nil {
			log.Printf("Error closing MongoDB connection: %s", err)
		}
	}(mongoClient)

	collection, err := mongoClient.GetCollectionFromAdmin("user_notification_settings")
	if err != nil {
		log.Printf("ERROR: getting MongoDB collection: %s", err)
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
		log.Printf("ERROR: getting MongoDB cursor: %s", err)
		return false
	}
	defer func(c *mongo.Cursor, ctx context.Context) {
		err := c.Close(ctx)
		if err != nil {
			log.Printf("Error closing MongoDB cursor: %s", err)
		}
	}(cursor, ctx)

	var notUserIds []int
	for cursor.Next(ctx) {
		user := &UserNotificationConfig{}
		err := cursor.Decode(user)
		if err != nil {
			log.Printf("ERROR: decoding MongoDB cursor: %s", err)
			return false
		}
		notUserIds = append(notUserIds, user.UserID)
	}

	var userNots []*UserNotification

	useTargets := request.TargetId != nil && request.TargetType != nil
	useText, err := getTextModelOrCreate(tx, *request.Text, request)

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
			newNotification.Text = *request.Text
		}

		userNots = append(userNots, newNotification)
	}

	subtx := tx.CreateInBatches(userNots, cfg.DatabaseBatchSize)

	if subtx.Error != nil {
		log.Printf("ERROR: creating user notifications: %s", subtx.Error)
		return false
	}

	if err := dbConnection.CommitTransaction(true); err != nil {
		log.Printf("ERROR: committing DB transaction: %s", err)
		return false
	}

	log.Printf("INFO: created %d user notifications", len(userNots))

	return true
}

func getTextModelOrCreate(tx *gorm.DB, text string, request *v1.NotificationCreateManualRequest) (*NotificationText, error) {
	if !request.TextAsModel {
		return nil, nil
	}

	textModel := &NotificationText{}

	tx = tx.Where("text ILIKE ?", fmt.Sprintf("%%%s%%", text)).First(textModel)

	if err := tx.Error; err != nil {
		log.Printf("ERROR: get text model: %s", err)
		return nil, err
	}

	if textModel.ID == 0 {
		textModel.Text = text
		tx = tx.Create(textModel)
		if err := tx.Error; err != nil {
			log.Printf("ERROR: creating text model: %s", err)
			return nil, err
		}
	}

	return textModel, nil
}
