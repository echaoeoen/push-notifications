package driver

import (
	"github.com/oeoen/push-notifications/access"
	"github.com/oeoen/push-notifications/driver/config"
	"github.com/oeoen/push-notifications/pkg/notification"
	"github.com/oeoen/push-notifications/pkg/storage"
	"github.com/oeoen/push-notifications/pkg/storage/sqls"
	"github.com/oeoen/push-notifications/pkg/tracing"
	"github.com/ory/x/logrusx"
)

type RegistrySQLs struct {
	*RegistryBase
	l  *logrusx.Logger
	c  config.Provider
	t  tracing.Trace
	db storage.Provider
}

func NewRegistrySqls(c config.Provider) (Registry, error) {
	registry := &RegistrySQLs{
		c:            c,
		l:            c.Logger(),
		db:           sqls.NewSQLS(c),
		RegistryBase: new(RegistryBase),
	}

	registry.RegistryBase.with(registry).WithConfig(c).WithNotificationManager(access.NewManager(registry.db, c))

	if err := registry.Init(); err != nil {
		return nil, err
	}

	return registry, nil
}

func (r *RegistrySQLs) Tracer() *tracing.Tracer {
	return r.Tracer()
}

func (r *RegistrySQLs) WithConfig(c config.Provider) Registry {
	r.c = c
	return r
}

func (r *RegistrySQLs) WithLogger(l *logrusx.Logger) Registry {
	r.l = l
	return r
}
func (r *RegistrySQLs) Init() error {
	return nil
}
func (r *RegistrySQLs) Provider() storage.Provider {
	return r.db
}
func (r *RegistrySQLs) NotificationManager() notification.Manager {
	return r.RegistryBase.m
}
func (r *RegistrySQLs) Configuration() notification.Configuration {
	return r.c
}
