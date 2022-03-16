package api

import (
	"github.com/gin-gonic/gin"
)

func BindRoutes(r gin.IRouter) {
  r.GET("/blog", GetBlogs)
  r.GET("/blog/:slug", GetPublicBlogBySlug)

  r.POST("/blog", PostBlog)
  r.PUT("/blog/:slug", PutBlog)
  r.DELETE("/blog/:slug", MarkBlogAsDeleted)
}
