package access

import (
	"context"

	"github.com/oeoen/push-notifications/driver/config"
	"github.com/oeoen/push-notifications/helper/errorp"
	"github.com/oeoen/push-notifications/pkg/fcm"
	"github.com/oeoen/push-notifications/pkg/notification"
)

type Manager struct {
	s notification.StorageManager
	c config.Provider
}

func NewManager(s notification.StorageManager, c config.Provider) *Manager {
	return &Manager{
		s: s,
		c: c,
	}
}
func (m *Manager) StorageManager() notification.StorageManager {
	return m.s
}

func (m *Manager) SendNotification(ctx context.Context, application, username string, content notification.Content) error {
	if err := m.s.SaveNotification(ctx, application, username, content); err != nil {
		return err
	}
	Token, err := m.s.GetFCMToken(ctx, application, username)
	if err != nil {
		return errorp.NewNotificationError(201, "warn_not_send_to_fcm", "Token for the username not setted", err.Error())
	}
	f := fcm.New(m.c)
	if err = f.Send(Token.FCMToken.FCMToken, content); err != nil {
		return err
	}
	return nil
}
