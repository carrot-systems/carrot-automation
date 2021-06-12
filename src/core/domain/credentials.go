package domain

type Credentials struct {
	Password string `json:"password" binding:"required" form:"password"`
}
