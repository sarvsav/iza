package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// whoamiCmd represents the whoami command
var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Prints the current user",
	Long: `Usage: iza whoami

Prints the current logged in user.`,
	Run: func(cmd *cobra.Command, args []string) {
		service, err := cmd.Flags().GetString("service")
		if err != nil {
			fmt.Println(err)
			return
		}
		switch service {
		case "cicd":
			result, err := application.CiCdService.WhoAmI()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(result)
		case "datastore":
			result, err := application.DataStoreService.WhoAmI()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(result)
		default:
			fmt.Println("Service not found")
		}
	},
}

func init() {
	rootCmd.AddCommand(whoamiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// whoamiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// whoamiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
