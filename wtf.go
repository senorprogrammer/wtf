package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gdamore/tcell"
	flags "github.com/jessevdk/go-flags"
	"github.com/olebedev/config"
	"github.com/radovskyb/watcher"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/bamboohr"
	"github.com/senorprogrammer/wtf/clocks"
	"github.com/senorprogrammer/wtf/cmdrunner"
	"github.com/senorprogrammer/wtf/gcal"
	"github.com/senorprogrammer/wtf/git"
	"github.com/senorprogrammer/wtf/github"
	"github.com/senorprogrammer/wtf/help"
	"github.com/senorprogrammer/wtf/jira"
	"github.com/senorprogrammer/wtf/newrelic"
	"github.com/senorprogrammer/wtf/opsgenie"
	"github.com/senorprogrammer/wtf/power"
	"github.com/senorprogrammer/wtf/security"
	"github.com/senorprogrammer/wtf/status"
	"github.com/senorprogrammer/wtf/system"
	"github.com/senorprogrammer/wtf/textfile"
	"github.com/senorprogrammer/wtf/todo"
	"github.com/senorprogrammer/wtf/weather"
	"github.com/senorprogrammer/wtf/wtf"
)

/* -------------------- Functions -------------------- */

func addToGrid(grid *tview.Grid, widget wtf.Wtfable) {
	if widget.Disabled() {
		return
	}

	grid.AddItem(
		widget.TextView(),
		widget.Top(),
		widget.Left(),
		widget.Height(),
		widget.Width(),
		0,
		0,
		false,
	)
}

// Grid stores all the widgets onscreen (like an HTML table)
func buildGrid(modules []wtf.Wtfable) *tview.Grid {
	grid := tview.NewGrid()
	grid.SetColumns(wtf.ToInts(Config.UList("wtf.grid.columns"))...)
	grid.SetRows(wtf.ToInts(Config.UList("wtf.grid.rows"))...)
	grid.SetBorder(false)

	for _, module := range modules {
		addToGrid(grid, module)
		go wtf.Schedule(module)
	}

	return grid
}

func keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyCtrlR:
		refreshAllWidgets()
	case tcell.KeyTab:
		FocusTracker.Next()
	case tcell.KeyBacktab:
		FocusTracker.Prev()
	case tcell.KeyEsc:
		FocusTracker.None()
	default:
		return event
	}

	return event
}

// redrawApp redraws the rendered views to screen on a defined interval (set in config.yml)
// Use this because each textView widget can have it's own update interval, and I don't want to
// manage drawing co-ordination amongst them all. If you need to have a
// widget redraw on it's own schedule, use the view's SetChangedFunc() and pass it `app`.
func redrawApp(app *tview.Application) {
	tick := time.NewTicker(time.Duration(Config.UInt("wtf.refreshInterval", 2)) * time.Second)
	quit := make(chan struct{})

	for {
		select {
		case <-tick.C:
			app.Draw()
		case <-quit:
			tick.Stop()
			return
		}
	}
}

func refreshAllWidgets() {
	for _, widget := range Widgets {
		go widget.Refresh()
	}
}

func watchForConfigChanges(app *tview.Application, configFlag *string, grid *tview.Grid, pages *tview.Pages) {
	watch := watcher.New()

	// notify write events.
	watch.FilterOps(watcher.Write)

	go func() {
		for {
			select {
			case <-watch.Event:
				loadConfig(configFlag)
				makeWidgets(app, pages)
				grid = buildGrid(Widgets)
				pages.AddPage("grid", grid, true, true)
			case err := <-watch.Error:
				log.Fatalln(err)
			case <-watch.Closed:
				return
			}
		}
	}()

	// Watch config file for changes.
	if err := watch.Add(*configFlag); err != nil {
		log.Fatalln(err)
	}

	// Start the watching process - it'll check for changes every 100ms.
	if err := watch.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}

/* -------------------- Main -------------------- */

var Config *config.Config
var FocusTracker wtf.FocusTracker
var Widgets []wtf.Wtfable

var (
	commit  = "dev"
	date    = "dev"
	version = "dev"
)

func makeWidgets(app *tview.Application, pages *tview.Pages) {
	bamboohr.Config = Config
	clocks.Config = Config
	cmdrunner.Config = Config
	gcal.Config = Config
	git.Config = Config
	github.Config = Config
	jira.Config = Config
	newrelic.Config = Config
	opsgenie.Config = Config
	power.Config = Config
	security.Config = Config
	status.Config = Config
	system.Config = Config
	textfile.Config = Config
	todo.Config = Config
	weather.Config = Config
	wtf.Config = Config

	Widgets = []wtf.Wtfable{
		bamboohr.NewWidget(),
		clocks.NewWidget(),
		cmdrunner.NewWidget(),
		gcal.NewWidget(),
		git.NewWidget(app, pages),
		github.NewWidget(app, pages),
		jira.NewWidget(),
		newrelic.NewWidget(),
		opsgenie.NewWidget(),
		power.NewWidget(),
		security.NewWidget(),
		status.NewWidget(),
		system.NewWidget(date, version),
		textfile.NewWidget(app, pages),
		todo.NewWidget(app, pages),
		weather.NewWidget(app, pages),
	}

	FocusTracker = wtf.FocusTracker{
		App:     app,
		Idx:     -1,
		Widgets: Widgets,
	}
}

func loadConfig(configFlag *string) {
	Config = wtf.LoadConfigFile(*configFlag)
}

func main() {

	/*
		This allows the user to pass flags in however they prefer. It supports the likes of:

		wtf -help    | --help
		wtf -version | --version
	*/
	homeDir, err := wtf.Home()
	if err != nil {
		os.Exit(1)
	}

	var cmdFlags struct {
		Version bool   `short:"v" long:"version" description:"Show Version Info"`
		Config  string `short:"c" long:"config" optional:"yes" description:"Path to config file"`
	}

	var parser = flags.NewParser(&cmdFlags, flags.Default)
	parser.Parse()

	if len(cmdFlags.Config) == 0 {
		cmdFlags.Config = filepath.Join(homeDir, ".wtf", "config.yml")
	}

	if cmdFlags.Version {
		help.DisplayVersionInfo(version)
	}

	/* -------------------- end flag parsing and handling -------------------- */

	// Responsible for creating the configuration directory and default
	// configuration file if they don't already exist
	wtf.CreateConfigDir()
	wtf.WriteConfigFile()

	loadConfig(&cmdFlags.Config)

	app := tview.NewApplication()
	pages := tview.NewPages()

	makeWidgets(app, pages)

	grid := buildGrid(Widgets)
	pages.AddPage("grid", grid, true, true)
	app.SetInputCapture(keyboardIntercept)

	// Loop in a routine to redraw the screen
	go redrawApp(app)
	go watchForConfigChanges(app, &cmdFlags.Config, grid, pages)

	if err := app.SetRoot(pages, true).Run(); err != nil {
		os.Exit(1)
	}
}
