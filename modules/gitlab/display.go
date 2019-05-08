package gitlab

import (
	"fmt"
)

func (widget *Widget) display() {

	project := widget.currentGitlabProject()
	if project == nil {
		widget.Redraw(widget.CommonSettings.Title, " Gitlab project data is unavailable ", true)
		return
	}

	title := fmt.Sprintf("%s- %s", widget.CommonSettings.Title, widget.title(project))

	_, _, width, _ := widget.View.GetRect()
	str := widget.settings.common.SigilStr(len(widget.GitlabProjects), widget.Idx, width) + "\n"
	str = str + " [red]Stats[white]\n"
	str = str + widget.displayStats(project)
	str = str + "\n"
	str = str + " [red]Open Approval Requests[white]\n"
	str = str + widget.displayMyApprovalRequests(project, widget.settings.username)
	str = str + "\n"
	str = str + " [red]My Merge Requests[white]\n"
	str = str + widget.displayMyMergeRequests(project, widget.settings.username)
	widget.Redraw(title, str, false)
}

func (widget *Widget) displayMyMergeRequests(project *GitlabProject, username string) string {
	mrs := project.myMergeRequests(username)

	if len(mrs) == 0 {
		return " [grey]none[white]\n"
	}

	str := ""
	for _, mr := range mrs {
		str = str + fmt.Sprintf(" [green]%4d[white] %s\n", mr.IID, mr.Title)
	}

	return str
}

func (widget *Widget) displayMyApprovalRequests(project *GitlabProject, username string) string {
	mrs := project.myApprovalRequests(username)

	if len(mrs) == 0 {
		return " [grey]none[white]\n"
	}

	str := ""
	for _, mr := range mrs {
		str = str + fmt.Sprintf(" [green]%4d[white] %s\n", mr.IID, mr.Title)
	}

	return str
}

func (widget *Widget) displayStats(project *GitlabProject) string {
	str := fmt.Sprintf(
		" MRs: %d  Issues: %d  Stars: %d\n",
		project.MergeRequestCount(),
		project.IssueCount(),
		project.StarCount(),
	)

	return str
}

func (widget *Widget) title(project *GitlabProject) string {
	return fmt.Sprintf("[green]%s [white]", project.Path)
}
