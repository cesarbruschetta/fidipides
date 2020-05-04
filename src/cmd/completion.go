package cmd

import (
	"os"

	"./utils"
	"github.com/spf13/cobra"
)

var newCompletionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate bash completion script for fidipides cli",
	Long: `Generate bash completion code for fidipides
To use it:

	$ source <(fidipides completion)

Then you'll have completion for all the basic fidipides commands. You may want
to add that line to your ~/.bash_profile or ~/.bashrc so that it runs
automatically after you login.

Make sure to have the bash-completion package installed. If you're on OS X also
pay attention to your bash version: it must be >= 4.0, which can be installed
by brew/ports and changed on your terminal emulator of preference.
    `,
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Root().GenBashCompletion(os.Stdout); err != nil {
			utils.PrintErrorAndExit(err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(newCompletionCmd)
}
