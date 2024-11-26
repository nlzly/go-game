package main

import (
	"fmt"
	"go-game/color"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	var input string
	fmt.Println()
	iterator := 0
	for iterator < 1 {
		fmt.Scanln(&input)
		color.PrintGreen(input)
		checkInput(input)
	}
}

func checkInput(input string) {
	if input == "clear" {
		clearScreen()
	}
	if input == "exit" {
		exitApp()
	}
}

func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func exitApp() {
	os.Exit(1)
}
