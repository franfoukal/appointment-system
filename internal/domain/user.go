package domain

type (
	User struct {
		ID        uint
		FirstName string
		LastName  string
		Username  string
		Email     string
		Password  string
		Role      string
	}

	TokenPayload struct {
		ID       uint
		Username string
		Email    string
		Role     string
	}
)
