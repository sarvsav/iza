package cmd

import (
	"github.com/sarvsav/iza/internals"
	"github.com/sarvsav/iza/models"
	"github.com/spf13/cobra"
)

func WithDuArgs(args []string) internals.OptionsDuFunc {
	return func(c *models.DuOptions) error { c.Args = args; return nil }
}

// duCmd represents the du command
var duCmd = &cobra.Command{
	Use:   "du",
	Short: "A brief description of your command",
	Long: `Usage: iza du DATABASE/COLLECTION...

Prints the disk usage of the specified database or collection.
`,
	Run: func(cmd *cobra.Command, args []string) {
		internals.Du(WithDuArgs(args))
	},
}

func init() {
	rootCmd.AddCommand(duCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// duCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// duCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
