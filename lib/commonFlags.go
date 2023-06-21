package lib

import (
	"github.com/spf13/cobra"
)

func RegisterCommonFlags(cmd *cobra.Command) {
	cmd.Flags().String("template", "my-theme", "template to use")
	cmd.Flags().String("resume", "resume.json", "resume to use")
	cmd.Flags().String("output", "resume.html", "output file")
}
