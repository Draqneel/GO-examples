package models

import (
	"time"
)

// Alarm - public interface
type Alarm interface {
	SetTime(time.Time)
	SetAlarmSound(string)
	Play()
}
