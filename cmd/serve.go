package cmd

import (
	"github.com/Tech-Arch1tect/resume-go/lib"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		lib.Serve(cmd.Flags().Lookup("file").Value.String(), cmd.Flags().Lookup("port").Value.String())
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringP("file", "f", "resume.html", "file to serve")
	serveCmd.Flags().StringP("port", "p", "8080", "port to serve the resume on")
}
