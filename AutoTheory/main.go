package main

func main() {
	gamels := make([]string, 0)
	gamels = append(gamels, "Сансара")
	gamels = append(gamels, "Привет")
	audiols := make([]string, 0)
	audiols = append(audiols, "Зомби")
	audiols = append(audiols, "Герои")

	machine := EnjoyMachine{gameLib:&gamels,
							playlist:&audiols,
							gameCost:5,
							audioCost:5,
							coins:0}

	firstState := GreetingState{&machine}

	firstState.NextFlagState(0)
}
