package driver

import (
	"github.com/oeoen/push-notifications/driver/config"
)

type Driver interface {
	Configuration() config.Provider
	Registry() Registry
	CallRegistry() Driver
}
