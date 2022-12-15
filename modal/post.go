package modal

import "time"

type Post struct {
	ID       int
	AuthorID int `gorm:"column:author_id"`
	PostID   int `gorm:"column:post_id"`
	Title    string
	Content  string
	CreateAt *time.Time `gorm:"column:createAt;<-:false"`
	UpdateAt *time.Time `gorm:"column:createAt;<-:false"`
}

type ParamPost struct {
	Title   string `form:"title"`
	Content string `form:"content"`
	PostID  int    `form:"post-id"`
}
