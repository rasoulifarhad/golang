package bcrypttest

import (
	"fmt"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestDo(t *testing.T) {
	rawPassword := "some-pw"
	pepper := "some-pepper"

	hashedBytes, err := bcrypt.GenerateFromPassword(
		[]byte(rawPassword+pepper), bcrypt.DefaultCost)
	if err != nil {
		// This really shouldn't happen unless we provide an invalid bcrypt cost or
		// are doing something else very wrong
	}

	// Use the hashedBytes, or convert to a string
	hashedString := string(hashedBytes)
	fmt.Printf("hasshed pass: %v\n", hashedString)

	err = bcrypt.CompareHashAndPassword(
		[]byte(hashedString),
		[]byte(rawPassword+pepper))
	if err != nil {
		// Passwords don't match - tell the user the login was invalid!
		fmt.Println("Passwords don't match - tell the user the login was invalid!")
	}
	// The passwords match if no error!
	fmt.Println("The passwords match if no error!")

}

func TestBCrypt(t *testing.T) {
	bc := NewBCrypt("some-pepper")
	rawPassword := "some-pw"
	hashedPW, err := bc.hash(rawPassword)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("hasshed pass: %v\n", hashedPW)
	}

	if bc.Equal(hashedPW, rawPassword) {
		fmt.Println("The passwords match if no error!")

	}
}
