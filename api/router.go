package main

import (
	"github.com/gin-gonic/gin"
)

func BindRoutes(r *gin.Engine) {
  r.GET("/blog", GetBlogs)
  r.GET("/blog/:slug", GetPublicBlogBySlug)

  r.POST("/blog", PostBlog)
  r.PUT("/blog/:slug", PutBlog)
}
