package util

import (
	"time"
)

const (
	Minute = 1000 * 60
	Hour   = Minute * 60
	Day    = Hour * 24
	Week   = Day * 7
	// Months = Weeks
	// Year =
)

func GetDuration(value int, duration int) int {
	return value * duration
}

type BeaconTime struct {
	*time.Time
}

// func (s *BeaconTime) IsZero() bool {
// 	if s.IsZero() {
// 		return true
// 	}

// 	return false
// }
