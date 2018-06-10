package security

import (
	"fmt"
	"strings"

	"github.com/andrewzolotukhin/wtf/wtf"
	"github.com/olebedev/config"
)

// Config is a pointer to the global config object
var Config *config.Config

type Widget struct {
	wtf.TextWidget
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" Security ", "security", false),
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	data := NewSecurityData()
	data.Fetch()

	widget.UpdateRefreshedAt()

	widget.View.SetText(fmt.Sprintf("%s", widget.contentFrom(data)))
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) contentFrom(data *SecurityData) string {
	str := " [red]WiFi[white]\n"
	str = str + fmt.Sprintf(" %8s: %s\n", "Network", data.WifiName)
	str = str + fmt.Sprintf(" %8s: %s\n", "Crypto", data.WifiEncryption)
	str = str + "\n"
	str = str + " [red]Firewall[white]        [red]DNS[white]\n"
	str = str + fmt.Sprintf(" %8s: [%s]%-3s[white]   %-16s\n", "Enabled", widget.labelColor(data.FirewallEnabled), data.FirewallEnabled, data.DnsAt(0))
	str = str + fmt.Sprintf(" %8s: [%s]%-3s[white]   %-16s\n", "Stealth", widget.labelColor(data.FirewallStealth), data.FirewallStealth, data.DnsAt(1))
	str = str + "\n"
	str = str + " [red]Users[white]\n"
	str = str + fmt.Sprintf(" %s", strings.Join(data.LoggedInUsers, ", "))

	return str
}

func (widget *Widget) labelColor(label string) string {
	switch label {
	case "on":
		return "green"
	case "off":
		return "red"
	default:
		return "white"
	}
}
