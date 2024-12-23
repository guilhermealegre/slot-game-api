package request

//swagger:parameters LoginRequest
type LoginRequest struct {
	// Body
	// in: body
	// required: true
	Body struct {
		// Email
		// required: true
		Email string `json:"email" validate:"required,email"`
		// Password
		// required: true
		Password string `json:"password" validate:"required,gte=6"`
	}
}

//swagger:parameters SignupRequest
type SignupRequest struct {
	// Body
	// in: body
	// required: true
	Body struct {
		// First Name
		// required: true
		FirstName string `json:"first_name" validate:"required"`
		// Last Name
		// required: true
		LastName string `json:"last_name" validate:"required"`
		// Avatar
		// required: true
		Avatar string `json:"avatar" validate:"required"`
		// Email
		// required: true
		Email string `json:"email" validate:"required,email"`
		// Password
		// required: true
		Password string `json:"password" validate:"required,gte=6"`
		// Confirm Password
		// required: true
		ConfirmPassword string `json:"confirm_password" validate:"eqfield=Password"`
	}
}
