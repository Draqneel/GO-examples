package models

import (
	"fmt"
	"math/rand"
	"time"
)

type HeadAlarm struct {
	ID    int64
	name  string
	time  time.Time
	sound string
}

func NewHeadAlarm(name string, sound string) *HeadAlarm {
	return &HeadAlarm{
		ID:    rand.Int63(),
		name:  name,
		sound: sound,}
}

func (alarm *HeadAlarm) SetTime(inTime time.Time) {
	alarm.time = inTime
}

func (alarm *HeadAlarm) SetAlarmSound(inString string) {
	alarm.sound = inString
}

func (alarm *HeadAlarm) Play() {
	fmt.Println("Playing alarm with name", alarm.name, "and ID:", alarm.ID, "sound", alarm.sound)
}
