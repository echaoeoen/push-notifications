package rest

import (
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/oeoen/push-notifications/driver/config"
	"github.com/oeoen/push-notifications/helper/errorp"
	"github.com/oeoen/push-notifications/pkg/notification"
)

type RestServer struct {
	service *echo.Echo
	m       notification.Manager
	c       config.Provider
}

func NewServer(m notification.Manager, c config.Provider) *RestServer {
	s := new(RestServer)
	s.service = echo.New()
	s.m = m
	s.c = c
	return s
}

func (s *RestServer) Serve() error {
	s.service.Use(TracingMiddleware("echo"))
	getRoutes(s.service, s.m, s.c)
	err := s.service.Start(s.Configuration().ListenHost() + ":" + s.Configuration().ListenPort())
	return err
}

func (s *RestServer) NotificationManager() notification.Manager {
	return s.m
}
func (s *RestServer) Configuration() notification.Configuration {
	return s.c
}

func responseError(c echo.Context, err error) error {
	e, ok := err.(*errorp.NotificationError)
	if ok {
		return c.JSON(int(e.Status()), e)
	}
	return err
}

func getQueries(q url.Values) [][3]string {
	var r [][3]string
	for k := range q {
		val := q.Get(k)
		op := operator(&val)
		r = append(r, [3]string{k, op, val})
	}
	return r
}

func operator(s *string) string {
	ops := []string{
		"LIKE",
		"<=",
		">=",
		"<",
		">",
	}
	for i := 0; i < len(ops); i++ {
		o := checkOperator(ops[i], *s)
		if o {
			*s = (*s)[len(ops[i]):len(*s)]
			return ops[i]
		}
	}
	return ops[0]
}

func checkOperator(o, s string) bool {
	if len(s) > len(o) {
		if s[0:len(o)] == o {
			return true
		}
	}
	return false
}

func parseParam(c echo.Context, p string) string {
	v := c.Param(p)
	parsed, err := url.QueryUnescape(v)
	if err != nil {
		return v
	}
	return parsed
}
