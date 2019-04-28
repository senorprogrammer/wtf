package zendesk

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget

	app      *tview.Application
	result   *TicketArray
	selected int
	settings *Settings
}

func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(app, settings.common, true),

		app:      app,
		settings: settings,
	}

	widget.View.SetInputCapture(widget.keyboardIntercept)

	return &widget
}

/* -------------------- Exported Functions -------------------- */
func (widget *Widget) Refresh() {
	ticketArray, err := widget.newTickets()
	ticketArray.Count = len(ticketArray.Tickets)
	if err != nil {
		log.Fatal(err)
	} else {
		widget.result = ticketArray
	}

	widget.app.QueueUpdateDraw(func() {
		widget.display()
	})
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) display() {
	widget.View.SetTitle(fmt.Sprintf("%s (%d)", widget.CommonSettings.Title, widget.result.Count))
	widget.View.SetText(widget.textContent(widget.result.Tickets))
}

func (widget *Widget) textContent(items []Ticket) string {
	if len(items) == 0 {
		return fmt.Sprintf("No unassigned tickets in queue - woop!!")
	}

	str := ""
	for idx, data := range items {
		str = str + widget.format(data, idx)
	}

	return str
}

func (widget *Widget) format(ticket Ticket, idx int) string {
	textColor := widget.settings.common.Colors.Background
	if idx == widget.selected {
		textColor = widget.settings.common.Colors.BorderFocused
	}

	requesterName := widget.parseRequester(ticket)
	str := fmt.Sprintf(" [%s:]%d - %s\n %s\n\n", textColor, ticket.Id, requesterName, ticket.Subject)
	return str
}

// this is a nasty means of extracting the actual name of the requester from the Via interface of the Ticket.
// very very open to improvements on this
func (widget *Widget) parseRequester(ticket Ticket) interface{} {
	viaMap := ticket.Via
	via := viaMap.(map[string]interface{})
	source := via["source"]
	fromMap, _ := source.(map[string]interface{})
	from := fromMap["from"]
	fromValMap := from.(map[string]interface{})
	fromName := fromValMap["name"]
	return fromName
}

func (widget *Widget) next() {
	widget.selected++
	if widget.result != nil && widget.selected >= len(widget.result.Tickets) {
		widget.selected = 0
	}
}

func (widget *Widget) prev() {
	widget.selected--
	if widget.selected < 0 && widget.result != nil {
		widget.selected = len(widget.result.Tickets) - 1
	}
}

func (widget *Widget) openTicket() {
	sel := widget.selected
	if sel >= 0 && widget.result != nil && sel < len(widget.result.Tickets) {
		issue := &widget.result.Tickets[widget.selected]
		ticketUrl := fmt.Sprintf("https://%s.zendesk.com/agent/tickets/%d", widget.settings.subdomain, issue.Id)
		wtf.OpenFile(ticketUrl)
	}
}

func (widget *Widget) unselect() {
	widget.selected = -1
}

func (widget *Widget) keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	switch string(event.Rune()) {
	case "j":
		// Select the next item down
		widget.next()
		widget.display()
		return nil
	case "k":
		// Select the next item up
		widget.prev()
		widget.display()
		return nil
	}
	switch event.Key() {
	case tcell.KeyDown:
		// Select the next item down
		widget.next()
		widget.display()
		return nil
	case tcell.KeyUp:
		// Select the next item up
		widget.prev()
		widget.display()
		return nil
	case tcell.KeyEnter:
		widget.openTicket()
		return nil
	case tcell.KeyEsc:
		// Unselect the current row
		widget.unselect()
		widget.display()
		return event
	default:
		// Pass it along
		return event
	}
}
