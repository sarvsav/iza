package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/sarvsav/iza/models"
	"github.com/spf13/cobra"
)

func WithCatArgs(args []string) models.OptionsCatFunc {
	return func(c *models.CatOptions) error { c.Args = args; return nil }
}

func printPrettyJSON(results models.CatResponse) {
	green := color.New(color.FgGreen).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	fmt.Println(green("Total number of Documents:"), results.Count)
	for i, doc := range results.Documents {
		fmt.Println(green("Document:"), i+1)
		for k, v := range doc {
			fmt.Printf("  %s: %v\n", cyan(k), v)
		}
		fmt.Println()
	}
}

// catCmd represents the cat command
var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "Concatenate Document(s) to standard output from collection(s).",
	Long: `Usage: iza cat [options] DATABASE/COLLECTION...

It will read the document(s) from the specified collection(s) and display them
on the standard output from each collection. If there is no database name
provided, then it will search for the collection in the test database. If the
database or collection does not exist, it will return empty result, and nothing
will be displayed. For example:

  iza cat demoDb/demoCollection01 testCollection02 sampleDb/sampleCollection03

It will display the contents of the documents from the following collections:
  1. demoCollection01 in the demoDb,
  2. testCollection02 in the test database,
  3. and sampleCollection03 in the sampleDb.

You can provide multiple arguments to read documents from multiple collections at once.`,
	Run: func(cmd *cobra.Command, args []string) {
		service, err := cmd.Flags().GetString("service")
		if err != nil {
			fmt.Println(err)
			return
		}
		switch service {
		case "cicd":
			err := application.CiCdService.Cat()
			if err != nil {
				fmt.Println(err)
				return
			}
		case "datastore":
			result, err := application.DataStoreService.Cat(WithCatArgs(args))
			if err != nil {
				fmt.Println(err)
				return
			}
			printPrettyJSON(result)
		default:
			fmt.Println("Service not found")
		}
	},
}

func init() {
	rootCmd.AddCommand(catCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// catCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// catCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
