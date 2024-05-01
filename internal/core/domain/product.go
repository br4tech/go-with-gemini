package domain

type Product struct {
	Name  string `validate:"required"`
	Code  string `validate:"required"`
	Image string
}

func NewProduct(name string, code string, image string) *Product {
	return &Product{
		Name:  name,
		Code:  code,
		Image: image,
	}
}
