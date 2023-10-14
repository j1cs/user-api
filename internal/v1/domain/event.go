package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Header struct {
	EventId   string `json:"eventId,omitempty"`
	TraceId   string `json:"traceId,omitempty"`
	Datetime  string `json:"datetime,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
}

func NewHeader() Header {
	guid := uuid.NewString()
	currentTime := time.Now()
	return Header{
		EventId:   guid,
		TraceId:   guid,
		Datetime:  currentTime.Format("2006-01-02T15:04:05.000Z"),
		Timestamp: fmt.Sprint(currentTime.Unix()),
	}
}
