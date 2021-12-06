package console

import "github.com/fatih/color"

func WriteInfo(message string) {
	c := color.New(color.FgGreen)
	c.Println(message)
}

func WriteError(message string) {
	c := color.New(color.FgRed)
	c.Println(message)
}

func WriteWarning(message string) {
	c := color.New(color.FgYellow)
	c.Println(message)
}
