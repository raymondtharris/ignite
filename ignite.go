package main

import (
	"fmt"
	"magna/magnaio"
	"magna/magnauser"
	"ignite/igniteconnect"
)

var theUser magnauser.User

func intro() {
	fmt.Println("Who awakens me?")
	dataString := magnaio.ReadInput()
	//fmt.Println(dataString)
	theUser = magnauser.User{dataString, 12290323094}
	fmt.Print("Hello " + theUser.Name + ". ")
}

func prompt() {
	fmt.Print("What would you like to know?\n")
	promptString := magnaio.ReadInput()
	fmt.Println(promptString)
}

func main() {
	igniteconnect.ConnectToMagna()
	intro()
	prompt()
}
