package main

import (
  "time"
  "blog/db"
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

type Post struct {
  Meta Meta `json:"meta"`
  Content string `json:"content"`
}

func PostFromPublicRow(post db.GetPublicPostBySlugRow) Post {
  return Post{
    Meta: Meta{
      Slug: post.Slug,
      Title: post.Title,
      Group: JsonNullString(post.GroupName),
      PostTime: post.PostTime,
      Summary: post.Summary,
      Tags: post.Tags,
    },
    Content: post.Content,
  }
}

func PostFromRow(post db.GetPostBySlugRow) Post {
  return Post{
    Meta: Meta{
      Slug: post.Slug,
      Title: post.Title,
      Group: JsonNullString(post.GroupName),
      PostTime: post.PostTime,
      Summary: post.Summary,
      Tags: post.Tags,
    },
    Content: post.Content,
  }
}

func PostFromMeta(content string, meta db.Meta) Post {
  return Post{
    Meta: MetaFrom(meta),
    Content: content,
  }
}
