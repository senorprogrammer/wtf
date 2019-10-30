# Changelog

## Unreleased

### ⚡️ Added

* Proper, usable [Docker file](https://github.com/wtfutil/wtf/blob/master/Dockerfile) added, by [@Boot-Error](https://github.com/Boot-Error)
* [GitLab](https://wtfutil.com/modules/gitlab/) module displays issues assigned to, and opened by, the user, by [@caalberts](https://github.com/caalberts)
* [TravisCI](https://wtfutil.com/modules/travisci/) now checks for uncommitted vendor changes, by [@indradhanush](https://github.com/indradhanush)
* Football module added, by [@C123R](https://github.com/C123R)
* [resourceuseage](https://wtfutil.com/modules/resourceusage/) now supports a `cpuCombined` setting, by [@madepolli](https://github.com/madepolli)
* [Twitter Stats](https://wtfutil.com/modules/twitter/twittertweets/) module added, by [@Ameobea](https://github.com/Ameobea)

### 🐞 Fixed

* Github PRs do not count against issues, by [@alexfornuto](https://github.com/alexfornuto)
* Todo scrolling now works properly, [#707](https://github.com/wtfutil/wtf/issues/707) by [3mard](https://github.com/3mard)

## 0.23.0

### ⚡️ Added

* [Azure DevOps](https://wtfutil.com/modules/azure-devops/) module added, by [@v-braun](https://github.com/v-braun)
* [Dev.to](https://wtfutil.com/modules/devto/) module added, by [@VictorAvelar](https://github.com/VictorAvelar)
* [TravisCI]() module now supports enterprise endpoints, [#652](https://github.com/wtfutil/wtf/issues/652) by [@scw007](https://github.com/scw007)
* [Subreddit](https://wtfutil.com/modules/subreddit/) module added, by [@lawrencecraft](https://github.com/lawrencecraft)
* [gCal](https://wtfutil.com/modules/google/gcal/) module now supports a `hourFormat` setting for defining whether to display 12 or 24-hour times, [#665](https://github.com/wtfutil/wtf/issues/665) by [@senorprogrammer](https://github.com/senorprogrammer)
* [Scarf](https://scarf.sh) installation instructions added to README, by [@aviaviavi](https://github.com/aviaviavi)
* Spotify widget now supports colour themes, [#659](https://github.com/wtfutil/wtf/issues/659) by [@Tardog](https://github.com/Tardog)
* [Buildkite](https://wtfutil.com/modules/buildkite/) module added, by [@jmks](https://github.com/jmks)
* [Improvements](https://github.com/wtfutil/wtf/pull/680) to the [CmdRunner](https://wtfutil.com/modules/cmdrunner/) module, by [@noxer](https://github.com/noxer)

### 🐞 Fixed

* gCal calendar event time colour can now be changed by setting the `eventTime` configuration setting, [#638](https://github.com/wtfutil/wtf/issues/638) by [@indradhanush](https://github.com/indradhanush)
* [Clocks](https://wtfutil.com/modules/clocks/) now obeys global row colour settings, [#658](https://github.com/wtfutil/wtf/issues/658) by [@senorprogrammer](https://github.com/senorprogrammer)
* [Transmission](https://wtfutil.com/modules/transmission/) module no longer blocks rendering when a Transmission daemon cannot be found, [#661](https://github.com/wtfutil/wtf/issues/661) by [@senorprogrammer](https://github.com/senorprogrammer)
* [Trello](https://wtfutil.com/modules/trello/) module now respects project list order, [#664](https://github.com/wtfutil/wtf/issues/664) by [@Seanstoppable](https://github.com/Seanstoppable)
* [Todo](https://wtfutil.com/modules/todo/) module now respects checkbox settings, [#616](https://github.com/wtfutil/wtf/issues/616) by [@Seanstoppable](https://github.com/Seanstoppable)
* [Todoist](https://wtfutil.com/modules/todoist/) module now properly handles todo items with due date and times, [#645](https://github.com/wtfutil/wtf/issues/645) by [@massa1240](https://github.com/massa1240)
* Invalid pointer error in [Azure DevOps](https://wtfutil.com/modules/azure-devops/) fixed by [@Boot-Error](https://github.com/Boot-Error)
* Renamed slice error in [Dev](https://wtfutil.com/modules/devto/) fixed by [@Boot-Error](https://github.com/Boot-Error)

### 👍 Updated

* Updated `go-datadog-api` to version v2.24.0
* Updated `go-github` to version 26.13
* Updated `watcher` to version 1.0.7
* Updated `google-api-go-client` to version 0.10.0
* Updated `chroma` to version 0.6.7
* Updated `go-gitlab` to version 0.20.1
* Updated `trello` to version 1.4.0
* Updated `tcell` to version 1.3.0
* Updated `gopsutil` to version 2.19.9+incompatible
* Updated `yaml` to version 2.2.4

## v0.22.0

### ⚡️ Added

* [Arpansagovau](https://wtfutil.com/modules/weather_services/arpansagovau/) (arpansa.gov.au) module added, by [@jeffz](https://github.com/jeffz)
* 'calendarReadLevel' setting added to gCal module, by [@mikkeljuhl](https://github.com/mikkeljuhl)
* Todoist module now catches and displays API errors, by [@Seanstoppable](https://github.com/Seanstoppable)
* TravisCI sort orders now configurable,  by [@nyourchuck](https://github.com/nyourchuck)
* Google Analytics module now supports real-time metrics, [#581](https://github.com/wtfutil/wtf/issues/581) by [@Ameobea](https://github.com/Ameobea)
* Colors in configuration can now be defined using long-form hex, i.e.: #ff0000, by [@Seanstoppable](https://github.com/Seanstoppable)
* GitHub module pull requests are now selectable and openable via keyboard, [#547](https://github.com/wtfutil/wtf/issues/547) by [@Midnight-Conqueror](https://github.com/Midnight-Conqueror)
* [Docker](https://wtfutil.com/modules/docker/) module added, [#594](https://github.com/wtfutil/wtf/issues/594) by [@v-braun](https://github.com/v-braun)
* NewRelic module now supports displaying data from multiple apps, [#471](https://github.com/wtfutil/wtf/issues/471) by [@ChrisDBrown](https://github.com/ChrisDBrown) and [@Seanstoppable](https://github.com/Seanstoppable)
* [Digital Clock](https://wtfutil.com/modules/digitalclock/) module added, by [@Narengowda](https://github.com/Narengowda)

### 🐞 Fixed

* ScrollableWidget bounds checking error fixed, [#578](https://github.com/wtfutil/wtf/issues/578) by [@Seanstoppable](https://github.com/Seanstoppable)
* Now properly URL-decodes Jenkins branch names, [#575](https://github.com/wtfutil/wtf/issues/575) by [@lesteenman](https://github.com/lesteenman)
* Jira column sizes render properly, [#574](https://github.com/wtfutil/wtf/issues/574) by [@lesteenman](https://github.com/lesteenman)
* Todoist module updated to latest API version, by [@Seanstoppable](https://github.com/Seanstoppable)
* gCal colour highlighting working again, [#611](https://github.com/wtfutil/wtf/issues/611) by [@senorprogrammer](https://github.com/senorprogrammer)
* Per-module background and text colour settings working again, [#568](https://github.com/wtfutil/wtf/issues/568) by [@Seanstoppable](https://github.com/Seanstoppable)
* Git module no longer forces sorting of repositories, [#608](https://github.com/wtfutil/wtf/pull/608) by [@Seanstoppable](https://github.com/Seanstoppable)
* GitHub PR icons render properly without phantom characters, by [@Midnight-Conqueror](https://github.com/Midnight-Conqueror)
* GitLab configuration now takes a list of project paths, [#566](https://github.com/wtfutil/wtf/issues/566) by [@senorprogrammer](https://github.com/senorprogrammer)
* Kubernetes configuration segfault fixed, [#549](https://github.com/wtfutil/wtf/issues/549) by [@ibaum](https://github.com/ibaum)

## v0.21.0

### ⚡️ Added

* Power Soure module support added for FreeBSD, by [@hxw](https://github.com/hxw)

### 🐞 Fixed

* Power indicator displays ∞ on Linux when fully-charged and on AC power, [#534](https://github.com/wtfutil/wtf/issues/534) by [@Seanstoppable](https://github.com/Seanstoppable)
* Default background color is now the terminal background color, making transparency support possible in MacOS and Linux,  by [@Seanstoppable](https://github.com/Seanstoppable)
* `xdg-open` now used as the opener for HTTP/HTTPS by default, by [@hxw](https://github.com/hxw)
* Transmission port over-ride now working, [#565](https://github.com/wtfutil/wtf/issues/565) by [@Seanstoppable](https://github.com/Seanstoppable)
* Default config is now created on first run, [#553](https://github.com/wtfutil/wtf/issues/553) by [@senorprogrammer](https://github.com/senorprogrammer)

## v0.20.0

### ⚡️ Added

* Kubernetes module added, [#142](https://github.com/wtfutil/wtf/issues/142) by [@sudermanjr](https://github.com/sudermanjr)

### 🐞 Fixed

* Tab and Esc keys work properly in modal dialogs, [#520](https://github.com/wtfutil/wtf/issues/520) by [@Seanstoppable](https://github.com/Seanstoppable)
* `wtfutil -m` flag now works with non-enabled modules, [#529](https://github.com/wtfutil/wtf/issues/529) by [@Seanstoppable](https://github.com/Seanstoppable)
* Jenkins job filtering preserved across redraws, [#532](https://github.com/wtfutil/wtf/issues/532) by [@Seanstoppable](https://github.com/Seanstoppable)

## v0.19.1

### ⚡️ Added

* Dockerfile, by [@senorprogrammer](https://github.com/senorprogrammer)
* Add build targets for arm and arm64 architectures, by [@senorprogrammer](https://github.com/senorprogrammer)

## v0.19.0

### ☠️ Breaking Change

* HIBP module now requires an API key to operate. See [Authentication and the Have I Been Pwned API](https://www.troyhunt.com/authentication-and-the-have-i-been-pwned-api/) for more details, [#508](https://github.com/wtfutil/wtf/issues/508) by [@senorprogrammer](https://github.com/senorprogrammer)

### ⚡️ Added

* OpsGenie module now supports "region" configuration option ("us" or "eu"), by [@l13t](https://github.com/l13t)

### 🐞 Fixed

* Fixes the error message shown when an explicitly-specified custom config file cannot be found or cannot be read, by [@senorprogrammer](https://github.com/senorprogrammer)
* Rollbar module works again, [#507](https://github.com/wtfutil/wtf/issues/507) by [@Seanstoppable](https://github.com/Seanstoppable)
* The default config that gets installed on first run is much improved, [#504](https://github.com/wtfutil/wtf/issues/504) by [@senorprogrammer](https://github.com/senorprogrammer)
* Default config file is now `chmod 0600` to ensure only the owning user can read it, by [@senorprogrammer](https://github.com/senorprogrammer)

## v0.18.0

### ⚡️ Added

* Google Analytics module, by [@DylanBartels](https://github.com/DylanBartels)

### 🐞 Fixed

* Now created ~/.config if that directory is missing, [#510](https://github.com/wtfutil/wtf/issues/510) by [@senorprogrammer](https://github.com/senorprogrammer)

## v0.17.1

### 🐞 Fixed

* Fixes an issue in which the default config file was not being created on first run

## v0.17.0

### 🐞 Fixed

* FeedReader story sorting bug fixed
* NewRelic dependency vendored

## v0.16.1
## v0.16.0

### ⚡️ Added

* Config and recipe added for installing via Homebrew

## v0.15.0

### ❗️Changed

* The installed binary has been renamed from `wtf` to `wtfutil`. [Read more about it here](https://wtfutil.com/blog/2019-07-10-wtfutil-release/).

## v0.14.0

### ⚡️ Added

* CmdRunner module now supports custom titles, by [@Seanstoppable](https://github.com/Seanstoppable)
* FeedReader module added (https://wtfutil.com/modules/feedreader/), a rudimentary Rss & Atom feed reader

### 🐞 Fixed

* Cryptolive module works again, [#481](https://github.com/wtfutil/wtf/issues/481) by [@Seanstoppable](https://github.com/Seanstoppable)
* gCal module now supports setting an explicit time zone via the "timezone" config attribute, [#382](https://github.com/wtfutil/wtf/issues/382) by [@jeangovil](https://github.com/jeangovil)
* Misconfigured module positions in `config.yaml` now attempt to provide informative error messages on launch, [#482](https://github.com/wtfutil/wtf/issues/482)

## v0.13.0

### ⚡️ Added

* Transmission module addedd (https://wtfutil.com/modules/transmission/)

## v0.12.0

### ⚡️ Added

* Textfile module's text wrapping is configurable via the 'wrapText' boolean setting
* Have I Been Pwned (HIBP) module added (https://wtfutil.com/modules/hibp/)

## v0.11.0

### ⚡️ Added

* GitHub module now supports custom queries for issues and pull requests, by [@Sean-Der](https://github.com/Sean-Der)

### 🐞 Fixed

* Todoist now properly updates list items when Refresh is called
* Keyboard modal displays properly when tabbing between widgets, [#467](https://github.com/wtfutil/wtf/issues/467)

## v0.10.3

### ❗️Changed

* Invalid glog dependency removed, by [@bosr](https://github.com/bosr)

## v0.10.2

### 🐞 Fixed

* Weather module no longer crashes if there's no weather data or no internet connection
* Gitlab no longer prevents installing with missing param, [#455](https://github.com/wtfutil/wtf/issues/455)

## v0.10.1

### 🐞 Fixed

* Trello now displays multiple lists properly, [#454](https://github.com/wtfutil/wtf/issues/454)

## v0.10.0

### ⚡️ Added

* DataDog module is now scrollable and interactive, by [@Seanstoppable](https://github.com/Seanstoppable)
* Focusable hot key numbers are now assigned in a stable order, [#435](https://github.com/wtfutil/wtf/issues/435) by [@Seanstoppable](https://github.com/Seanstoppable)
* Zendesk widget now has help text, by [@Seanstoppable](https://github.com/Seanstoppable)
* Scrollable widget added to provide common keyboard-navigation list functionality, by [@Seanstoppable](https://github.com/Seanstoppable)
* Logging functionality extracted out from Log module, by [@Seanstoppable](https://github.com/Seanstoppable)
* Improved sample configs with up-to-date attributes and examples, by [@retgits](https://github.com/retgits)
* PagerDuty config now supports schedule filtering using the `scheduleIDs` config option, by [@senporprogrammer](https://github.com/senporprogrammer)

## v0.9.2

### ⚡️ Added

* Keyboard management system for modules reworked. Now has a KeyboardWidget to simplify adding keyboard commands

### Fixed

* WTF versions are now prefixed with `v` again so module systems pick up the latest versions

## 0.9.1

### ⚡️ Added

* Increased the pagination limit for GitHub to 100, by [@Seanstoppable](https://github.com/Seanstoppable)
* Support for multiple instances of the same widget added, by [@Seanstoppable](https://github.com/Seanstoppable)

## 0.9.0

* Null release

## 0.8.0

### ⚡️ Added

* Dependencies are now managed and installed using Go modules. See README.md for details, [#406](https://github.com/wtfutil/wtf/issues/406) by [@retgits](https://github.com/retgits) 

## 0.7.2

### ⚡️ Added

* NBA Scores now navigable via arrow keys, [#415](https://github.com/wtfutil/wtf/issues/415)

### 🐞 Fixed

* Multi-page sigils off-by-one error fixed, [#413](https://github.com/wtfutil/wtf/issues/413)
* Many points of potential and probable race conditions have been improved to not have race conditions. WTF should be quite a bit more stable now
* In the Twitter module, the following have been fixed:
  * Help text says "Twitter" instead of "TextFile"
  * Keyboard-command "o" opens the current Twitter handle in the browser
  * Keyboard-command "o" is documented in the help text

## 0.7.1

### 🐞 Fixed

* HackerNews row selections are visible again, [#411](https://github.com/wtfutil/wtf/issues/411)

## 0.7.0

### ⚡️ Added

* Jenkins now supports coloured balls, [#358](https://github.com/wtfutil/wtf/issues/358) by [@rudolphjacksonm](https://github.com/rudolphjacksonm)
* Jenkins now supports regular expressions, [#359](https://github.com/wtfutil/wtf/issues/359) by [@rudolphjacksonm](https://github.com/rudolphjacksonm)
* Complete refactoring of the module settings system, reducing the dependency on `config` and making it possible to configure modules by other means, by [@senporprogrammer](https://github.com/senporprogrammer)

## 0.6.0

### ⚡️ Added

* Jira widget navigable via up/down arrow keys, by [@jdenoy](https://github.com/jdenoy)
* Windows security module improved, by [@E3V3A](https://github.com/E3V3A)
* Function modules moved into the `/modules` directory, by [@Seanstoppable](https://github.com/Seanstoppable)
* NBA Score module added by [@FriedCosey](https://github.com/FriedCosey)

### 🐞 Fixed

* Now displays an error on start-up if a widget has mis-configured `position` attributes ([#389](https://github.com/wtfutil/wtf/issues/389) by @senporprogrammer)

## 0.5.0

### ⚡️ Added

* Resource Usage module added by [@nicholas-eden](https://github.com/nicholas-eden)
* Recursive repo search in Git module ([#126](https://github.com/wtfutil/wtf/issues/126) by [@anandsudhir](http://github.com/anandsudhir)) 
* HTTP/HTTPS handling in OpenFile() util function by [@jdenoy](https://github.com/jdenoy)
* Honor system http proxies when using non-default transport by [@skymeyer](https://github.com/skymeyer)
* VictorOps module added by [ImDevinC](https://github.com/ImDevinC)
* Module templates added by [retgits](https://github.com/retgits)

## 0.4.0

### ⚡️ Added

* Mecurial module added ([@mweb](https://github.com/mweb))
* Can now define numeric hotkeys in config ([@mweb](https://github.com/mweb))
* Linux firewall support added ([@TheRedSpy15](https://github.com/TheRedSpy15))
* Spotify Web module added ([@StormFireFox1](https://github.com/StormFireFox1))

### 🐞 Fixed

* Google Calendar module now displays all-day events ([#306](https://github.com/wtfutil/wtf/issues/306) by [@nicholas-eden](https://github.com/nicholas-eden))
* Google Calendar configuration much improved ([#326](https://github.com/wtfutil/wtf/issues/326) by [@dvdmssmnn](https://github.com/dvdmssmnn))

## 0.3.0

### ⚡️ Added

* Spotify module added (@sticreations)
* Clocks module now supports configurable datetime formats (@danrabinowitz)
* Twitter module now supports subscribing to multiple screen names

### 🐞 Fixed

* Textfile module now watches files for changes ([#276](https://github.com/wtfutil/wtf/issues/276) by @senporprogrammer)
* Nav shortcuts now use numbers rather than letters to allow the use of letters in widget menus
* Twitter widget no longer crashes app when closing the help modal

## 0.2.2
#### Aug 25, 2018

### ⚡️ Added

* Twitter tweets are now colourized (@senorprogrammer)
* Twitter number of tweets to fetch is now customizable via config (@senorprogrammer)
* Google Calendar: widget is now focusable (@anandsudhir)
* [DataDog module](https://wtfutil.com/modules/datadog/) added (@Seanstoppable)

### 🐞 Fixed

* Textfile syntax highlighting now included in stand-alone binary ([#261](https://github.com/wtfutil/wtf/issues/261) by @senporprogrammer)
* Config param now supports relative paths starting with `~` ([#295](https://github.com/wtfutil/wtf/issues/295) by @anandsudhir)

## 0.2.1
#### Aug 17, 2018

### ⚡️ Added

* HackerNews widget is now scrollable (@anandsudhir)

### 🐞 Fixed

* Twitter screen name now configurable in configuration file (@senorprogrammer)
* Gerrit module no longer dies if it can't connect to the server (@anandsudhir)
* Pretty Weather properly displays colours again (([#298](https://github.com/wtfutil/wtf/issues/298) by @bertl4398)
* Clocks row colour configuration fixed (([#282](https://github.com/wtfutil/wtf/issues/282) by @anandsudhir)
* Sigils no longer display when there's only one option (([#291](https://github.com/wtfutil/wtf/issues/291) by @anandsudhir)
* Jira module now responds to the "/" key (([#268](https://github.com/wtfutil/wtf/issues/268)) by @senorprogrammer)

## 0.2.0
#### Aug 3, 2018

### ⚡️ Added

* [HackerNews module](https://wtfutil.com/modules/hackernews/) added (@anandsudhir)
* [Twitter module](https://wtfutil.com/modules/twitter/) added (@Trinergy)

### 🐞 Fixed

* TravisCI module now works with Pro version thanks to @ruggi
* Sensitive credentials can now be stored in config.yml instead of ENV vars
* GCal.showDeclined config added (@baustinanki)
* Gerrit widget is now interactive, added (@anandsudhir)

---

This file attempts to adhere to the principles of [keep a changelog](https://keepachangelog.com/en/1.0.0/).
