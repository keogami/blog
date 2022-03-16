package api

import (
	"blog/db"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PutBlog(c *gin.Context) {
  var post Post
  if err := c.ShouldBindJSON(&post); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{ "message": err.Error() })
    return
  }
  slug := c.Param("slug")

  ctx := context.Background()
  q := GetDBQueries()

  id, err := q.UpdatePost(ctx, db.UpdatePostParams{
    Slug: slug,
    Content: post.Content,
  })
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{ "message": err.Error() })
    return
  }

  err = q.UpdateMeta(ctx, db.UpdateMetaParams{
    PostID: id,
    Slug: post.Meta.Slug,
    Title: post.Meta.Title,
    Summary: post.Meta.Summary,
    GroupName: NullString(post.Meta.Group),
    Tags: post.Meta.Tags,
  })
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{ "message": err.Error() })
    return
  }

  c.JSON(http.StatusOK, post) // TODO stop being lazy
}
