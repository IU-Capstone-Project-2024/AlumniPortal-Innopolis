package helpers

import "time"

var layout = "2006-01-02 15:04:26"

func Convert(input string) time.Time {
	date, _ := time.Parse(layout, input)
	return date
}
