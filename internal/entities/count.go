package entities

import "time"

// CountHistoryEvent - score change event.
type CountHistoryEvent struct {
	Timestamp time.Time
	Delta     int16
}

// HistoryPeriod - event history period.
type HistoryPeriod string

// Possible event history periods.
const (
	HistoryPeriod3Days HistoryPeriod = "three_days"
	HistoryPeriodWeek  HistoryPeriod = "week"
)
