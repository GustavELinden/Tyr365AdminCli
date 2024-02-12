package uiList

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type menuItem struct {
	Name        string
	Description string
	HasFlags    bool
	Command     *cobra.Command
}

type model struct {
	items      []menuItem
	cursor     int
	parents    []*cobra.Command
	currentCmd *cobra.Command
}

func (m model) Init() tea.Cmd {
	// Initial setup can go here
	return nil
}
func NewModel(rootCmd *cobra.Command) model {
	m := model{
		currentCmd: rootCmd,
	}
	m.updateItemsFromCurrentCmd() // Populate the initial list with root command's subcommands
	return m
}
func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}
		case "enter":
			// Navigate into a command or execute action
			m.navigate()
		case "backspace":
			// Navigate back to parent command
			m.navigateBack()
		}
	}
	return m, nil
}
func (m *model) navigate() {
	selectedItem := m.items[m.cursor]

	// If the selected command has subcommands, show them
	if len(selectedItem.Command.Commands()) > 0 {
		m.parents = append(m.parents, m.currentCmd) // Push current command to parent stack
		m.currentCmd = selectedItem.Command          // Update current command
		m.updateItemsFromCurrentCmd()                // Update items to show subcommands
		m.cursor = 0                                 // Reset cursor for new list
	} else if selectedItem.HasFlags {
		// Handle flag input here or mark the command for flag processing
		// This could involve transitioning to a flag input state
		fmt.Println("Selected a command that requires flag input. Implement flag handling.")
	}
}

func (m *model) updateItemsFromCurrentCmd() {
	m.items = []menuItem{} // Reset items
	for _, cmd := range m.currentCmd.Commands() {
		item := menuItem{
			Name:        cmd.Name(),
			Description: cmd.Short,
			HasFlags:    len(cmd.Flags().FlagUsages()) > 0,
			Command:     cmd,
		}
		m.items = append(m.items, item)
	}
}
func (m *model) navigateBack() {
	if len(m.parents) > 0 {
		m.currentCmd = m.parents[len(m.parents)-1] // Pop the last parent command
		m.parents = m.parents[:len(m.parents)-1]   // Remove it from the stack
		m.updateItemsFromCurrentCmd()              // Update items to show current command's subcommands
		m.cursor = 0                               // Reset cursor for new list
	}
}

func (m model) View() string {
	var b strings.Builder
	for i, item := range m.items {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		flagIndicator := ""
		if item.HasFlags {
			flagIndicator = " *"
			// Example: Use color library to colorize
			flagIndicator = color.New(color.FgYellow).Sprint(flagIndicator)
		}
		fmt.Fprintf(&b, "%s %s%s\n", cursor, item.Name, flagIndicator)
	}
	return b.String()
}
