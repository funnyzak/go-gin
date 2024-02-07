package model

import (
	"go-gin/mappers"

	"gorm.io/gorm"
)

type Post struct {
	Common
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	Author  string `json:"author,omitempty"`
}

func (p Post) Create(form mappers.PostForm, db *gorm.DB) (err error) {
	p.Title = form.Title
	p.Content = form.Content
	p.Author = form.Author
	err = db.Model(&Post{}).Create(&p).Error
	return err
}

func (p Post) Update(form mappers.PostForm, db *gorm.DB) (err error) {
	if form.ID == 0 {
		return err
	}
	p.Title = form.Title
	p.Content = form.Content
	p.Author = form.Author
	err = db.Model(&Post{}).Where("id = ?", form.ID).Updates(&p).Error
	return err
}
