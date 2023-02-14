package carValidator

type CarInput struct {
	Make        string `json:"make" validate:"required"`
	Description string `json:"description" validate:"required"`
	Year        int    `json:"year" validate:"year"`
}
