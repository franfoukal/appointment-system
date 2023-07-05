package datemanager

import (
	"errors"
	"fmt"
	"time"
)

type dateOperation struct {
	originalDate time.Time
	years        int
	months       int
	days         int
	hours        int
	minutes      int
	seconds      int
}

type Operation func(*dateOperation)

const (
	DateLayout     = "2006-01-02"
	DatetimeLayout = "2006-01-02 15:04:05"
)

var ErrParsingDateFormat = errors.New("error parsing date string to time, check format")

// Return today time at the given hour, in UTC timezone
func BuildDateWithGivenHour(date time.Time, hourString string) (*time.Time, error) {
	currentDateStr := date.Format(DateLayout)
	timeStr := fmt.Sprintf("%s %s", currentDateStr, hourString)

	t, err := time.Parse(DatetimeLayout, timeStr)
	if err != nil {
		return nil, fmt.Errorf("%w. Cause: %s", ErrParsingDateFormat, err.Error())
	}

	return &t, nil
}

// Check if a given datetime is within a time range
func IsBetween(datetime, start, end time.Time) bool {
	return (datetime.After(start) || datetime.Equal(start)) && (datetime.Before(end) || datetime.Equal(end))
}

func OperateWithDatetime(originalTime time.Time, operations ...Operation) time.Time {
	dateOperation := &dateOperation{
		originalDate: originalTime,
	}

	for _, opt := range operations {
		opt(dateOperation)
	}

	return originalTime.
		AddDate(dateOperation.years, dateOperation.months, dateOperation.days).
		Add(time.Duration(dateOperation.hours)*time.Hour + time.Duration(dateOperation.minutes)*time.Minute + time.Duration(dateOperation.seconds)*time.Second)
}

func AddMonths(months int) Operation {
	return func(do *dateOperation) {
		do.months += months
	}
}

func AddYears(years int) Operation {
	return func(do *dateOperation) {
		do.years += years
	}
}

func AddDays(days int) Operation {
	return func(do *dateOperation) {
		do.days += days
	}
}

func AddHours(hours int) Operation {
	return func(do *dateOperation) {
		do.hours += hours
	}
}

func AddMinutes(minutes int) Operation {
	return func(do *dateOperation) {
		do.minutes += minutes
	}
}

func AddSeconds(seconds int) Operation {
	return func(do *dateOperation) {
		do.seconds += seconds
	}
}

func SubtractYears(years int) Operation {
	return func(do *dateOperation) {
		do.years -= years
	}
}

func SubtractDays(days int) Operation {
	return func(do *dateOperation) {
		do.days -= days
	}
}

func SubtractMonths(months int) Operation {
	return func(do *dateOperation) {
		do.months -= months
	}
}

func SubtractHours(hours int) Operation {
	return func(do *dateOperation) {
		do.hours -= hours
	}
}

func SubtractMinutes(minutes int) Operation {
	return func(do *dateOperation) {
		do.minutes -= minutes
	}
}

func SubtractSeconds(seconds int) Operation {
	return func(do *dateOperation) {
		do.seconds -= seconds
	}
}
