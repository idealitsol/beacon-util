package util

import (
	"time"

	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// TimeToGrpcTimestamp converts golang time.Time to google protobuf timestamp.Timestamp
func TimeToGrpcTimestamp(t *time.Time) *timestamp.Timestamp {
	if t.IsZero() {
		return nil
	}
	return &timestamp.Timestamp{Seconds: t.Unix()}
}

// GrpcTimestampToTime converts google protobuf timestamp.Timestamp to golang time.Time
func GrpcTimestampToTime(t *timestamp.Timestamp) *time.Time {
	if t.GetSeconds() == 0 {
		return nil
	}

	tt := time.Unix(t.GetSeconds(), 0)
	return &tt
}
