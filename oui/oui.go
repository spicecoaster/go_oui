package oui

import (
	"fmt"
)

const (
	OUI_URL = "http://standards-oui.ieee.org/oui.txt"
)

type ouiEntry struct {
	macPrefix    string
	hexString    string
	manufacturer string
}

type ouiList struct {
	ouiList []ouiEntry
}
