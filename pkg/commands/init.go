package commands

import "fmt"

type InitCmd struct {
	Backend string `short:"b" enum:"github,codeartifact" default:"github" help:"The backend to initialize localbrew with. This is your package repository type."`
	Auth    string `short:"a" help:"The authentication token for your selected backend."`
}

func (cmd *InitCmd) Run(globals *Globals) {
	switch cmd.Backend {
	case "github":
		fmt.Println("github")
	case "codeartifact":
		fmt.Println("codeartifact")
	}
}
