package main

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBlogBySlug(c *gin.Context) {
  slug := c.Param("slug")

  ctx := context.Background()
  q := GetDBQueries()

  post, err := q.GetPostBySlug(ctx, slug)
  if err != nil {
    switch err {
    case sql.ErrNoRows:
      c.JSON(http.StatusNotFound, gin.H{ "message": "BlogPost not found." })
    default:
      c.JSON(http.StatusInternalServerError, gin.H{ "message": err.Error() })
    }
    return
  }

  c.JSON(http.StatusOK, PostFromRow(post))
}
