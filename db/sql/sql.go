package sql

import (
	"database/sql"
	"embed"
  "github.com/pressly/goose/v3"
)

//go:embed 001_schema.sql
var migrationSchema embed.FS


func MigrateDB(db *sql.DB) error {
  goose.SetBaseFS(migrationSchema)

  if err := goose.SetDialect("postgres"); err != nil {
    return err
  }

  if err := goose.Up(db, ".", goose.WithNoVersioning()); err != nil {
    return err
  }

  return nil
}
