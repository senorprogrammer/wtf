package jenkins

import (
	"fmt"
	"strconv"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
	"regexp"
)

const HelpText = `
 Keyboard commands for Jenkins:

   /: Show/hide this help window
   j: Select the next job in the list
   k: Select the previous job in the list
   r: Refresh the data

   arrow down: Select the next job in the list
   arrow up:   Select the previous job in the list

   return: Open the selected job in a browser
`

type Widget struct {
	wtf.HelpfulWidget
	wtf.TextWidget

	app      *tview.Application
	selected int
	settings *Settings
	view     *View
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		HelpfulWidget: wtf.NewHelpfulWidget(app, pages, HelpText),
		TextWidget:    wtf.NewTextWidget(app, settings.common, true),

		app:      app,
		settings: settings,
	}

	widget.HelpfulWidget.SetView(widget.View)
	widget.unselect()

	widget.View.SetScrollable(true)
	widget.View.SetRegions(true)
	widget.View.SetInputCapture(widget.keyboardIntercept)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	view, err := widget.Create(
		widget.settings.url,
		widget.settings.user,
		widget.settings.apiKey,
	)
	widget.view = view

	if err != nil {
		widget.View.SetWrap(true)

		widget.app.QueueUpdateDraw(func() {
			widget.View.SetText(err.Error())
		})
	}

	widget.app.QueueUpdateDraw(func() {
		widget.display()
	})
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) display() {
	if widget.view == nil {
		return
	}

	widget.View.SetWrap(false)

	widget.View.Clear()
	widget.View.SetTitle(widget.ContextualTitle(fmt.Sprintf("%s: [red]%s", widget.CommonSettings.Title, widget.view.Name)))
	widget.View.SetText(widget.contentFrom(widget.view))
	widget.View.Highlight(strconv.Itoa(widget.selected)).ScrollToHighlight()
}

func (widget *Widget) contentFrom(view *View) string {
	var str string
	for idx, job := range view.Jobs {
		var validID = regexp.MustCompile(widget.settings.jobNameRegex)

		if validID.MatchString(job.Name) {
			str = str + fmt.Sprintf(
				`["%d"][""][%s] [%s]%-6s[white]`,
				idx,
				widget.rowColor(idx),
				widget.jobColor(&job),
				job.Name,
			)

			str = str + "\n"
		}
	}

	return str
}

func (widget *Widget) rowColor(idx int) string {
	if widget.View.HasFocus() && (idx == widget.selected) {
		return widget.settings.common.DefaultFocussedRowColor()
	}

	return widget.settings.common.RowColor(idx)
}

func (widget *Widget) jobColor(job *Job) string {
	switch job.Color {
	case "blue":
		// Override color if successBallColor boolean param provided in config
		return widget.settings.successBallColor
	case "red":
		return "red"
	default:
		return "white"
	}
}

func (widget *Widget) next() {
	widget.selected++
	if widget.view != nil && widget.selected >= len(widget.view.Jobs) {
		widget.selected = 0
	}

	widget.display()
}

func (widget *Widget) prev() {
	widget.selected--
	if widget.selected < 0 && widget.view != nil {
		widget.selected = len(widget.view.Jobs) - 1
	}

	widget.display()
}

func (widget *Widget) openJob() {
	sel := widget.selected
	if sel >= 0 && widget.view != nil && sel < len(widget.view.Jobs) {
		job := &widget.view.Jobs[widget.selected]
		wtf.OpenFile(job.Url)
	}
}

func (widget *Widget) unselect() {
	widget.selected = -1
	widget.display()
}

func (widget *Widget) keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	switch string(event.Rune()) {
	case "/":
		widget.ShowHelp()
	case "j":
		widget.next()
		return nil
	case "k":
		widget.prev()
		return nil
	case "r":
		widget.Refresh()
		return nil
	}

	switch event.Key() {
	case tcell.KeyDown:
		widget.next()
		return nil
	case tcell.KeyEnter:
		widget.openJob()
		return nil
	case tcell.KeyEsc:
		widget.unselect()
		return event
	case tcell.KeyUp:
		widget.prev()
		return nil
	default:
		return event
	}
}
