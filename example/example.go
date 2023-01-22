package main

import (
	"fmt"

	"github.com/antoniomralmeida/golibretranslate"
)

func main() {
	txt, err := golibretranslate.Translate("Bom Dia!", "auto", "en")
	if err == nil {
		fmt.Println(txt)
	}
}
