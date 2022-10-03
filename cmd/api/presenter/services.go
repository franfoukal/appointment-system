package presenter

type (
	Service struct {
		ID          int64  `json:"id"`
		Name        string `json:"name"`
		Duration    int64  `json:"duration"`
		Description string `json:"description"`
		ImageURL    string `json:"image"`
	}
)
