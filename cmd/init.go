package cmd

import (
	"github.com/Tech-Arch1tect/resume-go/lib"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new resume.json file",
	Long:  `Initialize a new resume.json file from https://raw.githubusercontent.com/jsonresume/resume-schema/master/sample.resume.json`,
	Run: func(cmd *cobra.Command, args []string) {
		lib.InitResume()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
