package api

import (
	"blog/db"
	"context"
	"database/sql"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

func JsonNullString(s sql.NullString) interface{} {
	if s.Valid {
		return s.String
	}
	return nil
}

func NullString(s interface{}) sql.NullString {
  switch s.(type) {
  case string: return sql.NullString{
    Valid: true,
    String: s.(string),
  }
  }
  return sql.NullString{ String: "", Valid: false }
}

type PostBlogParams struct {
  Title string `json:"title"`
  Summary string `json:"summary"`
  Content string `json:"content"`
  Tags []string `json:"tags"`
}

var NonAlphaRegex = regexp.MustCompile(`[^A-Za-z0-9\s]`)
var MultipleSpaceRegex = regexp.MustCompile(`[\s]+`)

func CreateSlugFromTitle(title string) string {
  res := NonAlphaRegex.ReplaceAllString(title, "")
  res = MultipleSpaceRegex.ReplaceAllString(res, "-")
  return strings.ToLower(res)
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

  c.JSON(http.StatusCreated, PostFromMeta(params.Content, meta))
}
