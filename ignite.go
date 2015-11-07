package main

import (
	"fmt"

	"github.com/ignite/igniteconnect"
	"github.com/magna/magnaio"
	"github.com/magna/magnauser"
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
	str := igniteconnect.FormatQuery(promptString, &theUser)
	igniteconnect.SendQuery(str)
}

func main() {
	//igniteconnect.ConnectToMagna()
	intro()
	prompt()
}
