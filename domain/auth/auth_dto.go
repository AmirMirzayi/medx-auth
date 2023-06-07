package auth

type LoginRequestBody struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ResponseBody struct {
	Token string      `json:"token"`
	User  interface{} `json:"user"`
}
