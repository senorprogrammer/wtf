package wtf

import (
	"strconv"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/cfg"
)

type ScrollableWidget struct {
	TextWidget

	selected       int
	maxItems       int
	RenderFunction func()
}

func NewScrollableWidget(app *tview.Application, pages *tview.Pages, commonSettings *cfg.Common, focusable bool) *ScrollableWidget {

	widget := ScrollableWidget{
		TextWidget: NewTextWidget(app, pages, commonSettings, focusable),
	}

	widget.Unselect()
	widget.View.SetScrollable(true)
	widget.View.SetRegions(true)
	widget.SetKeyboardChar("j", widget.Next, "Select next item")
	widget.SetKeyboardChar("k", widget.Prev, "Select previous item")
	widget.SetKeyboardKey(tcell.KeyDown, widget.Next, "Select next item")
	widget.SetKeyboardKey(tcell.KeyEsc, widget.Unselect, "Clear selection")
	widget.SetKeyboardKey(tcell.KeyUp, widget.Prev, "Select previous item")

	return &widget
}

func (widget *ScrollableWidget) SetRenderFunction(displayFunc func()) {
	widget.RenderFunction = displayFunc
}

func (widget *ScrollableWidget) SetItemCount(items int) {
	widget.maxItems = items
}

func (widget *ScrollableWidget) GetSelected() int {
	return widget.selected
}

func (widget *ScrollableWidget) RowColor(idx int) string {
	if widget.View.HasFocus() && (idx == widget.selected) {
		return widget.CommonSettings.DefaultFocussedRowColor()
	}

	return widget.CommonSettings.RowColor(idx)
}

func (widget *ScrollableWidget) Next() {
	widget.selected++
	if widget.selected >= widget.maxItems {
		widget.selected = 0
	}
	widget.RenderFunction()
}

func (widget *ScrollableWidget) Prev() {
	widget.selected--
	if widget.selected < 0 {
		widget.selected = widget.maxItems - 1
	}
	widget.RenderFunction()
}

func (widget *ScrollableWidget) Unselect() {
	widget.selected = -1
	if widget.RenderFunction != nil {
		widget.RenderFunction()
	}
}

func (widget *ScrollableWidget) Redraw(title, content string, wrap bool) {

	widget.TextWidget.Redraw(title, content, wrap)
	widget.app.QueueUpdateDraw(func() {
		widget.View.Highlight(strconv.Itoa(widget.selected)).ScrollToHighlight()
	})
}
