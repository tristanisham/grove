package app

import (

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type model struct {
	List list.Model
	err  error
}

func BrowseModal(list list.Model) model {
	return model{
		List: list,
		err:  nil,
	}
}

type statusMsg int
type errMsg struct {err error}
// For messages that contain errors it's often handy to also implement the
// error interface on the message.
func (e errMsg) Error() string { return e.err.Error() }


//Init is all about calling the setup functions for our view.
func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)

	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.List.View())
}

type Plant struct {
	title, desc, author, website string
}

func (p Plant) Title() string       { return p.title }
func (p Plant) Description() string { return p.desc }
func (p Plant) FilterValue() string { return p.title }
func (p Plant) Website() string     { return p.website }
func (p Plant) Author() string      { return p.author }


func NewPlant(title, desc, author, website string) Plant {
	return Plant{
		title:   title,
		desc:    desc,
		author:  author,
		website: website,
	}
}

