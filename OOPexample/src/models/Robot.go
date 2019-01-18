package models

type robot struct {
	x int
	y int
}

func (newRobot *robot) GetLocation() (int, int) {
	return newRobot.x, newRobot.y
}

// NewRobot - return link on new created robot object
func NewRobot(X int, Y int) *robot {
	return &robot{
		x: X,
		y: Y,
	}
}
