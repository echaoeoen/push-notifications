package config

import (
	"net/url"

	"github.com/ory/x/logrusx"
	"github.com/sirupsen/logrus"
	jConfig "github.com/uber/jaeger-client-go/config"
)

// Provider interface
type Provider interface {
	ServeHTTPS() bool
	DSN() string
	ListenHost() string
	ListenPort() string
	GetCookieSecrets() [][]byte
	Logger() *logrusx.Logger
	Service() string
	AppName() string
	FCMServerKey() string
	FetchNotificationSizePerReq() string
	TracingJaegerConfig() *jConfig.Configuration
}

func urlRoot(u *url.URL) *url.URL {
	if u.Path == "" {
		u.Path = "/"
	}
	return u
}

//validating
func MustValidate(l logrus.FieldLogger, p Provider) {

}
