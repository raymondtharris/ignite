package ignite_test

import (
	"fmt"
	"testing"
)

func TestNewUser(t *testing.T) {
	users := []string{
		"Johnny Cool",
		"Johnny Bravo",
		"Max Rager",
		"Cole Maxin",
		"Ronnie James",
		"Allen Wake",
		"Zack Thack",
	}
	for _, name := range users {
		fmt.Println(name)
	}
}
