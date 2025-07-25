package database

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rulessly/gin-base/internal/ent"
	"github.com/rulessly/gin-base/internal/ent/migrate"
	"github.com/rulessly/gin-base/pkg/config"
	"github.com/spf13/cast"
	"log"
)

var DB *ent.Client

func InitPostgres() {
	db, err := sql.Open("pgx", getDsn())
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxIdleConns(cast.ToInt(config.Config.Postgres.MaxIdle))
	db.SetMaxOpenConns(cast.ToInt(config.Config.Postgres.MaxOpen))
	drv := entsql.OpenDB(dialect.Postgres, db)
	DB = ent.NewClient(ent.Driver(drv))
	err = DB.Schema.Create(context.Background(), migrate.WithForeignKeys(false))
	if err != nil {
		log.Fatal(err)
	}
}

func getDsn() string {
	username := config.Config.Postgres.User
	password := config.Config.Postgres.Password
	host := config.Config.Postgres.Host
	port := config.Config.Postgres.Port
	dbname := config.Config.Postgres.DbName
	cfg := config.Config.Postgres.Config
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s %s", host, port, username, password, dbname, cfg)
}
