package cmd

import (
	"fmt"

	"github.com/sarvsav/iza/models"
	"github.com/spf13/cobra"
)

func WithDuArgs(args []string) models.OptionsDuFunc {
	return func(c *models.DuOptions) error { c.Args = args; return nil }
}

func prettyPrint(result models.DuResponse) {
	if result.Collection == "" {
		fmt.Println(result.Size, result.Database)
	} else {
		fmt.Println(result.Size, result.Database+"/"+result.Collection)
	}
}

// duCmd represents the du command
var duCmd = &cobra.Command{
	Use:   "du",
	Short: "A brief description of your command",
	Long: `Usage: iza du DATABASE/COLLECTION...

Prints the disk usage of the specified database or collection.
`,
	Run: func(cmd *cobra.Command, args []string) {
		service, err := cmd.Flags().GetString("service")
		if err != nil {
			fmt.Println(err)
			return
		}
		switch service {
		case "cicd":
			result, err := application.CiCdService.Du()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(result)
		case "datastore":
			result, err := application.DataStoreService.Du(WithDuArgs(args))
			if err != nil {
				fmt.Println(err)
				return
			}
			prettyPrint(result)
		default:
			fmt.Println("Service not found")
		}
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
