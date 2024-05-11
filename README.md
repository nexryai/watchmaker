# watchmaker
Advanced timer in golang

## Usage
### `timer.BaseInterval`
Set the base interval for the timer. 

### `timer.RunOnTheHour`
Counts BaseInterval based on the hour. For example, if the current time is 14:00:05 and BaseInterval is 10 seconds, the timer will run at 14:00:10.

### `timer.Delay`
Set the delay for the timer.

### `timer.Timezone`
Set the timezone for the timer.