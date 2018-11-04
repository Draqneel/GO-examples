package main

// 1st STRUCT and this methods

type User struct {
	Name string
	Surname string
	Age int8
}

// we can change object in this method
func (user *User) getAge() int8 {
	return user.Age
}

func (user *User) agePP()  {
	user.Age++
}

// 2nd STRUCT and this methods

type lowCaseUser struct {
	name string
	age int8
}

// we can't change object in this method
func (lowUser lowCaseUser) getLowAge() int8 {
	return lowUser.age
}

func (lowUser lowCaseUser) lowAgePP()  {
	lowUser.age++
}

// EXTENDS EXAMPLE

type Robot struct {
	X int
	Y int
}

func (robot *Robot) move(Xinc int, Yinc int) {
	robot.X = robot.X + Xinc
	robot.Y = robot.Y + Yinc
}

type BattleRobot struct {
	Robot
	Cartridges int64
}

func (robot *BattleRobot) attack(count int64)  {
	robot.Cartridges = robot.Cartridges - count
}

