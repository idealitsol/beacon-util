package util

import (
	"time"

	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// GoTimeToGrpcTime converts golang time.Time to google protobuf timestamp.Timestamp
func GoTimeToGrpcTime(t *time.Time) *timestamp.Timestamp {
	if t == nil {
		return nil
	}

	if t.IsZero() {
		return nil
	}
	return &timestamp.Timestamp{Seconds: t.Unix()}
}

// GrpcTimeToGoTime converts google protobuf timestamp.Timestamp to golang time.Time
func GrpcTimeToGoTime(t *timestamp.Timestamp) *time.Time {
	if t.GetSeconds() == 0 {
		return nil
	}

	tt := time.Unix(t.GetSeconds(), 0)
	return &tt
}
