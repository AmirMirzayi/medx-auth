package auth

type LoginRequestBody struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponseBody struct {
	Token string `json:"token"`
}
