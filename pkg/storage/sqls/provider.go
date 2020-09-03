package sqls

import (
	"context"
	"database/sql"
	"strings"
	"sync"

	"github.com/oeoen/push-notifications/driver/config"
	"github.com/oeoen/push-notifications/pkg/notification"
	"github.com/oeoen/push-notifications/pkg/storage/sqls/mysql"
)

type SQLSManager interface {
	DoMigration(c config.Provider, args ...string)
	DBService() *sql.DB
	Open() error
	Close() error
	Stmts() sync.Map
	Prepare(query string) (*sql.Stmt, error)
	Config() config.Provider

	SetFCMToken(ctx context.Context, application, username string, token notification.FCMToken) error
	GetFCMToken(ctx context.Context, application, username string) (*notification.UserData, error)
	SaveNotification(ctx context.Context, application, username string, content notification.Content) error
	FetchNotification(ctx context.Context, filter ...[3]string) ([]*notification.Content, error)
	ReadNotification(ctx context.Context, application, username, notificationID string) error
}

type SQLs struct {
	dbManager SQLSManager
}

func NewSQLS(c config.Provider) *SQLs {
	p := &SQLs{}
	p.DBInit(c)
	return p
}

func (p *SQLs) DBInit(c config.Provider) error {
	databaseKind := getDSN(c.DSN())
	if databaseKind == "mysql" {
		p.dbManager = mysql.NewManager(c)
		err := p.Manager().Open()
		if err != nil {
			c.Logger().Fatal(err)
		}
	}
	return nil
}

func (p *SQLs) DBDefer() error {
	return p.Manager().Close()
}
func (p *SQLs) DoMigration(c config.Provider, args ...string) {
	p.dbManager.DoMigration(c, args...)
}

func getDSN(dsn string) string {
	if dsn == "" {
		return "No DSN"
	}
	if strings.Contains(dsn, "mysql") {
		return "mysql"
	}
	return ""
}
func (p *SQLs) Manager() SQLSManager {
	return p.dbManager
}

func (p *SQLs) SetFCMToken(ctx context.Context, application, username string, token notification.FCMToken) error {
	return p.dbManager.SetFCMToken(ctx, application, username, token)
}
func (p *SQLs) GetFCMToken(ctx context.Context, application, username string) (*notification.UserData, error) {
	return p.dbManager.GetFCMToken(ctx, application, username)
}
func (p *SQLs) SaveNotification(ctx context.Context, application, username string, content notification.Content) error {
	return p.dbManager.SaveNotification(ctx, application, username, content)
}
func (p *SQLs) FetchNotification(ctx context.Context, application, username string, filter ...[3]string) ([]*notification.Content, error) {
	filter = append(filter, [][3]string{
		{"application", "=", application},
		{"username", "=", username},
		{"size", "=", p.dbManager.Config().FetchNotificationSizePerReq()},
	}...)
	return p.dbManager.FetchNotification(ctx, filter...)
}
func (p *SQLs) ReadNotification(ctx context.Context, application, username, notificationID string) error {
	return p.dbManager.ReadNotification(ctx, application, username, notificationID)
}
