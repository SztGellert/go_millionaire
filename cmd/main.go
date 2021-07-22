package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Become a millionaire!")
	time.Sleep(2)
	fmt.Println("What is your name?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	fmt.Println(fmt.Sprintf("Hi %s!", text))
	fmt.Println("How much hours is in a day?")
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	answer := scanner.Text()
	if answer == "24" {
		fmt.Println("Good answer!")
		time.Sleep(2)
		fmt.Println("Let's see the next question!")
	} else {
		fmt.Println("Wrong answer!")
		time.Sleep(2)
		fmt.Println("Bye")
	}
}
