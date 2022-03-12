package main

import (
	"blog/db"
	"context"
	"database/sql"
	"log"
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

type PostBlogParams struct {
	Title   string   `json:"title"`
	Summary string   `json:"summary"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

// TODO change CreateSlugFromTitle such that, for an input title like "[4] Making a compiler" it returns "4-making-a-compiler"
//
// Requirements:
// 1. all non-alphanumeric characters should be removed
// 2. all characters should be lower case
// 3. all spaces should be replaced with "-"
//
// Hint:
// 1. use the "strings" package
// 2. run `go test` command to check your implementation

//func to createSlugFromTitle
func CreateSlugFromTitle(title string) string {
	re, err := regexp.Compile("[^a-zA-Z0-9]+")

	if err != nil {
		log.Fatal(err)
	}

	RemoveNonAlpha := re.ReplaceAllLiteralString(title, "-")
	LowerCase := strings.ToLower(RemoveNonAlpha)
	LowerCase = strings.Trim(LowerCase, "-")
	return LowerCase
}

func PostBlog(c *gin.Context) {
	var params PostBlogParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx := context.Background()
	q := GetDBQueries()

	postID, err := q.CreatePost(ctx, params.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	meta, err := q.CreateMeta(ctx, db.CreateMetaParams{
		PostID:    postID,
		Slug:      CreateSlugFromTitle(params.Title),
		Title:     params.Title,
		Summary:   params.Summary,
		GroupName: sql.NullString{},
		Tags:      params.Tags,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, PostFromMeta(params.Content, meta))
}
