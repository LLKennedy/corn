package corn

import (
	"fmt"

	"github.com/spf13/cobra"
)

// CLI is a command line interface
type CLI struct {
	rootCmd *cobra.Command
	// Command is the root cobra Command for the user application
	Command *cobra.Command
}

// New creates a new CLI based on a root command
func New(cmd *cobra.Command) *CLI {
	c := &CLI{
		Command: cmd,
	}
	c.rootCmd = &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			c.getCommand().Execute()
		},
	}
	return c
}

// Run runs the CLI
func (c *CLI) Run() {
	c.getCommand().Execute()
}

func (c *CLI) getCommand() *cobra.Command {
	if c == nil || c.rootCmd == nil {
		return &cobra.Command{
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Fprintln(cmd.ErrOrStderr(), "Corn CLI: No command set")
			},
		}
	}
	return c.Command
}
