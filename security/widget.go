package security

import (
	"fmt"
	//"sort"
	"time"

	"github.com/gdamore/tcell"
	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
)

var Config *config.Config

type Widget struct {
	wtf.BaseWidget
	View *tview.TextView
}

func NewWidget() *Widget {
	widget := Widget{
		BaseWidget: wtf.BaseWidget{
			Name:        "Security",
			RefreshedAt: time.Now(),
			RefreshInt:  Config.UInt("wtf.security.refreshInterval", 3600),
		},
	}

	widget.addView()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	data := Fetch()

	widget.View.SetTitle(" 🤺 Security ")
	widget.RefreshedAt = time.Now()

	widget.View.Clear()
	fmt.Fprintf(widget.View, "%s", widget.contentFrom(data))
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) addView() {
	view := tview.NewTextView()

	view.SetBorder(true)
	view.SetBorderColor(tcell.ColorGray)
	view.SetDynamicColors(true)
	view.SetTitle(widget.Name)
	view.SetWrap(false)

	widget.View = view
}

func (widget *Widget) contentFrom(data map[string]string) string {
	str := fmt.Sprintf(" Firewall: %s  Network: %s\n", data["Enabled"], data["Network"])
	str = str + fmt.Sprintf("  Stealth: %s   Crypto: %s\n", data["Stealth"], data["Encryption"])

	return str
}
