package window

import (
	"log"
	"os"

	"gioui.org/app"
	"github.com/fyne-io/mobile/app"
)

func makeWindow() {
	w := app.NewWindow()

	err := run(w)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(0)

}

func main() {
	app.Main()
}
