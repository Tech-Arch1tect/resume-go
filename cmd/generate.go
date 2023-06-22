package cmd

import (
	"log"

	"github.com/Tech-Arch1tect/resume-go/lib"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate --resume <resume.json> --template <template> --output <resume.html>",
	Short: "The generate command generates a HTML resume from a JSON file and a template",
	Long: `The generate command generates a HTML resume from a JSON file and a template.
	
	Example:
	resume-go generate --resume resume.json --template ./my-theme --output resume.html`,
	Run: func(cmd *cobra.Command, args []string) {
		json := lib.GetMapFromJson(cmd.Flag("resume").Value.String())
		lib.GenerateResume(json, cmd.Flag("template").Value.String(), cmd.Flag("output").Value.String())
		log.Println("Resume generated")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP("resume", "r", "resume.json", "Path to resume.json")
	generateCmd.Flags().StringP("template", "t", "my-theme", "Path to template")
	generateCmd.Flags().StringP("output", "o", "resume.html", "Path to output file")
}
