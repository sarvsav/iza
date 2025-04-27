package app

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Model struct to store the application state
type Model struct {
	Stage       int
	Services    list.Model
	ServiceType string
	Inputs      []textinput.Model
	FocusIndex  int
}

// item struct for list items (service options)
type item struct {
	title, desc string
}

// Implement methods for list.Item interface
func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

// Lipgloss styles
var (
	greenStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("10")) // Green
	boldStyle    = lipgloss.NewStyle().Bold(true)
	selectedIcon = map[string]string{
		"Mongodb": "🍃",
		"Jenkins": "🛠️",
		"Jfrog":   "🐸",
	}
)

// InitialModel function initializes the first screen model
func InitialModel() Model {
	items := []list.Item{
		item{title: "MongoDB", desc: "🍃 Connect to a MongoDB database"},
		item{title: "Jenkins", desc: "🛠️ Configure Jenkins CI/CD"},
		item{title: "JFrog", desc: "🐸 Setup JFrog Artifactory"},
	}
	l := list.New(items, list.NewDefaultDelegate(), 30, 14)
	l.Title = "Choose a service"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)

	return Model{
		Services:   l,
		FocusIndex: 0,
	}
}

// Init method for initializing commands
func (m Model) Init() tea.Cmd {
	return nil
}

// Update method for handling user inputs
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			if m.Stage == 0 {
				selected := m.Services.SelectedItem().(item)
				m.ServiceType = strings.ToLower(selected.title)
				m.initInputsBasedOnService()
				m.Stage++
				return m, nil
			}
			if m.Stage == 1 {
				if m.FocusIndex == len(m.Inputs)-1 {
					m.Stage++
					return m, nil
				}
				m.FocusIndex++
				for i := range m.Inputs {
					if i == m.FocusIndex {
						m.Inputs[i].Focus()
					} else {
						m.Inputs[i].Blur()
					}
				}
				return m, nil
			}
			if m.Stage == 2 {
				return m, tea.Quit
			}
		case "up", "down":
			if m.Stage == 0 {
				var cmd tea.Cmd
				m.Services, cmd = m.Services.Update(msg)
				return m, cmd
			}
		}
	}

	// Update for inputs when stage 1 is active
	if m.Stage == 1 {
		cmds := make([]tea.Cmd, len(m.Inputs))
		for i := range m.Inputs {
			m.Inputs[i], cmds[i] = m.Inputs[i].Update(msg)
		}
		return m, tea.Batch(cmds...)
	}

	return m, nil
}

// Method to initialize inputs based on selected service
func (m *Model) initInputsBasedOnService() {
	var inputs []textinput.Model

	switch m.ServiceType {
	case "mongodb":
		url := textinput.New()
		url.Placeholder = "MongoDB URL"
		url.Focus()
		url.Width = 40

		username := textinput.New()
		username.Placeholder = "Username"
		username.Width = 40

		password := textinput.New()
		password.Placeholder = "Password"
		password.EchoMode = textinput.EchoPassword
		password.EchoCharacter = '•'
		password.Width = 40

		inputs = []textinput.Model{url, username, password}

	case "jenkins":
		url := textinput.New()
		url.Placeholder = "Base URL"
		url.Focus()
		url.Width = 40

		blueocean := textinput.New()
		blueocean.Placeholder = "BlueOcean Installed? (yes/no)"
		blueocean.Width = 40

		username := textinput.New()
		username.Placeholder = "Username"
		username.Width = 40

		token := textinput.New()
		token.Placeholder = "API Token"
		token.EchoMode = textinput.EchoPassword
		token.EchoCharacter = '•'
		token.Width = 40

		inputs = []textinput.Model{url, blueocean, username, token}

	case "jfrog":
		url := textinput.New()
		url.Placeholder = "JFrog URL"
		url.Focus()
		url.Width = 40

		repo := textinput.New()
		repo.Placeholder = "Repository Name"
		repo.Width = 40

		username := textinput.New()
		username.Placeholder = "Username"
		username.Width = 40

		password := textinput.New()
		password.Placeholder = "Password"
		password.EchoMode = textinput.EchoPassword
		password.EchoCharacter = '•'
		password.Width = 40

		inputs = []textinput.Model{url, repo, username, password}
	}

	m.Inputs = inputs
	m.FocusIndex = 0
}

// View method to render the UI
func (m Model) View() string {
	s := ""

	if m.Stage == 0 {
		s += boldStyle.Render("Select a service to configure:\n\n")
		s += m.Services.View()
		return s
	}

	if m.Stage == 1 {
		c := cases.Title(language.English)
		serviceName := c.String(m.ServiceType)
		icon := selectedIcon[serviceName]

		s += boldStyle.Render(fmt.Sprintf("%s Configure %s\n\n", icon, serviceName))

		labels := map[string][]string{
			"mongodb": {"MongoDB URL", "Username", "Password"},
			"jenkins": {"Base URL", "BlueOcean Installed?", "Username", "API Token"},
			"jfrog":   {"JFrog URL", "Repository Name", "Username", "Password"},
		}

		currentLabels := labels[m.ServiceType]

		// Add padding or no spaces before the first label
		for i, input := range m.Inputs {
			var prefix string
			if i == m.FocusIndex {
				// Show "➤" for the focused input
				prefix = greenStyle.Render("➤")
			} else {
				// For non-focused fields, keep an empty space or a short prefix
				prefix = "  "
			}

			// Align labels properly using fmt.Sprintf
			label := fmt.Sprintf("%s %s:", prefix, greenStyle.Render(currentLabels[i]))

			// Ensure proper alignment between the label and input field
			s += fmt.Sprintf("%-40s\n", label) // Align labels and inputs
			s += input.View() + "\n\n"
		}

		// Add instructions to move to the next step
		s += "\n[Press Enter to move next]\n"
		return s
	}

	if m.Stage == 2 {
		s += boldStyle.Render("🔒 Review your information:\n\n")

		labels := map[string][]string{
			"mongodb": {"MongoDB URL", "Username", "Password"},
			"jenkins": {"Base URL", "BlueOcean Installed?", "Username", "API Token"},
			"jfrog":   {"JFrog URL", "Repository Name", "Username", "Password"},
		}

		currentLabels := labels[m.ServiceType]

		// Find the maximum label length
		maxLen := 0
		for _, label := range currentLabels {
			if len(label) > maxLen {
				maxLen = len(label)
			}
		}

		maxLen += 2 // Extra padding after ➤

		for i, input := range m.Inputs {
			label := fmt.Sprintf("➤ %s", currentLabels[i])
			paddedLabel := fmt.Sprintf("%-*s", maxLen, label)
			value := input.Value()

			s += fmt.Sprintf("%s %s\n", greenStyle.Render(paddedLabel), value)
		}

		s += "\n[Press Enter to confirm and save]\n"
		return s
	}

	return s
}
