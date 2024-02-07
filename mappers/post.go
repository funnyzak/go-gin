package mappers

type PostForm struct {
	ID      int    `form:"id" json:"id" binding:"omitempty"`
	Title   string `form:"title" json:"title" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
	Author  string `form:"author" json:"author" binding:"required"`
}
