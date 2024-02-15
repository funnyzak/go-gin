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

func (p *Post) Get(id int, db *gorm.DB) (post Post, err error) {
	err = db.Model(&Post{}).Where("id = ?", id).First(&post).Error
	return post, err
}

func (p *Post) List(db *gorm.DB, where ...interface{}) (posts []Post, err error) {
	if len(where) == 0 {
		err = db.Model(&Post{}).Find(&posts).Error
		return posts, err
	}
	err = db.Model(&Post{}).Where(where).Find(&posts).Error
	return posts, err
}
