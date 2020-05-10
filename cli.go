package corn

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// InteractiveCLI is a command line interface
type InteractiveCLI struct {
	// Application is the application wrapped by the interactive CLI
	Application *cobra.Command
}

// New creates a new CLI based on a root command
func New(application *cobra.Command) *InteractiveCLI {
	c := &InteractiveCLI{
		Application: application,
	}
	return c
}

// Run runs the CLI
func (cli *InteractiveCLI) Run() {
	app := cli.getApp()
	if len(os.Args) > 1 {
		// TODO: initiate interactive session other than by no args?
		app.Execute()
		return
	}
	cli.InteractivePrompt(app)
}

type flagOption struct {
	*pflag.Flag
}

func (f *flagOption) String() string {
	flagString := ""
	switch {
	case f == nil:
		flagString = "Invalid (nil) flag"
	case f.Value == nil && f.DefValue == "":
		flagString = fmt.Sprintf("Name: %s; Value: (unset)", f.Name)
	case f.Value == nil:
		flagString = fmt.Sprintf("Name: %s; Value: (unset); Default: %s", f.Name, f.DefValue)
	default:
		flagString = fmt.Sprintf("Name: %s; Value: %s", f.Name, f.Value.String())
	}
	return flagString
}

// InteractivePrompt interactively prompts the user based on the command's arguments and subcommands
func (cli *InteractiveCLI) InteractivePrompt(cmd *cobra.Command) {
	switch {
	case !cmd.Runnable() && !cmd.HasSubCommands():
		fmt.Fprintln(cmd.ErrOrStderr(), ErrorMessage("cannot perform interactive prompts for command that is not runnable and has no subcommands"))
	case !cmd.HasSubCommands():
		// No subcommands, we should assume we're at the bottom of a tree and only prompt for flags
		flags := cmd.Flags()
		flagStrings := []*flagOption{{&pflag.Flag{
			Name: "Run command with current arguments",
		}}}
		flags.VisitAll(func(flag *pflag.Flag) {
			if flag != nil && !flag.Hidden {
				// We only prompt for flags that aren't hidden
				flagStrings = append(flagStrings, &flagOption{flag})
			}
		})
		prompt := promptui.Select{
			Label: "Customise flags, then run the command",
			Items: flagStrings,
		}
		prompt.Run()
	default:
		commands := []string{}
		if cmd.Runnable() {
			commands = append(commands, cmd.Use)
		}
		prompt := promptui.Select{
			Label: "Select a Command",
			Items: commands,
		}
		prompt.Run()
	}
}

func (cli *InteractiveCLI) getApp() *cobra.Command {
	if cli == nil {
		return &cobra.Command{
			Run: func(cmd *cobra.Command, args []string) {

			},
		}
	}
	return cli.Application
}
