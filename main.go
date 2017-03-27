package main

import (
	. "fmt"
)

func computerlogic() {
	Println("computerlogic")
}

func showCardsInHands() {
	Println("showCardsInHands")
}

func showCardDetails(num string) {
	Println("showCardDetails")
}

func castOneCard(num string) {
	Println("castOneCard")
}

func throwcard() {
	showCardsInHands()
	for {
		var cmd string
		Println("Deal with your cards:")
		Scanln(&cmd)
		switch cmd {
		case "pass":
			return
		case "show":
			Println("Card number:")
			Scanln(&cmd)
			showCardDetails(cmd)
		case "cast":
			Println("Card number:")
			Scanln(&cmd)
			castOneCard(cmd)
		case "help":
			fallthrough
		default:
			Println("support cmds:[pass,show,cast]")
		}
	}
}

func mainlogic() {
	for {
		var cmd string
		Println("Your order:")
		Scanln(&cmd)
		switch cmd {
		case "next":
			computerlogic()
		case "card":
			throwcard()
		case "help":
			fallthrough
		default:
			Println("support cmds:[next,card]")
		}
	}
}

func main() {
	Println("hello world.")
	mainlogic()
}
