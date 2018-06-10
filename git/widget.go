package git

import (
	"github.com/andrewzolotukhin/wtf/wtf"
	"github.com/gdamore/tcell"
	"github.com/olebedev/config"
	"github.com/rivo/tview"
)

// Config is a pointer to the global config object
var Config *config.Config

const HelpText = `
  Keyboard commands for Git:

    /: Show/hide this help window
    h: Previous git repository
    l: Next git repository
	p: Pull current git repository
	c: Checkout to branch

    arrow left:  Previous git repository
    arrow right: Next git repository
`

const offscreen = -1000
const modalWidth = 80
const modalHeight = 7

type Widget struct {
	wtf.TextWidget

	app   *tview.Application
	Data  []*GitRepo
	Idx   int
	pages *tview.Pages
}

func NewWidget(app *tview.Application, pages *tview.Pages) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" Git ", "git", true),

		app:   app,
		Idx:   0,
		pages: pages,
	}

	widget.View.SetInputCapture(widget.keyboardIntercept)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	repoPaths := wtf.ToStrs(Config.UList("wtf.mods.git.repositories"))

	widget.UpdateRefreshedAt()
	widget.Data = widget.gitRepos(repoPaths)
	widget.display()
}

func (widget *Widget) Next() {
	widget.Idx = widget.Idx + 1
	if widget.Idx == len(widget.Data) {
		widget.Idx = 0
	}

	widget.display()
}

func (widget *Widget) Prev() {
	widget.Idx = widget.Idx - 1
	if widget.Idx < 0 {
		widget.Idx = len(widget.Data) - 1
	}

	widget.display()
}
func (widget *Widget) Pull() {
	repoToPull := widget.Data[widget.Idx]
	repoToPull.pull()
	widget.Refresh()

}
func (widget *Widget) Checkout() {
	form := widget.modalForm("Branch to checkout:", "")

	checkoutFctn := func() {
		text := form.GetFormItem(0).(*tview.InputField).GetText()
		repoToCheckout := widget.Data[widget.Idx]
		repoToCheckout.checkout(text)
		widget.pages.RemovePage("modal")
		widget.app.SetFocus(widget.View)
		widget.display()
		widget.Refresh()
	}

	widget.addButtons(form, checkoutFctn)
	widget.modalFocus(form)

}

/* -------------------- Unexported Functions -------------------- */
func (widget *Widget) addCheckoutButton(form *tview.Form, fctn func()) {
	form.AddButton("Checkout", fctn)
}
func (widget *Widget) addButtons(form *tview.Form, checkoutFctn func()) {
	widget.addCheckoutButton(form, checkoutFctn)
	widget.addCancelButton(form)
}
func (widget *Widget) addCancelButton(form *tview.Form) {
	cancelFn := func() {
		widget.pages.RemovePage("modal")
		widget.app.SetFocus(widget.View)
		widget.display()
	}

	form.AddButton("Cancel", cancelFn)
	form.SetCancelFunc(cancelFn)
}
func (widget *Widget) modalFocus(form *tview.Form) {
	frame := widget.modalFrame(form)
	widget.pages.AddPage("modal", frame, false, true)
	widget.app.SetFocus(frame)
}

func (widget *Widget) modalForm(lbl, text string) *tview.Form {
	form := tview.NewForm().
		SetButtonsAlign(tview.AlignCenter).
		SetButtonTextColor(tview.Styles.PrimaryTextColor)

	form.AddInputField(lbl, text, 60, nil, nil)

	return form
}
func (widget *Widget) modalFrame(form *tview.Form) *tview.Frame {
	frame := tview.NewFrame(form).SetBorders(0, 0, 0, 0, 0, 0)
	frame.SetRect(offscreen, offscreen, modalWidth, modalHeight)
	frame.SetBorder(true)
	frame.SetBorders(1, 1, 0, 0, 1, 1)

	drawFunc := func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
		w, h := screen.Size()
		frame.SetRect((w/2)-(width/2), (h/2)-(height/2), width, height)
		return x, y, width, height
	}

	frame.SetDrawFunc(drawFunc)

	return frame
}

func (widget *Widget) currentData() *GitRepo {
	if len(widget.Data) == 0 {
		return nil
	}

	if widget.Idx < 0 || widget.Idx >= len(widget.Data) {
		return nil
	}

	return widget.Data[widget.Idx]
}

func (widget *Widget) gitRepos(repoPaths []string) []*GitRepo {
	repos := []*GitRepo{}

	for _, repoPath := range repoPaths {
		repo := NewGitRepo(repoPath)
		repos = append(repos, repo)
	}

	return repos
}

func (widget *Widget) keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	switch string(event.Rune()) {
	case "/":
		widget.showHelp()
		return nil
	case "h":
		widget.Prev()
		return nil
	case "l":
		widget.Next()
		return nil
	case "p":
		widget.Pull()
		return nil
	case "c":
		widget.Checkout()
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

func (widget *Widget) showHelp() {
	closeFunc := func() {
		widget.pages.RemovePage("help")
		widget.app.SetFocus(widget.View)
	}

	modal := wtf.NewBillboardModal(HelpText, closeFunc)

	widget.pages.AddPage("help", modal, false, true)
	widget.app.SetFocus(modal)
}
