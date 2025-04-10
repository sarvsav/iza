package cmd

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/sarvsav/iza/models"
	"github.com/spf13/cobra"
)

func WithLsLongListing(b bool) models.OptionsLsFunc {
	return func(c *models.LsOptions) error { c.LongListing = b; return nil }
}

func WithLsColor(color bool) models.OptionsLsFunc {
	return func(c *models.LsOptions) error { c.Color = color; return nil }
}

func WithLsArgs(args []string) models.OptionsLsFunc {
	return func(c *models.LsOptions) error { c.Args = args; return nil }
}

// Helper function to print each entry (database or collection) in ls -l style
func printEntry(typeChar, name, perms, owner, group string, size int64, lastModified time.Time) {
	// Format size with units (e.g., KB, MB, GB)
	sizeFormatted := formatSize(size)

	// Format last modified date
	dateFormatted := lastModified.Format("Jan 02 15:04")

	// Define a wrapper function to match the expected signature
	var colorFunc func(format string, args ...interface{})
	if typeChar == "d" {
		colorFunc = func(format string, args ...interface{}) {
			color.New(color.FgGreen).Printf(format, args...)
		}
	} else {
		colorFunc = func(format string, args ...interface{}) {
			color.New(color.FgBlue).Printf(format, args...)
		}
	}

	// Format output as "type owner group permissions size date name"
	colorFunc("%s%s %s %s %s %s %s\n", typeChar, perms, owner, group, sizeFormatted, dateFormatted, name)
}

// Function to format the size (bytes to KB/MB/GB)
func formatSize(size int64) string {
	if size >= 1024*1024*1024 {
		return fmt.Sprintf("%.2f GB", float64(size)/1024/1024/1024)
	} else if size >= 1024*1024 {
		return fmt.Sprintf("%.2f MB", float64(size)/1024/1024)
	} else if size >= 1024 {
		return fmt.Sprintf("%.2f KB", float64(size)/1024)
	}
	return fmt.Sprintf("%d bytes", size)
}

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
		longListingValue, err := cmd.Flags().GetBool("long")
		if err != nil {
			panic(err)
		}
		colorValue, err := cmd.Flags().GetBool("color")
		if err != nil {
			panic(err)
		}
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
			result, err := application.DataStoreService.Ls(
				WithLsLongListing(longListingValue),
				WithLsColor(colorValue),
				WithLsArgs(args))
			if err != nil {
				fmt.Println(err)
				return
			}
			for _, db := range result.Databases {
				printEntry("d", db.Name, db.Perms, db.Owner, db.Group, db.Size, db.LastModified)
			}
			for _, col := range result.Collections {
				printEntry(".", col.Name, col.Perms, col.Owner, col.Group, col.Size, col.LastModified)
			}
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
