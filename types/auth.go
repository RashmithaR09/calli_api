package types

type LoginRequest struct {
	Email       string `form:"email" binding:"required"`
	Password    string `form:"password" binding:"required"`
	RedirectURI string `form:"redirect_uri" binding:"required"`
	Scope       string `form:"scope" binding:"required"`
}
