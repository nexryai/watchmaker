package watchmaker

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	tz, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		t.Fatal(err)
	}

	timer := Timer{
		BaseInterval: 10 * time.Second,
		RunOnTheHour: false,
		Delay:        5 * time.Second,
		Timezone:     tz,
	}

	startedAt := time.Now()
	timer.WaitForNextScheduledTime()
	finishedAt := time.Now()

	if finishedAt.Sub(startedAt) < 15*time.Second {
		t.Errorf("Timer didn't wait for the next scheduled time")
	} else if finishedAt.Sub(startedAt) > 17*time.Second {
		t.Errorf("Timer waited too long")
	}
}

func testRunOnTheHour(t *testing.T) {
	tz, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		t.Fatal(err)
	}

	timer := Timer{
		BaseInterval: 10 * time.Second,
		RunOnTheHour: true,
		Delay:        10 * time.Second,
		Timezone:     tz,
	}

	startedAt := time.Now()
	timer.WaitForNextScheduledTime()
	finishedAt := time.Now()

	t.Logf("Started at: " + startedAt.Format("2006-01-02 15:04:05"))
	t.Logf("Finished at: " + finishedAt.Format("2006-01-02 15:04:05"))

	if !strings.HasSuffix(finishedAt.Format("2006-01-02 15:04:05"), "0") {
		t.Errorf("Timer didn't run on the hour: " + finishedAt.Format("2006-01-02 15:04:05"))
	} else {
		t.Logf("Timer ran on the hour: " + finishedAt.Format("2006-01-02 15:04:05"))
	}

	if finishedAt.Sub(startedAt) > 10*time.Second {
		t.Errorf("Timer waited too long")
	}
}

func TestOnTheHourManyTimes(t *testing.T) {
	for i := 1; i < 20; i++ {
		testRunOnTheHour(t)
		// 1-10秒ランダム待機
		time.Sleep(time.Duration(1+rand.Intn(9)) * time.Second)
	}
}

func TestBreakTime(t *testing.T) {
	tz, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatal(err)
	}

	timer := Timer{
		BaseInterval: 10 * time.Second,
		RunOnTheHour: false,
		Timezone:     tz,
		BreakTimes: []BreakTime{
			{
				From:  time.Now().In(tz),
				Until: time.Now().In(tz).Add(20 * time.Second),
			},
		},
	}

	startedAt := time.Now()
	timer.WaitForNextScheduledTime()
	finishedAt := time.Now()

	t.Logf("Started at: " + startedAt.Format("2006-01-02 15:04:05"))
	t.Logf("Finished at: " + finishedAt.Format("2006-01-02 15:04:05"))

	if finishedAt.Sub(startedAt) < 20*time.Second {
		t.Errorf("Timer didn't wait for the next scheduled time")
	} else if finishedAt.Sub(startedAt) > 21*time.Second {
		t.Errorf("Timer waited too long")
	}
}
