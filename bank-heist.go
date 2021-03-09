package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	isHeistOn := true
	eludedGuards := rand.Intn(100)

	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time?")
	}

	openedVault := rand.Intn(100)
	if isHeistOn == true && openedVault >= 70 {
		fmt.Println("Grab and GO!")
	} else if isHeistOn == false {
		fmt.Println("Vault can't be opened...")
	}

	fmt.Println("State of Heist: ", isHeistOn)

	leftSafely := rand.Intn(5)
	if isHeistOn == true {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("Plan a better heist...")
		case 1:
			isHeistOn = false
			fmt.Println("Turns out vault doors don't open from the inside...")
		case 2:
			isHeistOn = false
			fmt.Println("Plan a better heist...")
		case 3:
			isHeistOn = false
			fmt.Println("Ghen ghen...")
		case 4:
			isHeistOn = false
			fmt.Println("What is going on?")
		case 5:
			isHeistOn = false
			fmt.Println("This wasnt what we planned...")
		default:
			fmt.Println("Start the getaway car")
		}

		if isHeistOn == false {
			amtStolen := 10000 + rand.Intn(1000000)
			fmt.Println("Total amount stolen:", "$", amtStolen)
		}
	}

}
