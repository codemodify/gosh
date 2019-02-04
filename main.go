package main

import (
	"fmt"
	"strings"

	"github.com/codemodify/gosh/commands"
)

func main() {

	supportedCommands := []commands.Command{
		commands.NewParallelCP(),
		commands.NewExecuteWithBash(),
	}

	supportedCommandsExecutor := commands.NewExecutor(supportedCommands)

	err := supportedCommandsExecutor.Execute("pcp /home/richard/Temp/1 /home/richard/Temp/2")
	if err != nil {
		fmt.Println(strings.Replace(err.Error(), "bash: ", "", -1))
	}

	// reader := bufio.NewReader(os.Stdin)
	// // line := ""
	// for {
	// 	// rune, _, err := reader.ReadRune()
	// 	fmt.Println("")
	// 	fmt.Println("~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~")
	// 	fmt.Printf("--> ")

	// 	line, _, err := reader.ReadLine()
	// 	if err != nil {
	// 		fmt.Fprintln(os.Stderr, err)
	// 	}

	// 	// switch rune {
	// 	// case 13:
	// 	// 	fmt.Println("ENTER Pressed")
	// 	// 	fmt.Println(line)
	// 	// 	supportedCommandsExecutor.Execute(line)
	// 	// 	break
	// 	// default:
	// 	// 	fmt.Println(rune)
	// 	// 	line = line + string(rune)
	// 	// 	break
	// 	// }

	// 	err = supportedCommandsExecutor.Execute(string(line))
	// 	if err != nil {
	// 		fmt.Println(strings.Replace(err.Error(), "bash: ", "", -1))
	// 	}
	// }
}
