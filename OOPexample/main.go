package main

import (
	"./models"
)

func main()  {
	alarm := models.NewHeadAlarm("Chicken", "rock")
	alarm.SetAlarmSound("Queen")
	alarmPlays(alarm)
}

func alarmPlays(alarm models.Alarm)  {
	alarm.Play()
}
