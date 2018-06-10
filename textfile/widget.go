package textfile

import (
	"fmt"
	"io/ioutil"

	"github.com/gdamore/tcell"
	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/andrewzolotukhin/wtf/wtf"
)

// Config is a pointer to the global config object
var Config *config.Config

const HelpText = `
  Keyboard commands for Textfile:

    /: Show/hide this help window
    o: Open the text file in the operating system
`

type Widget struct {
	wtf.TextWidget

	app      *tview.Application
	filePath string
	pages    *tview.Pages
}

func NewWidget(app *tview.Application, pages *tview.Pages) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" Text File ", "textfile", true),

		app:      app,
		filePath: Config.UString("wtf.mods.textfile.filePath"),
		pages:    pages,
	}

	widget.View.SetWrap(true)
	widget.View.SetWordWrap(true)

	widget.View.SetInputCapture(widget.keyboardIntercept)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	widget.UpdateRefreshedAt()
	widget.View.SetTitle(fmt.Sprintf("%s %s", widget.Name, widget.filePath))

	filePath, _ := wtf.ExpandHomeDir(widget.filePath)

	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		fileData = []byte{}
	}

	if err != nil {
		widget.View.SetText(fmt.Sprintf("%s", err))
	} else {
		widget.View.SetText(fmt.Sprintf("%s", string(fileData)))
	}
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	switch string(event.Rune()) {
	case "/":
		widget.showHelp()
		return nil
	case "o":
		wtf.OpenFile(widget.filePath)
		return nil
	}

	return event
}

func (widget *Widget) showHelp() {
	closeFunc := func() {
		widget.pages.RemovePage("help")
		widget.app.SetFocus(widget.View)
	}

	modal := wtf.NewBillboardModal(HelpText, closeFunc)

	widget.pages.AddPage("help", modal, false, true)
	widget.app.SetFocus(modal)
}
