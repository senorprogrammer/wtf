package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/rivo/tview"
	log "github.com/wtfutil/wtf/logger"
	"github.com/wtfutil/wtf/wtf"
)

const maxBufferSize int64 = 1024

type Widget struct {
	wtf.TextWidget

	app      *tview.Application
	filePath string
	settings *Settings
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(app, pages, settings.common, true),

		app:      app,
		filePath: log.LogFilePath(),
		settings: settings,
	}

	widget.SetRefreshFunction(widget.Refresh)

	return &widget
}

// Refresh updates the onscreen contents of the widget
func (widget *Widget) Refresh() {
	if log.LogFileMissing() {
		return
	}

	logLines := widget.tailFile()

	widget.Redraw(widget.CommonSettings.Title, widget.contentFrom(logLines), false)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) contentFrom(logLines []string) string {
	str := ""

	for _, line := range logLines {
		chunks := strings.Split(line, " ")

		if len(chunks) >= 4 {
			str = str + fmt.Sprintf(
				"[green]%s[white] [yellow]%s[white] %s\n",
				chunks[0],
				chunks[1],
				strings.Join(chunks[3:], " "),
			)
		}
	}

	return str
}

func (widget *Widget) tailFile() []string {
	file, err := os.Open(widget.filePath)
	if err != nil {
		return []string{}
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return []string{}
	}

	bufferSize := maxBufferSize
	if maxBufferSize > stat.Size() {
		bufferSize = stat.Size()
	}

	startPos := stat.Size() - bufferSize

	buff := make([]byte, bufferSize)
	_, err = file.ReadAt(buff, startPos)
	if err != nil {
		return []string{}
	}

	logLines := strings.Split(string(buff), "\n")

	// Reverse the array of lines
	// Offset by two to account for the blank line at the end
	last := len(logLines) - 2
	for i := 0; i < len(logLines)/2; i++ {
		logLines[i], logLines[last-i] = logLines[last-i], logLines[i]
	}

	return logLines
}
