package validator

type CreatePost struct {
	Title       string `form:"title" binding:"required"`
	Description string `form:"description" binding:"required"`
	Slug        string `form:"slug"`
	UserID      string `form:"user_id"`
}
