package requests

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CheckCodeRequest struct {
	Email string `json:"email"`
	Code  int    `json:"code"`
}
