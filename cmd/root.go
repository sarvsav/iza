package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"path"
	"strings"
	"time"

	"github.com/sarvsav/iza/internals/app"
	"github.com/sarvsav/iza/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var application *app.Application

func checkCueConfig() error {
	if _, err := os.Stat("dev/config.cue"); os.IsNotExist(err) {
		return errors.New("❌ config.cue file not found. Please run `iza init` to create it")
	}
	return nil
}

func TouchFile(filePath string) error {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	return file.Close()
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "iza",
	Short: "interact mongodb with linux alike commands using iza",
	Long: `Iza is a CLI tool to interact with MongoDB using linux alike commands.

With simplicity in mind, iza provides a set of commands to interact with
MongoDB in a way that is familiar to linux users. The linux commands are mapped
to MongoDB operations, so you can use the commands you already know to interact
with MongoDB. It is designed to be simple, easy to use, and easy to remember.

A few examples of iza commands are:
  iza ls
  iza touch
  iza rm

You can also use iza to interact with MongoDB in a more advanced way, such as:
  iza find
  iza insert

And, detailed information about each command can be found in the help menu.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Skip config check for some special commands like "init"
		if cmd.Name() == "init" {
			return nil
		}
		return checkCueConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Check for version flag, print version information and exit
		if cmd.Flag("version").Changed {
			fmt.Println("iza:", version.Get())
			os.Exit(0)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(a *app.Application) {
	application = a
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory:", err)
		os.Exit(1)
	}

	logDir := userHomeDir + "/.iza/logs"
	logFileName := "iza_" + version.Get().String() + ".log"
	historyFile := path.Join(userHomeDir, ".iza", "history")
	if _, err := os.Stat(historyFile); os.IsNotExist(err) {
		err := TouchFile(historyFile)
		if err != nil {
			fmt.Println("Error creating history file:", err)
			os.Exit(1)
		}
		err = os.WriteFile(historyFile, []byte("host, user, working directory, timestamp, command\n"), 0644)
		if err != nil {
			fmt.Println("Error writing to history file:", err)
			os.Exit(1)
		}
	}

	hostname, _ := os.Hostname()
	currentUser, _ := user.Current()
	workingDir, _ := os.Getwd()
	startTime := time.Now().Format(time.RFC3339)

	hf, err := os.OpenFile(historyFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening history file:", err)
		os.Exit(1)
	}
	defer hf.Close()
	if _, err := hf.Write([]byte(fmt.Sprintf("%s, %s, %s, %s, %s\n", hostname, currentUser.Username, workingDir, startTime, strings.Join(os.Args[1:], " ")))); err != nil {
		fmt.Println("Error writing to history file:", err)
		os.Exit(1)
	}

	// Create log directory if it does not exist
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.MkdirAll(logDir, 0755)
		if err != nil {
			fmt.Println("Error creating log directory:", err)
			os.Exit(1)
		}
	}

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.iza.yaml)")
	rootCmd.PersistentFlags().StringP("service", "s", "datastore", "Default datastore service to interact with")
	rootCmd.PersistentFlags().StringP("log-file", "F", logFileName, "File to save the logs (optional), default name is iza_version_datetime.log")
	rootCmd.PersistentFlags().StringP("log-dir", "L", logDir, "Directory to save the logs (optional), default is $HOME/.iza/logs")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("version", "v", false, "Print iza version information and exit")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".iza" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".iza")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
