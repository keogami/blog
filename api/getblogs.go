package main

import (
	"blog/db"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Meta struct {
  Slug string `json:"slug"`
  Title string `json:"title"`
  Group interface{} `json:"group"`
  PostTime time.Time `json:"postTime"`
  Summary string `json:"summary"`
  Tags []string `json:"tags"`
}

func MetaFrom(m db.Meta) Meta {
  return Meta{
    Slug: m.Slug,
    Title: m.Title,
    Group: JsonNullString(m.GroupName),
    PostTime: m.PostTime,
    Summary: m.Summary,
    Tags: m.Tags,
  }
}

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
