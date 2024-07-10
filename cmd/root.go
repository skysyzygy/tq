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
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"runtime/debug"
	"strings"
	"syscall"

	"github.com/99designs/keyring"
	"github.com/muesli/reflow/indent"
	"github.com/muesli/reflow/wordwrap"
	"github.com/muesli/reflow/wrap"

	"github.com/skysyzygy/tq/auth"
	"github.com/skysyzygy/tq/tq"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

// Version number
const version string = "0.1.1"

var (
	cfgFile, jsonFile, logFile string
	verbose, dryRun            bool
	_tq                        *tq.TqConfig
	keys                       auth.Keyring
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
	if _tq != nil {
		fmt.Println(string(_tq.GetOutput()))
	}
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

	// enable case insensitive command names
	cobra.EnableCaseInsensitive = true
	// enable case insensitive flag names
	rootCmd.SetGlobalNormalizationFunc(func(f *pflag.FlagSet, name string) pflag.NormalizedName {
		return pflag.NormalizedName(strings.ToLower(name))
	})

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

	width, _, err := term.GetSize(int(syscall.Stdin))
	if err != nil {
		width = 0 // no wrapping
	}

	rootCmd.SetUsageTemplate(
		// Rename some things so that they align better with how they are used
		strings.NewReplacer("command", "verb", " Command", " Verb", "Examples", "Query",
			// Wrap flag usages using ANSI-aware function
			".FlagUsages", " | flagUsagesWrapped "+fmt.Sprint(width),
			// Indent example
			"{{.Example}}", "  {{.Example}}").
			Replace(rootCmd.UsageTemplate()))

	cobra.AddTemplateFunc("flagUsagesWrapped", flagUsagesWrapped)

	keys, _ = keyring.Open(keyring.Config{
		ServiceName: "tq",
	})

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

	initLog()
}

func initLog() {
	var log *os.File
	var err error
	if logFile != "" {
		// open log file for appending
		log, err = os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: cannot open log file %v for appending.", logFile)
		}
	}
	_tq = tq.New(log, verbose, dryRun)
}

// Initializes a tq instance with input from file or stdin
// and logs it in using the default authentication method.
// Shouldn't be called until the last minute in order to make sure
// all flags are set and that we don't unnecessarily ping the server.
func initTq(cmd *cobra.Command, args []string) (err error) {
	var input io.Reader
	var _err error
	if jsonFile != "" {
		// open input file for reading
		input, _err = os.OpenFile(jsonFile, os.O_RDONLY, 0644)
		if _err != nil {
			err = errors.Join(fmt.Errorf("cannot open input file %v for reading", jsonFile), _err, err)
		}
	} else if !term.IsTerminal(int(syscall.Stdin)) {
		input = cmd.InOrStdin()
	}

	a, _err := auth.FromString(viper.GetString("Login"))
	if _err != nil {
		err = errors.Join(fmt.Errorf("bad login string in config file"), _err, err)
	}

	err = errors.Join(a.Load(keys), err)

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

	_tq.SetInput(input)
	return _tq.Login(a)
}

// FlagUsagesWrapped is borrowed from pflag, lightly modified for
// ANSI-aware wrapping using Reflow. Wrapped to `cols` columns (0 for no
// wrapping)
func flagUsagesWrapped(cols int, f *pflag.FlagSet) string {
	buf := new(bytes.Buffer)

	lines := make([]string, 0, f.NFlag())

	maxlen := 0
	f.VisitAll(func(flag *pflag.Flag) {
		if flag.Hidden {
			return
		}

		line := ""
		if flag.Shorthand != "" && flag.ShorthandDeprecated == "" {
			line = fmt.Sprintf("  -%s, --%s", flag.Shorthand, flag.Name)
		} else {
			line = fmt.Sprintf("      --%s", flag.Name)
		}

		varname, usage := pflag.UnquoteUsage(flag)
		if varname != "" {
			line += " " + varname
		}
		if flag.NoOptDefVal != "" {
			switch flag.Value.Type() {
			case "string":
				line += fmt.Sprintf("[=\"%s\"]", flag.NoOptDefVal)
			case "bool":
				if flag.NoOptDefVal != "true" {
					line += fmt.Sprintf("[=%s]", flag.NoOptDefVal)
				}
			case "count":
				if flag.NoOptDefVal != "+1" {
					line += fmt.Sprintf("[=%s]", flag.NoOptDefVal)
				}
			default:
				line += fmt.Sprintf("[=%s]", flag.NoOptDefVal)
			}
		}

		// This special character will be replaced with spacing once the
		// correct alignment is calculated
		line += "\x00"
		if len(line) > maxlen {
			maxlen = len(line)
		}

		line += usage
		// if !flag.defaultIsZeroValue() {
		// 	if flag.Value.Type() == "string" {
		// 		line += fmt.Sprintf(" (default %q)", flag.DefValue)
		// 	} else {
		// 		line += fmt.Sprintf(" (default %s)", flag.DefValue)
		// 	}
		// }
		if len(flag.Deprecated) != 0 {
			line += fmt.Sprintf(" (DEPRECATED: %s)", flag.Deprecated)
		}

		lines = append(lines, line)
	})

	for _, line := range lines {
		sidx := strings.Index(line, "\x00")
		spacing := strings.Repeat(" ", maxlen-sidx)
		fmt.Fprint(buf, line[:sidx])
		// maxlen + 2 comes from + 1 for the \x00 and + 1 for the (deliberate) off-by-one in maxlen-sidx
		// try to wrap at cols-8 and if that fails enforce at cols wide
		wrapped := wrap.String(wordwrap.String(line[sidx+1:], cols-maxlen-10), cols-maxlen-2)
		for i, subline := range strings.Split(wrapped, "\n") {
			if i == 0 {
				fmt.Fprintln(buf, spacing, subline)
			} else {
				fmt.Fprintln(buf, indent.String(subline, uint(maxlen+1)))
			}
		}
	}

	return buf.String()
}
