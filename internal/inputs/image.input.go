package inputs

type UploadImageInput struct {
	File   string `form:"file" binding:"required"`
	Name   string `form:"name" binding:"required"`
	CardID uint   `form:"card_id" binding:"required"`
}

type ImageEditInput struct {
	ID     uint   `form:"id" binding:"required" json:"id"`
	Name   string `form:"name" binding:"required" json:"name"`
	CardID uint   `form:"card_id" binding:"required" json:"card_id"`
}

type ImageDeleteInput struct {
	ID uint `form:"id" binding:"required" json:"id"`
}
