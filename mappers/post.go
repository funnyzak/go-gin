package mappers

type PostForm struct {
	ID          int    `form:"id" json:"id" binding:"omitempty"`
	Title       string `form:"title" json:"title" binding:"required"`
	Content     string `form:"content" json:"content" binding:"required"`
	CreatedUser uint64 `form:"created_user" json:"created_user" binding:"omitempty"`
}
