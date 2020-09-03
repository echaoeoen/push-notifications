package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20200828000000_users_init, Down20200828000000_users_init)
}

func Up20200828000000_users_init(tx *sql.Tx) error {
	_, err := tx.Exec(`
    CREATE TABLE users (
				id BIGINT NOT NULL AUTO_INCREMENT,
				application VARCHAR(255) NOT NULL,
				username VARCHAR(255) NOT NULL,
				fcm_token VARCHAR(255) NOT NULL,
        created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT primary_policy PRIMARY KEY (id)
    );
  `)
	if err != nil {
		return err
	}
	return nil
}

func Down20200828000000_users_init(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE IF EXISTS users;")
	if err != nil {
		return err
	}
	return nil
}
