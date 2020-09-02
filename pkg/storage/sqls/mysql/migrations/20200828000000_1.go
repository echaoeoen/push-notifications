package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20200828000000_1_init, Down20200828000000_1_init)
}

func Up20200828000000_1_init(tx *sql.Tx) error {
	_, err := tx.Exec(`
    CREATE TABLE notifications (
				id BIGINT NOT NULL AUTO_INCREMENT,
				username varchar NOT NULL,
				application varchar NOT NULL,
				title VARCHAR(100) NOT NULL,
				subtitle VARCHAR(255) NOT NULL,
				message TEXT NOT NULL,
				action VARCHAR(100) NOT NULL,
				param VARCHAR(255) NOT NULL,
				readed TINYINT(1) DEFAULT 0,
        created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT primary_policy PRIMARY KEY (id)
    );
  `)
	if err != nil {
		return err
	}
	return nil
}

func Down20200828000000_1_init(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE IF EXISTS notifications;")
	if err != nil {
		return err
	}
	return nil
}
