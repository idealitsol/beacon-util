package util

import (
	"fmt"
	"testing"
	"time"
)

func TestGetDuration(t *testing.T) {
	dul := GetDuration(2, Week)
	dateString1 := "2018-09-23 19:47:59.724+00:00"
	dateString2 := "2018-09-26 19:47:59.724+00:00"
	time1, err := time.Parse(time.RFC3339, dateString1)
	time2, err := time.Parse(time.RFC3339, dateString2)
	if err != nil {
		fmt.Println("Error while parsing date :", err)
	}
	t.Error(dul)
	t.Error(time2.Sub(time1).Nanoseconds() / int64(time.Millisecond))
}
