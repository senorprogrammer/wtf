package victorops

import (
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

// Widget contains text info
type Widget struct {
	wtf.TextWidget

	teams    []OnCallTeam
	settings *Settings
}

// NewWidget creates a new widget
func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(app, pages, settings.common, true),
	}

	widget.SetRefreshFunction(widget.Refresh)
	widget.View.SetScrollable(true)
	widget.View.SetRegions(true)

	return &widget
}

// Refresh gets latest content for the widget
func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	teams, err := Fetch(widget.settings.apiID, widget.settings.apiKey)

	if err != nil {
		widget.Redraw(widget.CommonSettings.Title, err.Error(), true)
	} else {
		widget.teams = teams
		widget.Redraw(widget.CommonSettings.Title, widget.contentFrom(widget.teams), true)
	}
}

func (widget *Widget) contentFrom(teams []OnCallTeam) string {
	var str string

	if teams == nil || len(teams) == 0 {
		return "No teams specified"
	}

	for _, team := range teams {
		if len(widget.settings.team) > 0 && widget.settings.team != team.Slug {
			continue
		}

		str = fmt.Sprintf("%s[green]%s\n", str, team.Name)
		if len(team.OnCall) == 0 {
			str = fmt.Sprintf("%s[grey]no one\n", str)
		}
		for _, onCall := range team.OnCall {
			str = fmt.Sprintf("%s[white]%s - %s\n", str, onCall.Policy, onCall.Userlist)
		}

		str = fmt.Sprintf("%s\n", str)
	}

	if len(str) == 0 {
		str = "Could not find any teams to display"
	}
	return str
}
