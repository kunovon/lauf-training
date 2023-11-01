package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//var tage int
	var kilometer int
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\033[H\033[2J")
	fmt.Print("Bitte geben Sie die gewünschte Kilometerzahl ein: ")
	userInput, _ := reader.ReadString('\n')

	kilometer, err := strconv.Atoi(userInput)
	if err != nil {
		fmt.Println("Error during conversion")
		return
	}

	//userInput = kilometer
	fmt.Print("\033[H\033[2J")
	fmt.Println("Ihre Eingabe ist:", kilometer)
}
