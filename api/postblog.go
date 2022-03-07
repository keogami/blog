package main

import (
	"blog/db"
	"context"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JsonNullString(s sql.NullString) interface{} {
	if s.Valid {
		return s.String
	}
	return nil
}

type PostBlogParams struct {
  Title string `json:"title"`
  Summary string `json:"summary"`
  Content string `json:"content"`
  Tags []string `json:"tags"`
}

func CreateSlugFromTitle(title string) string {
  return title
}

func PostBlog(c *gin.Context) {
  var params PostBlogParams
  if err := c.ShouldBindJSON(&params); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{ "message": err.Error(), })
    return
  }

  ctx := context.Background()
  q := GetDBQueries()

  postID, err := q.CreatePost(ctx, params.Content)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{ "message": err.Error() })
    return
  }

  meta, err := q.CreateMeta(ctx, db.CreateMetaParams{
    PostID: postID,
    Slug: CreateSlugFromTitle(params.Title),
    Title: params.Title,
    Summary: params.Summary,
    GroupName: sql.NullString{},
    Tags: params.Tags,
  })
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{ "message": err.Error() })
    return
  }

  c.JSON(http.StatusCreated, gin.H{
    "content": params.Content,
    "meta": gin.H{
      "slug": meta.Slug,
      "title": meta.Title,
      "group": JsonNullString(meta.GroupName),
      "summary": meta.Summary,
      "postTime": meta.PostTime,
      "tags": meta.Tags,
    },
  })
}
