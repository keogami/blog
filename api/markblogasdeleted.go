package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MarkBlogAsDeleted(c *gin.Context) {
  slug := c.Param("slug")

  ctx := context.Background()
  q := GetDBQueries()

  count, err := q.MarkDeleted(ctx, slug)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{ "message": err.Error() })
    return
  }

  if count == int64(0) {
    c.JSON(http.StatusNotFound, gin.H{ "message": "BlogPost was not found" })
    return
  }

  post, err := q.GetPostBySlug(ctx, slug)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{ "message": err.Error() }) 
    return
  }

  c.JSON(http.StatusOK, PostFromRow(post))
}
