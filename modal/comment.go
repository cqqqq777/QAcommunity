package modal

import "time"

type Comment struct {
	ID        int
	CommentID int `gorm:"column:comment_id"`
	AuthorID  int `gorm:"column:author_id"`
	PostID    int `gorm:"column:post_id"`
	ParentID  int `gorm:"column:parent_id"`
	RootID    int `gorm:"column:root_id"`
	Content   string
	CreateAt  *time.Time `gorm:"column:createAt;<-:false"`
	UpdateAt  *time.Time `gorm:"column:updateAt;<-:false"`
}

type ParamComment struct {
	AuthorID int
	PostID   int    `form:"post-id" binding:"required"`
	ParentID int    `form:"parent-id" binding:"required"`
	RootID   int    `form:"root-id" binding:"required"`
	Content  string `form:"content" binding:"required"`
}
