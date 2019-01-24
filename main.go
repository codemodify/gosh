package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// supportedCommands := []commands.Command{
	// 	commands.NewExecuteWithBash(),
	// }

	// supportedCommandsExecutor := commands.NewExecutor(supportedCommands)

	reader := bufio.NewReader(os.Stdin)
	for {
		rune, _, err := reader.ReadRune()
		fmt.Println(rune)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		switch rune {
		case 'A':
			fmt.Println("A Key Pressed")
			break
		case 'a':
			fmt.Println("a Key Pressed")
			break
		}

		// supportedCommandsExecutor.Execute(input)
	}

}
