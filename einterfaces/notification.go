// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package einterfaces

import (
	"github.com/mattermost/mattermost-server/model"
)

type NotificationInterface interface {
	GetNotificationMessage(ack *model.PushNotificationAck, userId string) (*model.PushNotification, *model.AppError)
	CheckLicense() *model.AppError
}
