package forms

type Testing struct {
	Name string `json:"name" validate:"required,min=3,max=10"`
}
