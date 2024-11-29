package repl

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var replName string = "MonkeyREPL"

func printPrompt() {
	fmt.Print(replName, ">> ")
}

func printUnknown(command string) {
	fmt.Println(command, ": command not found")
}

func printHelp() {
	fmt.Printf("Listed are available commands for this %v: \n", replName)

	fmt.Println("/help	-	Show available commands")
	fmt.Println("/clear	-	Clear terminal")
	fmt.Println("/exit	-	Exit ", replName)
}

func clearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()

	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

}

func handleInvalidCmd(command string) {
	defer printUnknown(command)
}

func handleCmd(command string) {
	handleInvalidCmd(command)
}

func Start() {
	commands := map[string]interface{}{
		"/help":  printHelp,
		"/clear": clearScreen,
	}
	fmt.Printf("This is the %v. \n Type '/help' to get a list of available commands \n", replName)
	r := bufio.NewScanner(os.Stdin)
	printPrompt()
	for r.Scan() {
		input := r.Text()
		if command, exists := commands[input]; exists {
			command.(func())()

		} else if strings.EqualFold("/exit", input) {
			return
		} else {
			handleCmd(input)
		}
		printPrompt()
	}

}
