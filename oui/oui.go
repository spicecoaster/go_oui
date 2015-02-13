package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

const (
	MAX_OUI_ENTRIES = 100000
	MA_L_OUI_URL    = "http://standards-oui.ieee.org/oui.txt"
	MA_M_OUI_URL    = "http://standards.ieee.org/develop/regauth/oui28/mam.txt"
	MA_S_OUI_URL    = "http://standards.ieee.org/develop/regauth/iab/iab.txt"
	LOCAL_OUI_DB    = "oui_db.txt"
)

type ouiEntry struct {
	macPrefix    string
	manufacturer string
	address      [4]string
}

func getOUIFromIEEEOrg(ouiURL string) (string, error) {
	resp, err := http.Get(ouiURL)
	if err != nil {
		fmt.Print(err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
		return "", err
	} else {
		return string(body), nil
	}
	return "", nil
}

func getOUIFromLocalDB(ouiPath string) (string, error) {
	oui_data, err := ioutil.ReadFile(ouiPath)
	if err != nil {
		fmt.Print(err)
		return "", err
	}
	return string(oui_data), nil
}

func processOUIData(oui string) map[string]string {
	var oui_lines []string
	oui_list := make(map[string]string)

	oui_lines = strings.Split(oui, "\n")
	//fmt.Printf("OUI Array Length: %d\n", len(oui_lines))
	re := regexp.MustCompile("(hex)")
	for _, e := range oui_lines {
		if re.FindString(e) != "" {
			//fmt.Printf("%q\n", e)
			oui_parts := strings.Split(e, "\t")
			fmt.Printf("%s | %s | %s\n", oui_parts[0], oui_parts[1], oui_parts[2])
			//oui_entry := ouiEntry{macPrefix: strings.Trim(oui_parts[0], " "), manufacturer: oui_parts[3]}
			oui_list[strings.Trim(oui_parts[0], " ")] = oui_parts[2]
		}
	}
	return oui_list
}

func main() {
	oui_text, err := getOUIFromIEEEOrg(MA_L_OUI_URL)
	if err != nil {
		fmt.Print(err)
		oui_text, err = getOUIFromLocalDB(LOCAL_OUI_DB)
	}
	//fmt.Print(oui_text)
	m := processOUIData(oui_text)
	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}
}
