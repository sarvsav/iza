package app

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var badgeStyle = lipgloss.NewStyle().
	Padding(0, 1).
	MarginRight(1).
	Background(lipgloss.Color("#5E81AC")). // Nordic blue
	Foreground(lipgloss.Color("#ECEFF4")). // Light text
	Bold(true)

// Model struct to store the application state
type Model struct {
	Stage       int
	Services    list.Model
	ServiceType string
	Inputs      []textinput.Model
	FocusIndex  int
	ErrorMsg    string
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
	errorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Bold(true) // Red
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
		case "q":
			if m.Stage == 2 || m.Stage == 3 {
				return m, tea.Quit
			}
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
				// Validate current input before moving forward
				if strings.TrimSpace(m.Inputs[m.FocusIndex].Value()) == "" {
					m.ErrorMsg = "Please fill out the field before proceeding."
					return m, nil
				}
				m.ErrorMsg = "" // clear error if okay
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
				switch msg.String() {
				case "enter":
					// Save and go to success stage
					if err := m.SaveConfigToCueFile(); err != nil {
						m.ErrorMsg = "❌ Failed to save: " + err.Error()
						return m, nil
					}
					m.Stage = 3
					return m, nil
				case "q":
					// Exit without saving
					return m, tea.Quit
				}
			}
			if m.Stage == 3 {
				if msg.String() == "enter" || msg.String() == "q" {
					return m, tea.Quit
				}
			}
		case "up", "down":
			if m.Stage == 0 {
				var cmd tea.Cmd
				m.Services, cmd = m.Services.Update(msg)
				return m, cmd
			}
		}
	}

	// If typing in inputs
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
		badge := badgeStyle.Render(fmt.Sprintf("[%s %s]", icon, serviceName))
		s += boldStyle.Render(fmt.Sprintf("%s  Configure", badge))
		s += "\n\n"

		labels := map[string][]string{
			"mongodb": {"MongoDB URL", "Username", "Password"},
			"jenkins": {"Base URL", "BlueOcean Installed?", "Username", "API Token"},
			"jfrog":   {"JFrog URL", "Repository Name", "Username", "Password"},
		}

		currentLabels := labels[m.ServiceType]

		for i, input := range m.Inputs {
			var prefix string
			if i == m.FocusIndex {
				prefix = greenStyle.Render("➤")
			} else {
				prefix = "  "
			}

			s += fmt.Sprintf("%s %s:\n", prefix, greenStyle.Render(currentLabels[i]))
			s += input.View() + "\n\n"
		}

		if m.ErrorMsg != "" {
			s += "\n" + errorStyle.Render("❗ "+m.ErrorMsg) + "\n"
		}

		s += "\n[Press Enter to move next]\n"
		return s
	}

	if m.Stage == 2 {
		s += fmt.Sprintf("🔒 Review your information:\n\n")

		labels := map[string][]string{
			"mongodb": {"MongoDB URL", "Username", "Password"},
			"jenkins": {"Base URL", "BlueOcean Installed?", "Username", "API Token"},
			"jfrog":   {"JFrog URL", "Repository Name", "Username", "Password"},
		}

		currentLabels := labels[m.ServiceType]

		maxLen := 0
		for _, label := range currentLabels {
			if len(label) > maxLen {
				maxLen = len(label)
			}
		}
		maxLen += 2

		for i, input := range m.Inputs {
			label := fmt.Sprintf("➤ %s", currentLabels[i])
			paddedLabel := fmt.Sprintf("%-*s", maxLen, label)
			value := input.Value()

			s += fmt.Sprintf("%s %s\n", greenStyle.Render(paddedLabel), value)
		}

		s += "\n[Press Enter to confirm and save, or q to quit without saving]\n"
		return s
	}

	if m.Stage == 3 {
		msg := "✅ Details saved successfully!\n\n👉 You can now run:\n\n" +
			boldStyle.Render("myapp ls") + "\n\n" +
			"Press " + boldStyle.Render("q") + " or " + boldStyle.Render("Enter") + " to exit."

		successBox := lipgloss.NewStyle().
			Padding(1, 2).
			Margin(1, 0).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("10")).
			Render(msg)

		return successBox
	}

	return s
}

func (m Model) SaveConfigToCueFile() error {
	var sb strings.Builder

	// Start writing cue structure
	sb.WriteString("package dev\n\n")
	sb.WriteString("config: {\n")

	switch m.ServiceType {
	case "mongodb":
		sb.WriteString("  database: {\n")
		sb.WriteString("    mongo01: {\n")
		sb.WriteString("      kind: \"database\",\n")
		sb.WriteString("      type: \"mongodb\",\n")
		sb.WriteString(fmt.Sprintf("      host: \"%s\",\n", m.Inputs[0].Value()))
		sb.WriteString("      port: 27017,\n") // You could later make port dynamic too
		sb.WriteString("      credentials: {\n")
		sb.WriteString(fmt.Sprintf("        user: \"%s\",\n", m.Inputs[1].Value()))
		sb.WriteString(fmt.Sprintf("        pass: \"%s\",\n", m.Inputs[2].Value()))
		sb.WriteString("      },\n")
		sb.WriteString("    },\n")
		sb.WriteString("  },\n")

	case "jenkins":
		sb.WriteString("  ci_tools: {\n")
		sb.WriteString("    jenkins01: {\n")
		sb.WriteString("      kind: \"ci-tools\",\n")
		sb.WriteString("      tool: \"Jenkins\",\n")
		sb.WriteString(fmt.Sprintf("      endpoint: \"%s\",\n", m.Inputs[0].Value()))
		sb.WriteString("      auth: {\n")
		sb.WriteString(fmt.Sprintf("        method: \"%s\",\n", "token"))
		sb.WriteString(fmt.Sprintf("        user: \"%s\",\n", m.Inputs[2].Value()))
		sb.WriteString(fmt.Sprintf("        token: \"%s\",\n", m.Inputs[3].Value()))
		sb.WriteString("      },\n")
		sb.WriteString("    },\n")
		sb.WriteString("  },\n")

	case "jfrog":
		sb.WriteString("  artifactory: {\n")
		sb.WriteString("    artifactory01: {\n")
		sb.WriteString("      kind: \"artifactory\",\n")
		sb.WriteString(fmt.Sprintf("      url: \"%s\",\n", m.Inputs[0].Value()))
		sb.WriteString(fmt.Sprintf("      repo: \"%s\",\n", m.Inputs[1].Value()))
		sb.WriteString("      auth: {\n")
		sb.WriteString(fmt.Sprintf("        user: \"%s\",\n", m.Inputs[2].Value()))
		sb.WriteString(fmt.Sprintf("        pass: \"%s\",\n", m.Inputs[3].Value()))
		sb.WriteString("      },\n")
		sb.WriteString("    },\n")
		sb.WriteString("  },\n")
	}

	sb.WriteString("}\n")

	// Create folder if not exist
	if _, err := os.Stat("dev"); os.IsNotExist(err) {
		err := os.Mkdir("dev", 0755)
		if err != nil {
			return err
		}
	}

	// Write to dev/config.cue
	return os.WriteFile("dev/config.cue", []byte(sb.String()), 0644)
}
