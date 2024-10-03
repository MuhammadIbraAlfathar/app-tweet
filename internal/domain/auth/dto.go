package auth

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email,max=320"`
	Name     string `json:"name" binding:"required,max=50"`
	Password string `json:"password" binding:"required,max=72,min=8"`
	Role     string `json:"role" binding:"required,oneof=admin user guest"`
	Gender   string `json:"gender" binding:"required,oneof=male female"`
}
