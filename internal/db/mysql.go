package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/halowahyudi/rest-api-structure-golang/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init(cfg *config.Config) error {
    var err error

    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
        cfg.Database.User,
        cfg.Database.Password,
        cfg.Database.Host,
        cfg.Database.Name,
    )

    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        return fmt.Errorf("could not open db: %w", err)
    }

    if err = DB.Ping(); err != nil {
        return fmt.Errorf("could not ping db: %w", err)
    }

    log.Println("Connected to the database successfully")
    return nil
}

func Close() {
    if err := DB.Close(); err != nil {
        log.Fatalf("Could not close db: %v", err)
    }
}
