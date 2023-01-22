# golibretranslate


<h1>free api translation for golang<h1>

go get github.com/antoniomralmeida/golibretranslate


example:

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


