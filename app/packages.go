package app

// Example to follow : https://github.com/charmbracelet/bubbletea/tree/master/tutorials/commands/

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