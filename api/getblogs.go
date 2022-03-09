package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBlogs(c *gin.Context) {
  ctx := context.Background()
  q := GetDBQueries()

  dbmetas, err := q.ListMetas(ctx)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{ "message": err.Error() })
    return
  }

  metas := make([]Meta, 0, len(dbmetas))
  for _, item := range dbmetas {
    metas = append(metas, MetaFrom(item))
  }

  c.JSON(http.StatusOK, metas)
}
