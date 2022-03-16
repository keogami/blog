package api

import (
	"blog/db"
	"fmt"

	_ "github.com/lib/pq"
)

var queries *db.Queries
func GetDBQueries() *db.Queries {
  return queries
}

func SetDBQueries(q *db.Queries) {
  if q == nil {
    panic(fmt.Errorf("The queries provided to API is nil"))
  }
  queries = q
}

