package cmd

import (
	"fmt"

	"github.com/sarvsav/iza/models"
	"github.com/spf13/cobra"
)

func WithTouchArgs(args []string) models.OptionsTouchFunc {
	return func(c *models.TouchOptions) error { c.Args = args; return nil }
}

// touchCmd represents the touch command
var touchCmd = &cobra.Command{
	Use:   "touch",
	Short: "Creates an empty collection in your database",
	Long: `Usage: iza touch DATABASE/COLLECTION...

If value for the database is empty, then it will be added to test database.
If the database doesn't exist, then it will be created.
If the collection already exists, then it will not be created or modified.
You can provide multiple arguments to create multiple collections at once.

For example:
  iza touch demoDb/demoCollection01 testCollection02 sampleDb/sampleCollection03

It will create three collections:
 1. demoCollection01 in the demoDb,
 2. testCollection02 in the test database,
 3. and sampleCollection03 in the sampleDb.`,
	Run: func(cmd *cobra.Command, args []string) {
		service, err := cmd.Flags().GetString("service")
		if err != nil {
			fmt.Println(err)
			return
		}
		switch service {
		case "cicd":
			result, err := application.CiCdService.Touch()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(result)
		case "datastore":
			result, err := application.DataStoreService.Touch()
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
	rootCmd.AddCommand(touchCmd)
}
