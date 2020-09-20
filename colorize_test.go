package utilities

import (
	"fmt"
	"testing"
)

func TestColor(t *testing.T) {
	fmt.Printf("%s\n%s\n%s\n", Success("I'm Success"), Red("I'm Red"), Yellow("I'm Yellow"))
	fmt.Printf("%s\n%s\n%s\n", Info("I'm Info"), Warn("I'm Warn"), Fatal("I'm Fatal"))
	fmt.Printf("%s\n%s\n%s\n", Normal("I'm Normal"), Purple("I'm Purple"), Magenta("I'm Magenta"))
	fmt.Printf("%s\n%s\n%s\n", White("I'm White"), LightWhite("I'm LightWhite"), LightRed("I'm LightRed"))
	fmt.Printf("%s\n%s\n", LightGreen("I'm LightGreen"), LightYellow("I'm LightYellow"))
	fmt.Printf("%s\n%s\n", LightBlue("I'm LightBlue"), LightMagenta("I'm LightMagenta"))
	fmt.Printf("%s\n%s\n%s\n", LightTeal("I'm LightTeal"), LightWhite("I'm LightWhite"), Normal("I'm _Normal"))
}
