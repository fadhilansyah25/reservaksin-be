package request

import "ca-reservaksin/businesses/admin"

type Admin struct {
	Role     string `json:"role"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AdminLogin struct {
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required"`
}

func (req *Admin) ToDomain() *admin.Domain {
	return &admin.Domain{
		Role:     req.Role,
		Username: req.Username,
		Password: req.Password,
	}
}
