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
	theUser = MakeNewUser(dataString)
	fmt.Print("Hello " + theUser.Name + ". ")
}

func prompt() {
	fmt.Print("What would you like to know?\n")
	promptString := magnaio.ReadInput()
	fmt.Println(promptString)
	str := igniteconnect.FormatQuery(promptString, &theUser)
	igniteconnect.SendQuery(str)
}

func MakeNewUser(aName string) magnauser.User {
	newUser := magnauser.User{aName, 1223124312342}
	return newUser
}

func main() {
	//igniteconnect.ConnectToMagna()
	intro()
	prompt()
}
