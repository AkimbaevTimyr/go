package requests

type LoginRequest struct {
	Email    string `json:"email" validate:"required,min=10,max=50"`
	Password string `json:"password" validate:"required,min=8,max=20"`
}

type CheckCodeRequest struct {
	Email string `json:"email"`
	Code  int    `json:"code"`
}
