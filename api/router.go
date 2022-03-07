package main

import (
	"github.com/gin-gonic/gin"
)

func BindRoutes(r *gin.Engine) {
  r.GET("/blog", GetBlogs)
  r.GET("/blog/:slug", GetBlogBySlug)

  r.POST("/blog", PostBlog)
}
