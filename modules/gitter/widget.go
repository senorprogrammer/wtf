package gitter

import (
	"fmt"
	"strconv"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

const HelpText = `
 Keyboard commands for Gitter:

   /: Show/hide this help window
   j: Select the next message in the list
   k: Select the previous message in the list
   r: Refresh the data

   arrow down: Select the next message in the list
   arrow up:   Select the previous message in the list
`

// A Widget represents a Gitter widget
type Widget struct {
	wtf.HelpfulWidget
	wtf.KeyboardWidget
	wtf.TextWidget

	app *tview.Application

	messages []Message
	selected int
	settings *Settings
}

// NewWidget creates a new instance of a widget
func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		HelpfulWidget:  wtf.NewHelpfulWidget(app, pages, HelpText),
		KeyboardWidget: wtf.NewKeyboardWidget(),
		TextWidget:     wtf.NewTextWidget(app, settings.common, true),

		app:      app,
		settings: settings,
	}

	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	widget.unselect()

	widget.View.SetScrollable(true)
	widget.View.SetRegions(true)

	widget.HelpfulWidget.SetView(widget.View)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	room, err := GetRoom(widget.settings.roomURI, widget.settings.apiToken)
	if err != nil {
		widget.View.SetWrap(true)
		widget.View.SetTitle(widget.CommonSettings.Title)
		widget.View.SetText(err.Error())
		return
	}

	if room == nil {
		return
	}

	messages, err := GetMessages(room.ID, widget.settings.numberOfMessages, widget.settings.apiToken)

	if err != nil {
		widget.View.SetWrap(true)

		widget.Redraw(widget.CommonSettings.Title, err.Error(), true)
		return
	}
	widget.messages = messages

	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) display() {
	if widget.messages == nil {
		return
	}

	title := fmt.Sprintf("%s - %s", widget.CommonSettings.Title, widget.settings.roomURI)

	widget.Redraw(title, widget.contentFrom(widget.messages), true)
	widget.app.QueueUpdateDraw(func() {
		widget.View.Highlight(strconv.Itoa(widget.selected)).ScrollToHighlight()
		widget.View.ScrollToEnd()
	})
}

func (widget *Widget) contentFrom(messages []Message) string {
	var str string
	for idx, message := range messages {
		str = str + fmt.Sprintf(
			`["%d"][""][%s] [blue]%s [lightslategray]%s: [%s]%s [aqua]%s`,
			idx,
			widget.rowColor(idx),
			message.From.DisplayName,
			message.From.Username,
			widget.rowColor(idx),
			message.Text,
			message.Sent.Format("Jan 02, 15:04 MST"),
		)

		str = str + "\n"
	}

	return str
}

func (widget *Widget) rowColor(idx int) string {
	if widget.View.HasFocus() && (idx == widget.selected) {
		return widget.settings.common.DefaultFocussedRowColor()
	}

	return widget.settings.common.RowColor(idx)
}

func (widget *Widget) next() {
	widget.selected++
	if widget.messages != nil && widget.selected >= len(widget.messages) {
		widget.selected = 0
	}

	widget.display()
}

func (widget *Widget) prev() {
	widget.selected--
	if widget.selected < 0 && widget.messages != nil {
		widget.selected = len(widget.messages) - 1
	}

	widget.display()
}

func (widget *Widget) openMessage() {
	sel := widget.selected
	if sel >= 0 && widget.messages != nil && sel < len(widget.messages) {
		message := &widget.messages[widget.selected]
		wtf.OpenFile(message.Text)
	}
}

func (widget *Widget) unselect() {
	widget.selected = -1
	widget.display()
}
