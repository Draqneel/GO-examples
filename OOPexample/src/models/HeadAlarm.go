package models

import (
	"fmt"
	"math/rand"
	"time"
)

type headAlarm struct {
	ID    int64
	name  string
	time  time.Time
	sound string
}

// NewHeadAlarm - create new Object of headAlarm and return link
func NewHeadAlarm(name string, sound string) *headAlarm {
	return &headAlarm{
		ID:    rand.Int63(),
		name:  name,
		time:  time.Now(),
		sound: sound}
}

func (alarm *headAlarm) SetTime(inTime time.Time) {
	alarm.time = inTime
}

func (alarm *headAlarm) SetAlarmSound(inString string) {
	alarm.sound = inString
}

func (alarm *headAlarm) Play() {
	fmt.Println("Playing alarm with name", alarm.name, "in Time:", alarm.time.Format("01-02-2006 3:4:5"), "sound", alarm.sound)
}
