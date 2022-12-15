package services

import (
	"encoding/json"
	mysqlDao "main/dao/mysql"
	"main/modal"
	"main/utils"
)

func GetAllPostTitle() (string, error) {
	post := make([]modal.Post, 0)
	if err := mysqlDao.GetAllPostTitle(post); err != nil {
		return "", err
	}
	var titleSlice = make([]string, 0)
	for _, val := range post {
		titleSlice = append(titleSlice, val.Title)
	}
	title, err := json.Marshal(titleSlice)
	if err != nil {
		return "", err
	}
	return string(title), nil
}

func PostDetail(id int) (*modal.Post, error) {
	Post := new(modal.Post)
	Post.PostID = id
	err := mysqlDao.PostDetail(Post)
	if err != nil {
		return nil, err
	}
	return Post, err
}

func QueryPostComment(postID int) (string, error) {
	commentSlice := make([]modal.Comment, 0)
	if err := mysqlDao.QueryPostComment(commentSlice, postID); err != nil {
		return "", err
	}
	//创建一个map切片存放帖子的评论
	var comment = make([]map[int]string, 0)
	for _, val := range commentSlice {
		comment = append(comment, map[int]string{val.AuthorID: val.Content})
	}
	//把切片序列化后转成string返回
	comments, err := json.Marshal(comment)
	if err != nil {
		return "", nil
	}
	return string(comments), nil
}

func PublishPost(p *modal.ParamPost, authorID int) error {
	post := new(modal.Post)
	ID, err := utils.GetID()
	if err != nil {
		return err
	}
	post.PostID = ID
	post.AuthorID = authorID
	post.Content = p.Content
	post.Title = p.Title
	return mysqlDao.PublishPost(post)
}

func GetPersonalPost(authorID int) (string, error) {
	postsSlice := make([]modal.Post, 0)
	if err := mysqlDao.GetPersonalPost(postsSlice, authorID); err != nil {
		return "", err
	}
	posts := make([]map[string]string, 0)
	for _, val := range postsSlice {
		posts = append(posts, map[string]string{val.Title: val.Content})
	}
	postsStr, err := json.Marshal(posts)
	if err != nil {
		return "", err
	}
	return string(postsStr), nil
}

func UpdatePost(p *modal.ParamPost, authorID int) error {
	post := new(modal.Post)
	post.PostID = p.PostID
	post.Title = p.Title
	post.Content = p.Content
	post.AuthorID = authorID
	return mysqlDao.UpdatePost(post)
}

func DeletePost(authorID, postID int) error {
	return mysqlDao.DeletePost(postID, authorID)
}
