package main

import (
	"fmt"
	"os"
)

type State interface {
	NextFlagState(flag int)
	Off()
}

type EnjoyMachine struct {
	currentState State
	gameCost int
	audioCost int
	gameLib *[]string
	playlist *[]string
	currentPlay string
	coins int
}


type DefaultState struct {
	machine *EnjoyMachine
}

func (state *DefaultState) init() {

}

func (state *DefaultState) greeting() {
	fmt.Println("Hello, guest! You have", state.machine.coins, "coins!")
	fmt.Println("1) More coins\n2)Select song\n3)Select game")
}

func (state *DefaultState) NextState() {
	methodLockedOutput()
}

func (state *DefaultState) Off() {
	methodLockedOutput()
}

func (state *DefaultState) NextFlagState(flag int) {
	if flag == 1 {
		state.machine.currentState = &PaymentState{state.machine}
		state.machine.currentState.NextFlagState(0)
	} else if flag == 2 {
		state.machine.currentState = &AudioState{state.machine}
		state.machine.currentState.NextFlagState(0)
	} else if flag == 3 {
		state.machine.currentState = &GameState{state.machine}
		state.machine.currentState.NextFlagState(0)
	} else {
		fmt.Println("Incorrect select!")
	}
}

type PaymentState struct {
	machine *EnjoyMachine
}

func (state *PaymentState) Off() {
	nextState :=  &DefaultState{machine:state.machine}
	nextState.greeting()
	state.machine.currentState = nextState

	var inputFlag int
	fmt.Fscan(os.Stdin, &inputFlag)

	nextState.NextFlagState(inputFlag)
}

func (state *PaymentState) NextFlagState(flag int) {
	fmt.Println("Введите количество монет")
	var inputCoins int
	fmt.Fscan(os.Stdin, &inputCoins)
	nextState :=  &DefaultState{machine:state.machine}
	nextState.greeting()
	state.machine.currentState = nextState

	var inputFlag int
	fmt.Fscan(os.Stdin, &inputFlag)
	nextState.machine.coins = inputCoins
	nextState.greeting()
	nextState.NextFlagState(inputFlag)
}


type AudioState struct {
	machine *EnjoyMachine
}

func (state *AudioState) Off() {
	nextState :=  &DefaultState{machine:state.machine}
	nextState.greeting()
	state.machine.currentState = nextState

	var inputFlag int
	fmt.Fscan(os.Stdin, &inputFlag)

	nextState.NextFlagState(inputFlag)
}

func (state *AudioState) NextFlagState(flag int) {
	fmt.Println("List of all songs:")

	ls := *state.machine.playlist

	for i, val := range ls {
		print(i+1,":",val)
	}

	state.PlayMusic(ls)

	fmt.Println("Do you want listen something else? +/-")

	var inputOperation string
	fmt.Fscan(os.Stdin, &inputOperation)

	if inputOperation == "+" {
		state.PlayMusic(ls)
	} else {
		state.Off()
	}
}

func (state *AudioState) PlayMusic(ls []string) {
	if state.machine.coins < state.machine.audioCost {
		print("You have no enough money!")
		state.Off()
		return
	}

	var inputIndex int
	fmt.Fscan(os.Stdin, &inputIndex)

	fmt.Println("Enjoy machine plays ", ls[inputIndex], "song!")
	state.machine.coins -= state.machine.audioCost
}


type GameState struct {
	machine *EnjoyMachine
}

func (state *GameState) Off() {
	nextState :=  &DefaultState{machine:state.machine}
	nextState.greeting()
	state.machine.currentState = nextState

	var inputFlag int
	fmt.Fscan(os.Stdin, &inputFlag)

	nextState.NextFlagState(inputFlag)
}

func (state *GameState) NextFlagState(flag int) {
	fmt.Println("List of all games")

	ls := *state.machine.gameLib

	for i, val := range ls {
		print(i+1,":",val)
	}

	state.PlayGame(ls)

	fmt.Println("Do you want playing something else? +/-")

	var inputOperation string
	fmt.Fscan(os.Stdin, &inputOperation)

	if inputOperation == "+" {
		state.PlayGame(ls)
	} else {
		state.Off()
	}
}

func (state *GameState) PlayGame(ls []string) {
	if state.machine.coins < state.machine.audioCost {
		print("You have no enough money!")
		state.Off()
		return
	}

	var inputIndex int
	fmt.Fscan(os.Stdin, &inputIndex)

	fmt.Println("Enjoy machine plays ", ls[inputIndex], "game")
	state.machine.coins -= state.machine.audioCost
}


type GreetingState struct {
	machine *EnjoyMachine
}

func (state *GreetingState) Off() {
	methodLockedOutput()
}

func (state *GreetingState) NextFlagState(flag int) {
	nextState :=  &DefaultState{machine:state.machine}
	nextState.greeting()
	state.machine.currentState = nextState

	var inputFlag int
	fmt.Fscan(os.Stdin, &inputFlag)

	nextState.NextFlagState(inputFlag)
}

func methodLockedOutput()  {
	fmt.Println("Method locked in this state")
}

