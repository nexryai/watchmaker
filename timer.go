package watchmaker

import (
	"time"
)

type Timer struct {
	BaseInterval time.Duration
	RunOnTheHour bool
	Delay        time.Duration
	Timezone     *time.Location
	LastRun      time.Time
}

type BreakTime struct {
	From  time.Time
	Until time.Time
}

func (t *Timer) WaitForNextScheduledTime() {
	if t.Timezone == nil {
		t.Timezone = time.Local
	}

	baseTime := time.Now().In(t.Timezone)

	if t.RunOnTheHour {
		// Truncate to the hour
		// ex: 2019-01-01 12:34:56 -> 2019-01-01 12:00:00
		baseTime = baseTime.Truncate(time.Hour)
	}

	nextTime := baseTime.Add(t.BaseInterval)
	if t.Delay > 0 {
		nextTime = nextTime.Add(t.Delay)
	}

	// もし次の時間が過去の時間だったら、baseIntervalを足していく
	for {
		if nextTime.Before(time.Now().In(t.Timezone)) {
			nextTime = nextTime.Add(t.BaseInterval)
		} else {
			break
		}
	}

	// Wait until the next scheduled time
	time.Sleep(nextTime.Sub(time.Now().In(t.Timezone)))
	t.LastRun = time.Now().In(t.Timezone)

	return
}
