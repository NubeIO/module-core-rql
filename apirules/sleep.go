package apirules

import "time"

// Sleep will delay the program for the `duration` passed in (duration is units seconds)
func (inst *RQL) Sleep(duration int) {
	d := time.Duration(duration)
	time.Sleep(d * time.Second)
}
