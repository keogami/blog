package main

import (
	"blog/db"
	mysql "blog/db/sql"
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var queries *db.Queries
func GetDBQueries() *db.Queries {
  return queries
}

func main() {
  dbconn, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
  if err != nil {
    panic(fmt.Errorf("Couldn't open DB Connection: %w", err))
  }

  err = mysql.MigrateDB(dbconn)
  if err != nil {
    panic(fmt.Errorf("Couldn't migrate DB: %w", err))
  }

  queries = db.New(dbconn)

  r := gin.Default()
  BindRoutes(r)

  r.Run()

  fmt.Println("that worked well")
}
