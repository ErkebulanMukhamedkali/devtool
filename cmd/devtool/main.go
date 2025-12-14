package main

import (
	"fmt"
	"os"

	"github.com/ErkebulanMukhamedkali/devtool/internal/commands"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "init":
		if len(os.Args) > 2 {
			fmt.Println("init have no extra arguments")
			os.Exit(1)
		}

		if err := commands.Init(); err != nil {
			fmt.Println(err)
		}
	case "add", "remove":
		if len(os.Args) < 3 {
			fmt.Println("not enough arguments")
			os.Exit(1)
		}

		if len(os.Args) > 3 {
			fmt.Println("only one argument is expected")
			os.Exit(1)
		}

		switch os.Args[1] {
		case "add":
			if err := commands.Add(os.Args[2]); err != nil {
				fmt.Println(err)
			}
		case "remove":
			if err := commands.Remove(os.Args[2]); err != nil {
				fmt.Println(err)
			}
		}
	case "list":
		if len(os.Args) > 2 {
			fmt.Println("list have no extra arguments")
			os.Exit(1)
		}

		if err := commands.List(); err != nil {
			fmt.Println(err)
		}
	case "run":
		if len(os.Args) < 3 {
			fmt.Println("not enough arguments")
			os.Exit(1)
		}

		if err := commands.Run(os.Args[2:]); err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Printf("unknown command %q\n", os.Args[1])
		os.Exit(1)
	}
}
