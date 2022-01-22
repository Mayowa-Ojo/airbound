package actors

type UserRegistrationVM struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Street    string `json:"street" binding:"required"`
	City      string `json:"city" binding:"required"`
	State     string `json:"state" binding:"required"`
	Zipcode   string `json:"zipcode" binding:"required"`
}

type UserLoginVM struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	TwoFaCode string `json:"two_fa_code"`
}

type AccountVerificationVM struct {
	Token string `form:"tk" binding:"required"`
}

type TwoFaSetupVM struct {
	UserID    string `json:"user_id" binding:"required"`
	TwoFaCode string `json:"two_fa_code"`
}
