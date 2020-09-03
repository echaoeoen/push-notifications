package driver

import (
	"github.com/oeoen/push-notifications/driver/config"
	"github.com/oeoen/push-notifications/pkg/handler"
	"github.com/oeoen/push-notifications/pkg/handler/rest"
	"github.com/oeoen/push-notifications/pkg/notification"
)

type RegistryBase struct {
	r Registry
	c config.Provider
	h handler.Provider
	m notification.Manager
}

func (b *RegistryBase) with(r Registry) *RegistryBase {
	b.r = r
	return b
}

func (b *RegistryBase) WithConfig(c config.Provider) *RegistryBase {
	b.c = c
	return b
}
func (b *RegistryBase) WithNotificationManager(m notification.Manager) *RegistryBase {
	b.m = m
	return b
}

func (b *RegistryBase) Handler() handler.Provider {
	if b.c.Service() != "rest" {

	}
	return rest.NewServer(b.r.NotificationManager(), b.c)
}

func (b *RegistryBase) NotificationManager() notification.Manager {
	return b.m
}
