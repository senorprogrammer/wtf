package wtf

import (
	"github.com/rivo/tview"
)

type Wtfable interface {
	Enabler
	Scheduler

	BorderColor() string
	FocusChar() string
	Focusable() bool
	Name() string
	SetFocusChar(string)
	TextView() *tview.TextView

	Height() int
	Left() int
	Top() int
	Width() int
}
