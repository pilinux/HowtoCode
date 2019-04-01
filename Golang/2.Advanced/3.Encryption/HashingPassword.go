package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func main() {

	for {

		// Enter a password and generate a salted hash
		pwd := getPwd()
		hash := hashAndSalt(pwd)
		fmt.Println("Salted Hash", hash)

		// Enter the same password again and compare it with the
		// first password entered
		pwd2 := getPwd2()
		pwdMatch := comparePasswords(hash, pwd2)
		fmt.Println("Passwords Match?:", pwdMatch)
	}
}

func getPwd() []byte {

	// Prompt the user to enter a password
	fmt.Println("Enter a password")

	// Variable to store the user's input
	var pwd string

	// Read the user's input
	_, err := fmt.Scan(&pwd)
	if err != nil {
		log.Println(err)
	}

	// Return the user's input as a byte slice which will save us
	// from having to do this conversion later on
	return []byte(pwd)

}

func getPwd2() []byte {

	// Prompt the user to enter a password
	fmt.Println("Enter the password again")

	// Variable to store the user's input
	var pwd2 string

	// Read the user's input
	_, err := fmt.Scan(&pwd2)
	if err != nil {
		log.Println(err)
	}

	// Return the user's input as a byte slice which will save us
	// from having to do this conversion later on
	return []byte(pwd2)

}

func hashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)

}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {

	// Since the hashed password will be read from the DB, it will
	// be a string which  needs to be converted into a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true

}
