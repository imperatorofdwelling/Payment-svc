package postgres

import (
	"database/sql"
	"fmt"
	"github.com/imperatorofdwelling/payment-svc/internal/config"
	_ "github.com/lib/pq"
)

func NewPsqlStorage(c config.Postgres) (*sql.DB, error) {
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", c.Host, c.Port, c.Username, c.Password, c.DbName, c.SSLMode)

	fmt.Println(psqlConn)

	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		return nil, fmt.Errorf("could not connect to postgres: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("could not ping postgres: %v", err)
	}

	return db, nil
}
