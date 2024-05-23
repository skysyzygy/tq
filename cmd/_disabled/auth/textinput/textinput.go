package auth

// A simple example demonstrating the use of multiple text input components
// from the Bubbles component library.
// Forked from https://github.com/charmbracelet/bubbletea/blob/master/examples/textinputs/main.go

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle = focusedStyle.Copy().Foreground(lipgloss.Color("240"))
	cursorStyle  = focusedStyle.Copy()

	promptStyle = focusedStyle.Copy().Width(20)
	noStyle     = lipgloss.NewStyle().Width(20)

	focusedButton = focusedStyle.Copy().Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

type model struct {
	focusIndex int
	inputs     []textinput.Model
	cursorMode cursor.Mode
	errorMsg   string
	submitted  bool
}

func initialModel() model {
	m := model{
		inputs: make([]textinput.Model, 5),
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 64
		t.PromptStyle = noStyle
		t.Validate = func(v string) (err error) {
			if v == "" {
				err = errors.New("Please enter all required values")
			}
			return
		}
		switch i {
		case 0:
			t.Prompt = "API server:"
			t.Placeholder = "https://"
			t.Focus()
			t.PromptStyle = promptStyle
			t.TextStyle = focusedStyle
		case 1:
			t.Prompt = "Username:"
			t.Placeholder = "Enter your Tessitura username"
		case 2:
			t.Prompt = "Usergroup:"
			t.Placeholder = "(Optional) Enter the short description of your usergroup (e.g. PlanMem)"
			t.Validate = nil
		case 3:
			t.Prompt = "Location/machine:"
			t.Placeholder = "Usually your Tessitura username followed by `-14` will work"
		case 4:
			t.Prompt = "Password:"
			t.Placeholder = "Enter your password"
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '•'
		}

		m.inputs[i] = t
	}

	return m
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && m.focusIndex == len(m.inputs) {
				// Validate inputs
				for _, input := range m.inputs {
					if err := input.Validate(input.Value()); err != nil {
						m.errorMsg = err.Error()
						return m, nil
					}
				}
				m.submitted = true
				return m, tea.Quit
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = promptStyle
					m.inputs[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = noStyle
				m.inputs[i].TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m model) View() string {
	var b strings.Builder

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n%s\n\n", *button, m.errorMsg)

	return b.String()
}

// Spin up the bubbles/bubbletea textinput and return the values from the user as an array
func AuthInput() ([]string, error) {
	m, err := tea.NewProgram(initialModel()).Run()

	if err != nil {
		fmt.Printf("Could not start the auth/textinput module: %s\n", err)
		os.Exit(1)
	}

	if !m.(model).submitted {
		return nil, errors.New("operation cancelled")
	}

	values := make([]string, 5)
	for i, val := range m.(model).inputs {
		values[i] = val.Value()
	}

	return values, nil
}
