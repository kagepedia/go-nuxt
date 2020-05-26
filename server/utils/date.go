package utils

import (
	"time"
)

func Datetime(datetime time.Time) string {
	const layout = "2006-01-02 15:04:05"
	return datetime.Format(layout)
}
