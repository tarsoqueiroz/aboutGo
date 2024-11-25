package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	//
	// Bcrypt: The Classic Password Smoothie
	//
	fmt.Printf("===== Bcrypt: The Classic Password Smoothie\n=====\n")

	password := []byte("iLoveCrypto123")

	// let's blend this password
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic("===== Oops, our cryptographic blender broke!!!")
	}

	// password smoothie
	fmt.Printf("===== Bcrypt: our password smoothie: %x\n", hashedPassword)

	// now, let's see if we can recognize our original password
	err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		fmt.Printf("===== Bcrypt: nope, that's not our password!!!\n=====\n")
	} else {
		fmt.Printf("===== Bcrypt: yep, that's our password alright!!!\n=====\n")
	}
}
