package datemanager

import "time"

type CustomTime struct {
	time.Time
}

func (c CustomTime) String() string {
	return c.Time.Format(time.RFC3339)
}

func Parse(value string) (*CustomTime, error) {
	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return nil, err
	}

	return &CustomTime{t}, nil
}
