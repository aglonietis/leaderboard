package requests

type ScoreStoreRequest struct {
	Name string `json:"name" validate:"required,max=191"`
	Score int `json:"score" validate:"required"`
}

