package models

import (
	"time"
)

type Alarm interface {
	SetTime(time.Time)
	SetAlarmSound(string)
	Play()
}
