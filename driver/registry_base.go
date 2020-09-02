package driver

import (
	"github.com/oeoen/push-notifications/driver/config"
	"github.com/oeoen/push-notifications/pkg/handler"
	"github.com/oeoen/push-notifications/pkg/handler/rest"
)

type RegistryBase struct {
	r Registry
	c config.Provider
	h handler.Provider
}

func (b *RegistryBase) with(r Registry) *RegistryBase {
	b.r = r
	return b
}

func (b *RegistryBase) WithConfig(c config.Provider) *RegistryBase {
	b.c = c
	return b
}

func (b *RegistryBase) Handler() handler.Provider {
	if b.c.Service() != "rest" {

	}
	return rest.NewServer(b.r.NotificationManager(), b.c)
}
