package config

import (
	"strings"

	"github.com/ory/viper"
	"github.com/ory/x/logrusx"
	"github.com/ory/x/viperx"
	jConfig "github.com/uber/jaeger-client-go/config"
)

type ViperProvider struct {
	l               *logrusx.Logger
	ss              [][]byte
	generatedSecret []byte
	forcedHTTP      bool
}

const (
	ViperKeyPublicURL                = "urls.public"
	ViperKeyDSN                      = "dsn"
	ViperKeyHost                     = "serve.host"
	ViperKeyPort                     = "serve.port"
	ViperKeyGetCookieSecrets         = "secrets.cookie"
	ViperKeyService                  = "service"
	ViperFCMServerKey                = "fcm.server.key"
	ViperFetchNotificationSizePerReq = "fetch.notification.size.per.request"
)

func init() {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

}

func NewViperProvider(l *logrusx.Logger, forcedHTTP bool) Provider {
	return &ViperProvider{
		l:          l,
		forcedHTTP: forcedHTTP,
	}
}

func (v *ViperProvider) DSN() string {
	return viperx.GetString(v.l, ViperKeyDSN, "", "DATABASE_URL")
}

func (v *ViperProvider) TracingJaegerConfig() *jConfig.Configuration {
	c := &jConfig.Configuration{
		Reporter: &jConfig.ReporterConfig{
			CollectorEndpoint: viperx.GetString(v.l, "tracing.providers.jaeger.endpoint", "", "TRACING_PROVIDER_JAEGER_ENDPOINT"),
		},
		ServiceName: "push-notifications",
		Sampler: &jConfig.SamplerConfig{
			SamplingServerURL: viperx.GetString(v.l, "tracing.providers.jaeger.sampling.server", "", "TRACING_PROVIDER_JAEGER_SAMPLING_SERVER_URL"),
			Param:             viperx.GetFloat64(v.l, "tracing.providers.jaeger.sampling.value", float64(1), "TRACING_PROVIDER_JAEGER_SAMPLING_VALUE"),
			Type:              viperx.GetString(v.l, "tracing.providers.jaeger.sampling.type", "const", "TRACING_PROVIDER_JAEGER_SAMPLING_TYPE"),
		},
	}

	return c
}

func (v *ViperProvider) GetCookieSecrets() [][]byte {
	return [][]byte{
		[]byte(viperx.GetString(v.l, ViperKeyGetCookieSecrets, "", "COOKIE_SECRET")),
	}
}

func (v *ViperProvider) ListenHost() string {
	return viperx.GetString(v.l, ViperKeyHost, "", "HOST")
}

func (v *ViperProvider) ListenPort() string {
	return viperx.GetString(v.l, ViperKeyPort, "", "PORT")
}
func (v *ViperProvider) Service() string {
	return viperx.GetString(v.l, ViperKeyService, "", "SERVICE")
}

func (v *ViperProvider) ServeHTTPS() bool {
	return !v.forcedHTTP
}
func (v *ViperProvider) Logger() *logrusx.Logger {
	return v.l
}

func (v *ViperProvider) AppName() string {
	return "push-notifications"
}
func (v *ViperProvider) FCMServerKey() string {
	return viperx.GetString(v.l, ViperFCMServerKey, "", "FCM_SERVER_KEY")
}
func (v *ViperProvider) FetchNotificationSizePerReq() string {
	return viperx.GetString(v.l, ViperFCMServerKey, "10", "FETCH_NOTIFICATION_SIZE_PER_REQUEST")

}
