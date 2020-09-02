package driver

import (
	"github.com/oeoen/push-notifications/driver/config"
	"github.com/oeoen/push-notifications/pkg/handler"
	"github.com/oeoen/push-notifications/pkg/notification"
	"github.com/oeoen/push-notifications/pkg/tracing"
	"github.com/sirupsen/logrus"
)

type Registry interface {
	Tracer() *tracing.Tracer
	WithConfig(c config.Provider) Registry
	WithLogger(l logrus.FieldLogger) Registry
	Init() error
	notification.Registry
	Handler() handler.Provider
}

func CallRegistry(r Registry) {
	r.Tracer()
}
