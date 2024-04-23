package watchmaker

import "time"

type Timer struct {
	BaseInterval time.Time
	SetOnTheHour bool
	Delay        time.Time
	// 型なんとかする
	Timezone string
}

type BreakTime struct {
	From time.Time
	To   time.Time
}
