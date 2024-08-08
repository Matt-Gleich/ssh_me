package cmds

import (
	"fmt"
	"io"

	"github.com/charmbracelet/ssh"
	"github.com/gleich/lumber/v2"
	"github.com/gleich/terminal/internal/output"
	"golang.org/x/term"
)

func Terminal(s ssh.Session, styles output.Styles) {
	prefix := styles.Green.Render("λ ")
	terminal := term.NewTerminal(s, prefix)

	for {
		cmd, err := terminal.ReadLine()
		if err == io.EOF {
			fmt.Fprintln(s)
			return
		}
		if err != nil {
			lumber.Error(err, "failed to process command")
			fmt.Fprintln(s, "processing command failed, closing connection")
			return
		}

		switch cmd {
		case "":
		case "exit":
			return
		case "help":
			fmt.Fprintln(s, output.Help(styles))
		case "clear", "c":
			styles.Renderer.Output().ClearScreen()
		case "workouts":
			workouts(s, styles)
		case "projects":
			projects(s, styles)
		case "games":
			games(s, styles)
		default:
			fmt.Fprintf(
				s,
				"\nInvalid command '%s'. Type `help` to see available commands.\n\n",
				cmd,
			)
		}
	}
}
