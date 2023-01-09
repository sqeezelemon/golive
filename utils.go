package golive

import (
	"strconv"
	"strings"
	"time"
)

////// ERRORS

type ApiError int

func (e ApiError) Error() string {
	var description string
	switch e {
	case 0:
		description = "OK"
	case 1:
		description = "User not found"
	case 2:
		description = "Missing request parameters"
	case 3:
		description = "Endpoint error"
	case 4:
		description = "Not authorised (check API key)"
	case 5:
		description = "Server not found"
	case 6:
		description = "Flight not found"
	case 7:
		description = "No ATIS available"
	default:
		description = "Undocumented error code"
	}
	return "Live API error " + strconv.Itoa(int(e)) + ": " + description
}

////// TIME

const (
	layoutWithoutT = "2006-01-02 15:04:05Z"
)

type TimeWithoutT time.Time

func (t *TimeWithoutT) UnmarshalJSON(b []byte) error {
	time, err := time.Parse(layoutWithoutT, strings.Trim(string(b), `"`))
	if err != nil {
		return err
	}
	*t = TimeWithoutT(time)
	return nil
}

func (t TimeWithoutT) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(t).Format(layoutWithoutT)), nil
}
