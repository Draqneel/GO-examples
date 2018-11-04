package main

import "fmt"

func main()  {
	// testing work with fields of public structures
	user := User{"Anton","Tikhonov",19}
	userTwo := User{"Anatoliy", "Plotnitskiy", 20}
	userThree := User{
		Age: 99,
	}

	fmt.Println(user.getAge())
	fmt.Println(userTwo.getAge())
	fmt.Println(userThree.Age)
	user.agePP()
	fmt.Println(user.getAge())
	// testing work with fields of default structures and return field of copy struct
	lowUser := lowCaseUser{"Vasya", 88}
	fmt.Println(lowUser.getLowAge())
	lowAgePP := lowUser.lowAgePP
	fmt.Println(lowUser.getLowAge()," < ", lowAgePP)

	// extends example
	 robot1 := BattleRobot{Robot{1, 2}, 100}
	 robot1.attack(50)
	 fmt.Println(robot1)
}
