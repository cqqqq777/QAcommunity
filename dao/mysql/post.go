package mysqlDao

import (
	g "main/global"
	"main/modal"
)

func GetAllPostTitle(p []modal.Post) error {
	return g.Mdb.Select("title").Find(&p).Error
}

func PostDetail(p *modal.Post) error {
	return g.Mdb.Model(p).Where("post_id=?", p.PostID).First(p).Error
}

func QueryPostComment(c []modal.Comment, PostID int) error {
	return g.Mdb.Select("author_id", "content").Where("post_id = ?", PostID).Find(&c).Error
}

func PublishPost(p *modal.Post) error {
	return g.Mdb.Create(p).Error
}

func GetPersonalPost(p []modal.Post, authID int) error {
	return g.Mdb.Select("title", "content").Where("author_id = ?", authID).Find(&p).Error
}

func UpdatePost(p *modal.Post) error {
	return g.Mdb.Model(p).Where("post_id = ?", p.PostID).Updates(map[string]interface{}{"title": p.Title, "content": p.Content}).Error
}

func DeletePost(postID, authorID int) error {
	return g.Mdb.Where("post_id = ? and author_id = ?", postID, authorID).Delete(&modal.Post{}).Error
}
