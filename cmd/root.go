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
	"errors"
	"fmt"
	"io"
	"os"
	"runtime/debug"
	"slices"
	"strings"

	"github.com/skysyzygy/tq/auth"
	"github.com/skysyzygy/tq/tq"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Version number
const version string = "0.1"

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
		"tq is basically a high-level API for common tasks in Tessi. "),
	Version: version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		if _tq != nil && _tq.Log != nil {
			_tq.Log.Error(err.Error())
		} else {
			fmt.Println("Error: ", err.Error())
		}
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	settings := make(map[string]string)
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			settings[setting.Key] = setting.Value
		}
	}
	commit := strings.Join([]string{settings["vcs"], settings["vcs.revision"], settings["vcs.time"]}, " ")
	rootCmd.Version = rootCmd.Version + " (" + commit + ")"

	//rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.tq)")
	rootCmd.PersistentFlags().StringVarP(&jsonFile, "file", "f", "", "JSON file to read (default is to read from stdin)")
	rootCmd.PersistentFlags().StringVarP(&logFile, "log", "l", "", "log file to write to (default is no log)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "turns on additional diagnostic output")
	rootCmd.PersistentFlags().BoolVarP(&dryRun, "dryrun", "n", false, "don't actually do anything, just show what would have happened")

	// Hide global flags from auth command
	authenticateCmd.SetUsageFunc(func(cmd *cobra.Command) error {
		authenticateCmd.InheritedFlags().VisitAll(func(f *pflag.Flag) { f.Hidden = true })
		return rootCmd.UsageFunc()(cmd)
	})

	rootCmd.SetUsageTemplate(
		strings.NewReplacer("command", "verb", " Command", " Verb").
			Replace(rootCmd.UsageTemplate()))

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
		cfgFile = home + string(os.PathSeparator) + ".tq"
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	} else {
		cfg, err := os.OpenFile(cfgFile, os.O_CREATE|os.O_WRONLY, 0644)
		cfg.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Warning: couldn't access config file")
		}
	}
}

// Initializes a tq instance with logger, verbosity, and dryness,
// and logs it in using the default authentication method.
// Defined here but shouldn't be called until the last minute in order to make sure
// all flags are set and that we don't unnecessarily ping the server.
func tqInit(cmd *cobra.Command, args []string) (err error) {
	var log, json *os.File
	var _err error

	if logFile != "" {
		// open log file for appending
		log, _err = os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if _err != nil {
			err = errors.Join(fmt.Errorf("cannot open log file for appending"), _err, err)
		}
	}
	_tq = tq.New(log, verbose, dryRun)

	if jsonFile != "" {
		// open input file for reading
		json, _err = os.OpenFile(jsonFile, os.O_RDONLY, 0644)
		if _err != nil {
			err = errors.Join(fmt.Errorf("cannot open input file for reading"), _err, err)
		}
		input, _err := io.ReadAll(json)
		json.Close()
		if _err != nil {
			err = errors.Join(fmt.Errorf("cannot read from input file"), _err, err)
		}
		cmd.SetArgs(slices.Concat([]string{string(input)}, args))
	}

	a, _err := auth.FromString(viper.GetString("Login"))
	if _err != nil {
		err = errors.Join(fmt.Errorf("bad login string in config file"), _err, err)
	}

	err = errors.Join(a.Load(), err)

	if valid, _err := a.Validate(); !valid || _err != nil {
		err = errors.Join(fmt.Errorf("invalid login"), _err, err)
	}

	if err != nil {
		authString, _ := a.String()
		_tq.Log.Error(err.Error(),
			"logFile", logFile,
			"jsonFile", jsonFile,
			"auth", authString)
		return err
	}

	return _tq.Login(a)
}
