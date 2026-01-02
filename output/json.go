package output

import (
	"encoding/json"
	"os"
	"strings"
)

type WaybarOutput struct {
	Text    string `json:"text"`
	Tooltip string `json:"tooltip"`
	Class   string `json:"class"`
}

func (o WaybarOutput) Print() error {
	return json.NewEncoder(os.Stdout).Encode(o)
}

func FormatTooltip(format, name, ip string) string {
	tooltip := strings.Replace(format, "{name}", name, -1)
	tooltip = strings.Replace(tooltip, "{ip}", ip, -1)
	return tooltip
}
