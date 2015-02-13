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

func processOUIData(oui string) (list ouiList) {
	var oui_lines []string
	var oui_list ouiList

	oui_lines = strings.Split(oui, "\n")
	fmt.Printf("OUI Array Length: %d\n", len(oui_lines))
	re := regexp.MustCompile("(hex)")
	i := 0
	for _, e := range oui_lines {
		if re.FindString(e) != "" {
			fmt.Printf("%q\n", e)
			oui_parts := strings.Split(e, "\t")
			oui_entry := ouiEntry{macPrefix: strings.Trim(oui_parts[0], " "), manufacturer: oui_parts[3]}
			fmt.Println(oui_entry)
			oui_list[i] = oui_entry
			i++
		}
	}
	return oui_list
}

func main() {
	oui_text := getOUIFile(MA_M_OUI_URL)
	fmt.Print(oui_text)
	processOUIData(oui_text)
}
