package auth

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email         string `json:"email"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	DisplayName   string `json:"display_name"`
	PhoneNumber   string `json:"phone_number"`
	AvatarURL     string `json:"avatar_url"`
	CheckPassword string `json:"check_password"`
}
