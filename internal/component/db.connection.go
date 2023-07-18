package component

import (
	"database/sql"
	"fmt"
	"gofiber-ewallet/internal/config"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func ConnectDB(config *config.Config) *sql.DB {
	connStr := fmt.Sprintf(
		"host=%s "+"port=%s "+"user%s "+"password%s "+"dbName%s "+"sslmode=disable",
		config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password,
		config.Database.Name,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Printf("Failed Connect Database %s", err.Error())
	}

	logger.Println("Database Connected")

	return db
}
