package github

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
)

const HelpText = `
  Keyboard commands for GitHub:

    /: Show/hide this help window
    h: Previous git repository
    l: Next git repository
    r: Refresh the data

    arrow left:  Previous git repository
    arrow right: Next git repository
`

type Widget struct {
	wtf.HelpfulWidget
	wtf.TextWidget

	GithubRepos []*GithubRepo
	Activity    *GithubActivity
	Idx         int
}

func NewWidget(app *tview.Application, pages *tview.Pages) *Widget {
	widget := Widget{
		HelpfulWidget: wtf.NewHelpfulWidget(app, pages, HelpText),
		TextWidget:    wtf.NewTextWidget("GitHub", "github", true),

		Idx: 0,
	}

	widget.GithubRepos = widget.buildRepoCollection(wtf.Config.UMap("wtf.mods.github.repositories"))
	notificationUser, err := wtf.Config.String("wtf.mods.github.notificationUser")
	if err == nil {
		widget.Activity = NewGithubActivities(notificationUser)
	}

	widget.HelpfulWidget.SetView(widget.View)
	widget.View.SetInputCapture(widget.keyboardIntercept)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	for _, repo := range widget.GithubRepos {
		repo.Refresh()
	}

	if widget.Activity != nil {
		widget.Activity.Refresh()
	}

	widget.UpdateRefreshedAt()
	widget.display()
}

func (widget *Widget) Next() {
	widget.Idx = widget.Idx + 1
	if widget.Idx == len(widget.GithubRepos) {
		widget.Idx = 0
	}

	widget.display()
}

func (widget *Widget) Prev() {
	widget.Idx = widget.Idx - 1
	if widget.Idx < 0 {
		widget.Idx = len(widget.GithubRepos) - 1
	}

	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) buildRepoCollection(repoData map[string]interface{}) []*GithubRepo {
	githubRepos := []*GithubRepo{}

	for name, owner := range repoData {
		repo := NewGithubRepo(name, owner.(string))
		githubRepos = append(githubRepos, repo)
	}

	return githubRepos
}

func (widget *Widget) currentGithubRepo() *GithubRepo {
	if len(widget.GithubRepos) == 0 {
		return nil
	}

	if widget.Idx < 0 || widget.Idx >= len(widget.GithubRepos) {
		return nil
	}

	return widget.GithubRepos[widget.Idx]
}

func (widget *Widget) keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	switch string(event.Rune()) {
	case "/":
		widget.ShowHelp()
		return nil
	case "h":
		widget.Prev()
		return nil
	case "l":
		widget.Next()
		return nil
	case "r":
		widget.Refresh()
		return nil
	}

	switch event.Key() {
	case tcell.KeyLeft:
		widget.Prev()
		return nil
	case tcell.KeyRight:
		widget.Next()
		return nil
	default:
		return event
	}
}
