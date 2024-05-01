/*
Copyright © 2024 Sky Syzygy ssyzygy@bam.org

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/skysyzygy/tq/auth"
	"github.com/skysyzygy/tq/tq"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile, jsonFile, logFile string
	verbose, dryRun            bool
	_tq                        *tq.TqConfig
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tq",
	Short: "A toolkit for Tessitura",
	Long: helpParagraph("tq is a wrapper around the Tessitura API that reads " +
		"JSON-formatted data and outputs a series of API calls to Tessitura. " +
		"It internally handles authentication, session creation and " +
		"closure, and batch/concurrent processing so that humans like " +
		"you can focus on the data and not the intricacies of the API.\n\n" +
		"tq is basically a high-level 	API for common tasks in Tessi. "),

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
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

	//rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.tq.yaml)")
	rootCmd.PersistentFlags().StringVarP(&jsonFile, "file", "f", "", "JSON file to read (default is to read from stdin)")
	rootCmd.PersistentFlags().StringVarP(&logFile, "log", "l", "", "log file to write to (default is no log)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "turns on additional diagnostic output")
	rootCmd.PersistentFlags().BoolVarP(&dryRun, "dryrun", "n", false, "don't actually do anything, just show what would have happened")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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

		// Search config in home directory with name ".tq" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".tq")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

// Initializes a tq instance with logger, verbosity, and dryness,
// and logs it in using the default authentication method.
// Defined here but shouldn't be called until the last minute in order to make sure
// all flags are set and that we don't unnecessarily ping the server.
func tqInit(cmd *cobra.Command, args []string) {
	var log *os.File
	var err error
	if logFile != "" {
		// open log file for appending
		log, err = os.OpenFile(logFile, os.O_APPEND&os.O_CREATE, 0644)
		if err != nil {
			fmt.Println("Couldn't open log file: ", logFile)
			os.Exit(1)
		}
	}
	_tq = tq.New(log, verbose, dryRun)
	a, err := auth.FromString(viper.GetString("Login"))
	if err != nil {
		_tq.Log.Error("Bad login string in config file", "error", err.Error(), "login", a)
		os.Exit(1)
	}
	a.Load()
	if valid, err := a.Validate(); !valid || err != nil {
		_tq.Log.Error("Invalid login", "error", err.Error(), "login", a)
		os.Exit(1)
	}
	_tq.Login(a)
}
