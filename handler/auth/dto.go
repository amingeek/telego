package auth

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email         string `json:"email"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	CheckPassword string `json:"check_password"`
}
