package utils

import (
	"time"
)

func CurrentNanos() int64 {
	return time.Now().UnixNano()
}

func FromTimeInNanos(fromTimeNanos int64) int64 {
	return time.Now().UnixNano() - fromTimeNanos
}
