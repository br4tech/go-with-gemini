package domain

type Opinion struct {
	Content   string `validate:"required"`
	ProductID int    `validate:"required"`
}

func NewOpinion(content string, productID int) *Opinion {
	return &Opinion{
		Content:   content,
		ProductID: productID,
	}
}
