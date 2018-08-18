package github

import (
	"fmt"

	"github.com/google/go-github/github"
	"github.com/senorprogrammer/wtf/wtf"
)

func (widget *Widget) display() {
	repo := widget.currentGithubRepo()
	if repo == nil {
		widget.View.SetText(" GitHub repo data is unavailable ")
		return
	}

	widget.View.SetTitle(widget.ContextualTitle(fmt.Sprintf("%s - %s", widget.Name, widget.title(repo))))

	str := wtf.SigilStr(len(widget.GithubRepos), widget.Idx, widget.View) + "\n"
	str = str + " [red]Stats[white]\n"
	str = str + widget.displayStats(repo)
	str = str + "\n"
	str = str + " [red]Open Review Requests[white]\n"
	str = str + widget.displayMyReviewRequests(repo, wtf.Config.UString("wtf.mods.github.username"))
	str = str + "\n"
	str = str + " [red]My Pull Requests[white]\n"
	str = str + widget.displayMyPullRequests(repo, wtf.Config.UString("wtf.mods.github.username"))
	str = str + "\n"
	str = str + widget.displayNotifications()
	str = str + "\n"

	widget.View.SetText(str)
}

func (widget *Widget) displayMyPullRequests(repo *GithubRepo, username string) string {
	prs := repo.myPullRequests(username)

	if len(prs) == 0 {
		return " [grey]none[white]\n"
	}

	str := ""
	for _, pr := range prs {
		str = str + fmt.Sprintf(" %s[green]%4d[white] %s\n", mergeString(pr), *pr.Number, *pr.Title)
	}

	return str
}

func (widget *Widget) displayMyReviewRequests(repo *GithubRepo, username string) string {
	prs := repo.myReviewRequests(username)

	if len(prs) == 0 {
		return " [grey]none[white]\n"
	}

	str := ""
	for _, pr := range prs {
		str = str + fmt.Sprintf(" [green]%4d[white] %s\n", *pr.Number, *pr.Title)
	}

	return str
}

func (widget *Widget) displayNotifications() string {
	if widget.Activity == nil {
		return ""
	}

	str := wtf.SigilStr(len(widget.Activity.Notifications), widget.Idx, widget.View) + "\n"
	for _, notification := range widget.Activity.Notifications {
		str = str + "(" + notification.Type + ") "
		str = str + notification.Title
		str = str + " - "
		str = str + notification.URL
		str = str + "\n"
	}

	return str
}

func (widget *Widget) displayStats(repo *GithubRepo) string {
	str := fmt.Sprintf(
		" PRs: %d  Issues: %d  Stars: %d\n",
		repo.PullRequestCount(),
		repo.IssueCount(),
		repo.StarCount(),
	)

	return str
}

func (widget *Widget) title(repo *GithubRepo) string {
	return fmt.Sprintf("[green]%s - %s[white]", repo.Owner, repo.Name)
}

func showStatus() bool {
	return wtf.Config.UBool("wtf.mods.github.enableStatus", false)
}

var mergeIcons = map[string]string{
	"dirty":    "[red]![white] ",
	"clean":    "[green]✔[white] ",
	"unstable": "[red]✖[white] ",
	"blocked":  "[red]✖[white] ",
}

func mergeString(pr *github.PullRequest) string {
	if !showStatus() {
		return ""
	}
	if str, ok := mergeIcons[pr.GetMergeableState()]; ok {
		return str
	}
	return "? "
}
