package cmd

import (
	"bytes"
	jsonEnc "encoding/json"
	"fmt"
	"regexp"
	"strings"
	"syscall"

	"github.com/alecthomas/chroma/v2/quick"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/indent"
	"github.com/muesli/reflow/wordwrap"
	"github.com/muesli/reflow/wrap"
	"github.com/skysyzygy/tq/tq"
	"github.com/spf13/pflag"
	"golang.org/x/term"
)

var (
	width  = 80
	subtle = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	// highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	// success   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}

	// successStyle = lipgloss.NewStyle().Foreground(success)

	paraStyle = lipgloss.NewStyle().
			Align(lipgloss.Left).
		//Foreground(lipgloss.Color("#FAFAFA")).
		//Background(highlight).
		Padding(1, 2)

	// titleStyle = lipgloss.NewStyle().
	// 		MarginLeft(1).
	// 		MarginRight(5).
	// 		Padding(0, 1).
	// 		Italic(true) //.
	// 	//Foreground(highlight)

	hiliteStyle = lipgloss.NewStyle().
		//Foreground(lipgloss.Color("#FFFDF5")).
		//Background(lipgloss.Color("#FF5F87")).
		Padding(0, 1)
)

func helpParagraph(para string) string {
	return lipgloss.JoinVertical(lipgloss.Center,
		paraStyle.Width(width).Margin(1, 0).Render(para),
		lipgloss.PlaceHorizontal(width, lipgloss.Center,
			hiliteStyle.Margin(0, 1).Render(`Tlön, Uqbar, Orbis Tertius`),
			lipgloss.WithWhitespaceChars("="),
			lipgloss.WithWhitespaceForeground(subtle),
		))
}

// flagUsagesWrapped returns a string containing the usage information
// for all flags in the FlagSet. Borrowed from pflag, and made ANSI-aware
// using Reflow. Wrapped to `cols` columns (0 for no wrapping)
func flagUsagesWrapped(cols int, flatten bool, f *pflag.FlagSet) string {
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

		prose := regexp.MustCompile("{.+}$").ReplaceAllString(usage, "")
		json := regexp.MustCompile("{.+}$").FindString(usage)

		line += prose + jsonHighlight(json, flatten)

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

// exampleWrapped indents and wraps the `example` (query) text
// for tq / cobra commands using ANSI-aware wrapping to `cols` width and a
// 2-column indent
func exampleWrapped(cols int, flatten bool, example string) string {
	buf := new(bytes.Buffer)
	example = jsonHighlight(example, flatten)
	// try to wrap at cols-8 and if that fails enforce at cols wide
	wrapped := wrap.String(wordwrap.String(example, cols-8), cols)
	for _, subline := range strings.Split(wrapped, "\n") {
		fmt.Fprintln(buf, indent.String(subline, 2))
	}
	return buf.String()
}

// Syntax highlighting for json strings using chroma
// and ptionally flatten using FlattenJSONMap
func jsonHighlight(json string, flatten bool) string {
	w := new(bytes.Buffer)
	if flatten && len(json) > 0 {
		json = strings.ReplaceAll(json, ", ...", "")
		j, err := tq.FlattenJSONMap([]byte(json), "")
		if err == nil {
			jm, _ := jsonEnc.Marshal(j)
			json = string(jm)
			json = strings.ReplaceAll(json, ",", ", ")
		}
	}
	if highlight || term.IsTerminal(syscall.Stdout) && !noHighlight {
		err := quick.Highlight(w, json, "json", "terminal256", "vulcan")
		if err == nil {
			return w.String()
		}
	}
	// fallback - return input unchanged
	return json
}
