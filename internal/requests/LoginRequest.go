package requests

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=1,max=191"`
	Password string `json:"password" validate:"required,min=1,max=191"`
}
