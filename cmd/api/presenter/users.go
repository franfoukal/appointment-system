package presenter

type (
	Registration struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Username  string `json:"username"`
		Email     string `json:"email"`
	}

	Token struct {
		JWT string `json:"token"`
	}
)
