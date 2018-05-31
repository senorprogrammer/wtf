package newrelic

import (
	"fmt"

	"github.com/olebedev/config"
	"github.com/senorprogrammer/wtf/wtf"
	nr "github.com/yfronto/newrelic"
)

// Config is a pointer to the global config object
var Config *config.Config

type Widget struct {
	wtf.TextWidget
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" New Relic ", "newrelic", false),
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	app, appErr := Application()
	deploys, depErr := Deployments()

	appName := "error"
	if appErr == nil {
		appName = app.Name
	}

	widget.UpdateRefreshedAt()
	widget.View.SetTitle(fmt.Sprintf(" New Relic: [green]%s[white] ", appName))
	widget.View.Clear()

	if depErr != nil {
		widget.View.SetWrap(true)
		fmt.Fprintf(widget.View, "%s", depErr)
	} else {
		widget.View.SetWrap(false)
		fmt.Fprintf(widget.View, "%s", widget.contentFrom(deploys))
	}
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) contentFrom(deploys []nr.ApplicationDeployment) string {
	str := fmt.Sprintf(
		" %s\n",
		"[red]Latest Deploys[white]",
	)

	revisions := []string{}

	for _, deploy := range deploys {
		if (deploy.Revision != "") && wtf.Exclude(revisions, deploy.Revision) {
			lineColor := "white"
			if wtf.IsToday(deploy.Timestamp) {
				lineColor = "lightblue"
			}

			str = str + fmt.Sprintf(
				" [green]%s[%s] %s %-.16s[white]\n",
				deploy.Revision[0:],
				lineColor,
				deploy.Timestamp.Format("Jan 02, 15:04 MST"),
				wtf.NameFromEmail(deploy.User),
			)

			revisions = append(revisions, deploy.Revision)

			if len(revisions) == Config.UInt("wtf.mods.newrelic.deployCount", 5) {
				break
			}
		}
	}

	return str
}
