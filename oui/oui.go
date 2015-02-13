package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

const (
	MAX_OUI_ENTRIES = 1000
	MA_L_OUI_URL    = "http://standards-oui.ieee.org/oui.txt"
	MA_M_OUI_URL    = "http://standards.ieee.org/develop/regauth/oui28/mam.txt"
	MA_S_OUI_URL    = "http://standards.ieee.org/develop/regauth/iab/iab.txt"
)

type ouiEntry struct {
	macPrefix    string
	hexString    string
	manufacturer string
	address      [4]string
}

type ouiList [MAX_OUI_ENTRIES]ouiEntry

func GetMAMOUI() {
}

func getOUIFile(ouiURL string) string {
	resp, err := http.Get(ouiURL)
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
	} else {
		return string(body)
	}
	return ""
}

func main() {
	oui_text := getOUIFile(MA_M_OUI_URL)
	fmt.Print(oui_text)
	var oui_lines []string
	oui_lines = strings.Split(oui_text, "\n")
	fmt.Printf("OUI Array Length: %d\n", len(oui_lines))
	re := regexp.MustCompile("(hex)")
	for _, e := range oui_lines {
		if re.FindString(e) != "" {
			fmt.Printf("%s\n", e)
		}
	}
}
