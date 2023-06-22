package cmd

import (
	"github.com/Tech-Arch1tect/resume-go/lib"
	"github.com/spf13/cobra"
)

// watchCmd represents the watch command
var watchCmd = &cobra.Command{
	Use:   "watch --resume <resume.json> --template <template> --output <resume.html>",
	Short: "watch: watch for changes in the template and re-generates the resume (optionally serves it)",
	Long: `watch: watch for changes in the template and re-generates the resume (optionally serves it)
	
	Example:
	resume-go watch --resume resume.json --template ./my-theme --output resume.html
	
	Optionally serve the resume on a local webserver:
	resume-go watch --resume resume.json --template ./my-theme --output resume.html --serve true --port 8080`,
	Run: func(cmd *cobra.Command, args []string) {
		// generate the resume initially
		lib.GenerateResume(lib.GetMapFromJson(cmd.Flags().Lookup("resume").Value.String()), cmd.Flags().Lookup("template").Value.String(), cmd.Flags().Lookup("output").Value.String())
		// watch for changes in the template and re-generate the resume
		// optionally serve the resume on a local webserver
		if cmd.Flags().Lookup("serve").Value.String() == "true" {
			// run the watch command in a goroutine if the serve flag is set
			go lib.Watch(cmd.Flags().Lookup("template").Value.String(), cmd)
			lib.Serve(cmd.Flags().Lookup("output").Value.String(), cmd.Flags().Lookup("port").Value.String())
		} else {
			lib.Watch(cmd.Flags().Lookup("template").Value.String(), cmd)
		}
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
	watchCmd.Flags().StringP("resume", "r", "resume.json", "Path to resume.json")
	watchCmd.Flags().StringP("template", "t", "my-theme", "Path to template")
	watchCmd.Flags().StringP("output", "o", "resume.html", "Path to output file")
	watchCmd.Flags().BoolP("serve", "s", false, "serve the resume on a local webserver")
	watchCmd.Flags().StringP("port", "p", "8080", "port to serve the resume on")
}
