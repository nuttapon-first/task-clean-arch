package databases

import (
	"database/sql"
	"log"

	"github.com/nuttapon-first/task-clean-arch/configs"
	_ "github.com/proullon/ramsql/driver"
)

func NewSQLDBConnection(config *configs.Configs) (*sql.DB, error) {
	db, err := sql.Open(config.Database.DriverName, config.Database.DbName)
	if err != nil {
		log.Fatal("error:", err)
		return nil, err
	}

	createTb := `
	CREATE TABLE IF NOT EXISTS tasks (
	id INT AUTO_INCREMENT,
	name TEXT NOT NULL,
	status INT NOT NULL,
	PRIMARY KEY (id)
	);
	`

	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("error:", err)
		return nil, err
	}

	return db, nil
}
