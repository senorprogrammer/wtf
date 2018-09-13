---
title: "Jenkins"
date: 2018-06-09T20:53:35-07:00
draft: false
weight: 130
---

<img class="screenshot" src="/imgs/modules/jenkins.png" alt="jenkins screenshot" width="320" height="68" />

Added in `v0.0.8`.

Displays jenkins status of given builds in a project or view

## Source Code

```bash
wtf/jenkins/
```

## Keyboard Commands

<span class="caption">Key:</span> `[return]` <br />
<span class="caption">Action:</span> Open the selected job in the browser.

<span class="caption">Key:</span> `j` <br />
<span class="caption">Action:</span> Select the next job in the list.

<span class="caption">Key:</span> `k` <br />
<span class="caption">Action:</span> Select the previous job in the list.

<span class="caption">Key:</span> `r` <br />
<span class="caption">Action:</span> Refresh the data.

<span class="caption">Key:</span> `↓` <br />
<span class="caption">Action:</span> Select the next job in the list.

<span class="caption">Key:</span> `↑` <br />
<span class="caption">Action:</span> Select the previous job in the list.

## Configuration

```yaml
jenkins:
  apiKey: "3276d7155dd9ee27b8b14f8743a408a9"
  enabled: true
  position:
    top: 2
    left: 3
    height: 2
    width: 3
  refreshInterval: 300
  url: "https://jenkins.domain.com/jenkins/view_url"
  user: "username"
  verifyServerCertificate: true
```

### Attributes

`apiKey` <br />
Value: Your <a href="https://wiki.jenkins.io/display/JENKINS/Remote+access+API">Jenkins API</a> key.

`enabled` <br />
Determines whether or not this module is executed and if its data displayed onscreen. <br />
Values: `true`, `false`.

`position` <br />
Defines where in the grid this module's widget will be displayed.

`refreshInterval` <br />
How often, in seconds, this module will update its data. <br />
Values: A positive integer, `0..n`.

`user` <br />
Your Jenkins username. <br />

`url` <br />
The url to your Jenkins project or view. <br />
Values: A valid URI.

`verifyServerCertificate` <br />
_Optional_ <br />
Determines whether or not the server's certificate chain and host name are verified. <br />
Values: `true`, `false`.
