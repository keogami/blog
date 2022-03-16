package main

import (
	"blog/api"
	"blog/db"
	mysql "blog/db/sql"
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
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

  queries := db.New(dbconn)
  api.SetDBQueries(queries)

  r := gin.Default()
  api.BindRoutes(r.Group("api"))

  r.Run()

  fmt.Println("that worked well")
}
