package requests

type ScoreStoreRequest struct {
	Name  string `json:"name" validate:"required,min=1,max=191"`
	Score int    `json:"score" validate:"required"`
}
