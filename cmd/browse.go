package cmd

import (
	"grove/app"
	"log"
	"github.com/charmbracelet/bubbles/list"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var browse *cobra.Command = &cobra.Command{
	Use:   "browse",
	Short: "Opens Grove's TUI",
	Long: `Grove has a TUI (text user interface) that allows you to tab around and view details of available packages
		 without needing to find and navigate an ancient website. Grove's TUI looks good on any device with a terminal and runs fast too!`,
	Run: func(cmd *cobra.Command, args []string) {
		items := []list.Item{
			app.NewPlant("Example", "A dummy package for testing terminal development", "hi", "hi"),
			app.NewPlant("Example1", "A dummy package for testing terminal development", "hi", "hi"),
			app.NewPlant("Example2", "A dummy package for testing terminal development", "hi", "hi"),
			app.NewPlant("Example3", "A dummy package for testing terminal development", "hi", "hi"),
			app.NewPlant("Example4", "A dummy package for testing terminal development", "hi", "hi"),
			app.NewPlant("Example5", "A dummy package for testing terminal development", "hi", "hi"),
			app.NewPlant("Example6", "A dummy package for testing terminal development", "hi", "hi"),
			app.NewPlant("Example7", "A dummy package for testing terminal development", "hi", "hi"),
			app.NewPlant("Example8", "A dummy package for testing terminal development", "hi", "hi"),

		}
		m := app.NewModel(list.New(items, list.NewDefaultDelegate(), 0, 0))
		m.List.Title = "Your Grove"
		p := tea.NewProgram(m, tea.WithAltScreen())
		if err := p.Start(); err != nil {
			log.Fatalf("Oops, there was an error running Grove: %v", err)
		}
	},
}
