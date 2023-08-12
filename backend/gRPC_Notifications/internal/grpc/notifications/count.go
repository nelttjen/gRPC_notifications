package notifications

import (
	"github.com/sirupsen/logrus"
	v1 "notification_grpc/api"
	"notification_grpc/internal/config"
	"notification_grpc/pkg/database"
	"notification_grpc/pkg/logger"
)

type notificationCount struct {
	Count          int `gorm:"column:count"`
	CountImportant int `gorm:"column:countimportant"`
}

func GetCountNotifications(conn database.PostgresConnection, userId uint64) (*v1.UserCountNotificationResponse, bool) {
	dbConn, _ := conn.GetDBConnection()

	massCounter := &notificationCount{}
	userCounter := &notificationCount{}
	response := &v1.UserCountNotificationResponse{
		Count:        0,
		HasImportant: false,
	}

	resultMass := dbConn.Raw(`
	SELECT 
	    count(umn.id) as Count,
       	count(CASE WHEN mn.important THEN 1 END) as CountImportant
	FROM user_mass_notifications umn
	LEFT JOIN mass_notifications mn on mn.id = umn.notification_id
	WHERE umn.user_id = ? AND umn.read = ?;
`, userId, false).Scan(massCounter)

	resultUser := dbConn.Raw(`
    SELECT 
        count(notif.id) as Count,
	    count(CASE WHEN notif.important THEN 1 END) as CountImportant
    FROM user_notifications notif
    WHERE notif.user_id = ? AND notif.read = ?;
`, userId, false).Scan(userCounter)

	if resultUser.Error != nil || resultMass.Error != nil {
		logger.LogflnIfExists("error", "One of the queries to get the count failed: mass query - %v, user query - %v", logrus.ErrorLevel, config.LoggerLevelImportant, resultMass.Error, resultUser.Error)
		return response, false
	}
	response.Count = uint32(massCounter.Count + userCounter.Count)
	response.HasImportant = massCounter.CountImportant > 0 || userCounter.CountImportant > 0
	
	logger.LogflnIfExists("debug", "User %d has %d notifications and %v important", logrus.DebugLevel, config.LoggerLevelAll, userId, response.Count, response.HasImportant)
	return response, true
}
