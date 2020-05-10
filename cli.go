package corn

import "github.com/spf13/cobra"

// CLI is a command line interface
type CLI struct {
	// Command is the root cobra Command
	Command *cobra.Command
}

// New creates a new CLI based on a root command
func New(cmd *cobra.Command) *CLI {
	c := &CLI{
		Command: cmd,
	}
	return c
}

// Run runs the CLI
func (c *CLI) Run() {

}
