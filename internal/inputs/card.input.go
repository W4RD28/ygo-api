package inputs

type CardAddInput struct {
	Name        string `json:"name" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Race        string `json:"race" binding:"required"`
	Attribute   string `json:"attribute"`
	Level       int    `json:"level"`
	Attack      int    `json:"attack"`
	Defense     int    `json:"defense"`
	Description string `json:"description" binding:"required"`

	ImageID uint `json:"image_id"`
}

type CardUpdateInput struct {
	ID          uint   `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Race        string `json:"race" binding:"required"`
	Attribute   string `json:"attribute"`
	Level       int    `json:"level"`
	Attack      int    `json:"attack"`
	Defense     int    `json:"defense"`
	Description string `json:"description" binding:"required"`

	ImageID uint `json:"image_id"`
}

type CardDeleteInput struct {
	ID uint `json:"id" binding:"required"`
}
