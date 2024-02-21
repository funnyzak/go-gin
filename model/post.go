package model

import (
	"go-gin/mappers"

	"gorm.io/gorm"
)

type Post struct {
	Common
	Title       string `json:"title,omitempty" gorm:"column:title"`
	Content     string `json:"content,omitempty" gorm:"column:content"`
	CreatedUser uint64 `json:"created_user,omitempty" gorm:"column:created_user"`
}

func NewPost() *Post {
	return &Post{}
}

func (p *Post) Create(form mappers.PostForm, db *gorm.DB) (err error) {
	p.Title = form.Title
	p.Content = form.Content
	p.CreatedUser = form.CreatedUser
	err = db.Model(&Post{}).Create(&p).Error
	return err
}

func (p *Post) Update(form mappers.PostForm, db *gorm.DB) (err error) {
	if form.ID == 0 {
		return err
	}
	p.Title = form.Title
	p.Content = form.Content
	p.CreatedUser = form.CreatedUser
	err = db.Model(&Post{}).Where("id = ?", form.ID).Updates(&p).Error
	return err
}

func (p *Post) Delete(id int, db *gorm.DB) (err error) {
	err = db.Model(&Post{}).Where("id = ?", id).Delete(&p).Error
	return err
}

func (p *Post) Get(id int, db *gorm.DB) (err error) {
	err = db.Model(&Post{}).Where("id = ?", id).First(&p).Error
	return err
}

func (p *Post) List(db *gorm.DB, query interface{}, args ...interface{}) (posts []Post, err error) {
	err = db.Model(&Post{}).Where(query, args).Find(&posts).Error
	return posts, err
}
