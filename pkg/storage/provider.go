package storage

import (
	"github.com/oeoen/push-notifications/driver/config"
	"github.com/oeoen/push-notifications/pkg/notification"
)

type Provider interface {
	DBInit(c config.Provider) error
	DBDefer() error
	notification.StorageManager
}
