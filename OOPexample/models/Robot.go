package models

type Robot struct {
	x int
	y int
}

func (robot *Robot) GetLocation() (int, int) {
	return robot.x, robot.y
}

func NewRobot(X int, Y int) *Robot {
	return &Robot{
		x: X,
		y: Y,
	}
}
