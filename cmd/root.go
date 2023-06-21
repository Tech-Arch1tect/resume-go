package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "resume-go [command] [arguments]",
	Short: "Generates a json-resume from a templat",
	Long: `Generates a json-resume from a template.
	Available commands:
	- watch: watches for changes in the template and re-generates the resume (optionally serves it)
	- generate: generates the resume from a json file and a template and outputs it to a HTML file
	- serve: serves the resume on a local webserver (does not generate the resume)`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
