package main

import (
	mysql "blog/db/sql"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {
  dbconn, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
  if err != nil {
    panic(fmt.Errorf("Couldn't open DB Connection: %w", err))
  }

  err = mysql.MigrateDB(dbconn)
  if err != nil {
    panic(fmt.Errorf("Couldn't migrate DB: %w", err))
  }

  fmt.Println("that worked well")
}
