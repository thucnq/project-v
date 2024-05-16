package dateutils

import (
	"errors"
	"time"
)

var ErrInvalidDateFormat = errors.New("invalid date format")

const layoutDate = "02/01/2006" // dd/MM/yyyy

// GetStartTimestampOfDate convert from 0:00 of date dd/MM/yyyy to timestamp milliseconds
func GetStartTimestampOfDate(dateStr string, location *time.Location) (int64, error) {
	date, err := time.ParseInLocation(layoutDate, dateStr, location)
	if err != nil {
		return 0, ErrInvalidDateFormat
	}

	timestampMs := date.UnixMilli()
	return timestampMs, nil
}

// GetEndTimestampOfDate convert from 23:59 of date dd/MM/yyyy to timestamp milliseconds
func GetEndTimestampOfDate(dateStr string, location *time.Location) (int64, error) {
	date, err := time.ParseInLocation(layoutDate, dateStr, location)
	if err != nil {
		return 0, ErrInvalidDateFormat
	}

	timestampMs := date.Add(time.Hour * 23).Add(time.Minute * 59).UnixMilli()
	return timestampMs, nil
}

// ParseDateFromTimestamp convert timestamp milliseconds to date format dd/MM/yyyy
func ParseDateFromTimestamp(timestampMs int64) string {
	timeObj := time.UnixMilli(timestampMs)

	date := timeObj.Format(layoutDate)
	return date
}

// GetDateFromDateStr convert date dd/MM/yyyy to time.Time
func GetDateFromDateStr(timestampMs int64) time.Time {
	timeObj := time.UnixMilli(timestampMs)
	return timeObj
}
