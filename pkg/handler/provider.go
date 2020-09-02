package handler

import "github.com/oeoen/push-notifications/pkg/notification"

type Provider interface {
	Serve() error
	notification.Registry
}
