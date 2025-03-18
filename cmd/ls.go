package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// func WithLongListing(b bool) internals.OptionsLsFunc {
// 	return func(c *models.LsOptions) error { c.LongListing = b; return nil }
// }

// func WithColor(color bool) internals.OptionsLsFunc {
// 	return func(c *models.LsOptions) error { c.Color = color; return nil }
// }

// func WithArgs(args []string) internals.OptionsLsFunc {
// 	return func(c *models.LsOptions) error { c.Args = args; return nil }
// }

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List information about the databases and collections",
	Long: `Usage: iza ls [OPTIONS] [DATABASE/COLLECTION...]

List information about the databases and collections. If no arguments are
provided, it will list all databases. If a database is provided, it will list
all collections in that database. If a collection is provided, it will list
information about that collection.

For example:
  iza ls
  iza ls demoDb
  iza ls demoDb/demoCollection01
  iza ls demoDb/demoCollection01 testDb/testCollection02

It will list:
  1. all databases,
  2. all collections in demoDb,
  3. information about demoCollection01 in demoDb,
  4. and information about demoCollection01 in demoDb and testCollection02 in testDb.`,
	Run: func(cmd *cobra.Command, args []string) {
		// longListingValue, err := cmd.Flags().GetBool("long")
		// if err != nil {
		// 	panic(err)
		// }
		// colorValue, err := cmd.Flags().GetBool("color")
		// if err != nil {
		// 	panic(err)
		// }
		// internals.Ls(WithLongListing(longListingValue), WithColor(colorValue), WithArgs(args))
		service, err := cmd.Flags().GetString("service")
		if err != nil {
			fmt.Println(err)
			return
		}
		switch service {
		case "cicd":
			result, err := application.CiCdService.Ls()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(result)
		case "datastore":
			result, err := application.DataStoreService.Ls()
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
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	lsCmd.Flags().BoolP("long", "l", false, "Long listing format of databases and collections")
	lsCmd.Flags().BoolP("color", "c", false, "Add colors to the output")
}
