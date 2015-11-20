package main

import (
	"bufio"
	"fmt"
	"github.com/howeyc/gopass"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"os"
)

// This is a program to generate user-password files
// Passwords are encrypted with bcrypt, and the username
// and password are concatenated by a ':'.
func main() {
	// Get desired username
	fmt.Printf("Please enter a username: ")
	consoleReader := bufio.NewReader(os.Stdin)
	user, err := consoleReader.ReadString('\n')
	if err != nil {
		log.Println(err)
		return
	}
	user = user[:len(user)-1] // Trim newline character

	// Get desired password
	fmt.Printf("Please enter a password: ")
	password := gopass.GetPasswd() // gopass.GetPasswd() does not echo back to the user

	// Create hash of password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return
	}

	// Write username and password to `user.pass`
	// Format: `user:password`
	err = ioutil.WriteFile("user.pass", []byte(user+":"+string(hashedPassword)), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
}
