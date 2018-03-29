package main

import (
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/bamboohr"
	"github.com/senorprogrammer/wtf/gcal"
	"github.com/senorprogrammer/wtf/status"
	"github.com/senorprogrammer/wtf/weather"
)

func main() {
	app := tview.NewApplication()

	grid := tview.NewGrid()
	grid.SetRows(14, 36, 4) // How _high_ the row is, in terminal rows
	grid.SetColumns(40, 40) // How _wide_ the column is, in terminal columns
	grid.SetBorder(false)

	grid.AddItem(bamboohr.Widget(), 0, 0, 1, 1, 0, 0, false)
	grid.AddItem(gcal.Widget(), 1, 0, 1, 1, 0, 0, false)
	grid.AddItem(status.Widget(), 2, 0, 2, 3, 0, 0, false)
	grid.AddItem(weather.Widget(), 0, 1, 1, 1, 0, 0, false)

	if err := app.SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}
