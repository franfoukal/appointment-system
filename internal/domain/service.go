package domain

import (
	"net/url"
	"time"
)

type (
	Service struct {
		ID          int64
		Name        string
		Duration    time.Duration
		Description string
		ImageURL    url.URL
	}
)
