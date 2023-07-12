package config

import (
	"errors"
	"fmt"
)

const ActionTitleNewName = "title_new_name"
const ActionTitleNewChapter = "title_new_chapter"
const ActionSiteNotification = "site_notification"
const ActionChapterFree = "title_chapter_free"

func GetActionByKey(key string) (val int32, err error) {
	switch key {
	case ActionTitleNewName:
		val = 1
	case ActionTitleNewChapter:
		val = 2
	case ActionSiteNotification:
		val = 3
	case ActionChapterFree:
		val = 4

	default:
		val = 0
		err = errors.New(fmt.Sprintf("Invalid action key: %s", key))
	}

	return
}
