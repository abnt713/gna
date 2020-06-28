package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/abnt713/bullseye/hero"
)

func main() {
	fmt.Println("!!! Welcome to the hero control !!!")

	reader := bufio.NewReader(os.Stdin)

	quiver, err := hero.NewQuiver(1, 1)
	if err != nil {
		panic(err)
	}

	myHero := hero.NewHero(
		hero.NewStdoutMind(),
		hero.NewBody(),
		quiver,
		hero.NewBow(),
	)

	running := true
	for running {
		instruction, err := reader.ReadString('\n')
		instruction = strings.Replace(instruction, "\n", "", -1)

		if err != nil {
			panic(err)
		}
		switch instruction {
		case "up":
			myHero.MoveUp()
		case "down":
			myHero.MoveDown()
		case "pull":
			myHero.PullBow()
		case "ready":
			myHero.ReadyArrow()
		case "shoot":
			myHero.ShootArrow()
		case "quit":
			running = false
		default:
			fmt.Printf("Unknown command: %v\n", instruction)
		}
	}

	fmt.Println("See ya :D")
}
