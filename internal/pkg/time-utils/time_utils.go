package timeutils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidTimeFormat = errors.New("invalid time format")

// ConvertHourMinuteStrToMinutes convert from hh:mm (24 hours) to the number minutes from the start of the day (0:00)
func ConvertHourMinuteStrToMinutes(t string) (int, error) {
	part := strings.Split(t, ":")
	if len(part) != 2 {
		return 0, ErrInvalidTimeFormat
	}

	hh := part[0]
	hours, err := strconv.Atoi(hh)
	if err != nil {
		return 0, ErrInvalidTimeFormat
	}
	if hours < 0 || hours > 23 {
		return 0, ErrInvalidTimeFormat
	}

	mm := part[1]
	minutes, err := strconv.Atoi(mm)
	if err != nil {
		return 0, ErrInvalidTimeFormat
	}
	if minutes < 0 || minutes > 59 {
		return 0, ErrInvalidTimeFormat
	}

	return hours*60 + minutes, nil
}

// ConvertMinutesToHourMinutesStr convert from number minutes from the start of the day (0:00) to hh:mm (24 hours)
func ConvertMinutesToHourMinutesStr(minutes int) string {
	hours := minutes / 60
	minutes = minutes % 60

	return fmt.Sprintf("%02d:%02d", hours, minutes)
}
