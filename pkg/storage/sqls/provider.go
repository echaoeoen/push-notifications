package sqls

import (
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
	notification.StorageManager
	Stmts() sync.Map
	Prepare(query string) (*sql.Stmt, error)
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
