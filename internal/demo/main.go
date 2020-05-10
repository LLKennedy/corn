package main

import (
	"fmt"

	"github.com/llkennedy/corn"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "demo",
		Short: "A demo application",
	}
	rootCmd.AddCommand(&cobra.Command{
		Use:   "test",
		Short: "A test command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(cmd.OutOrStdout(), "Hello, world!")
		},
	})
	cli := corn.New(rootCmd)
	cli.Run()
}
