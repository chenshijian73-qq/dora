package pkg

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"os"
)

var colors = []color.Attribute{
	color.FgRed,
	color.FgGreen,
	color.FgYellow,
	color.FgBlue,
	color.FgMagenta,
	color.FgCyan,
}

type ColorLine struct {
	Prefix string
	Value  string
}

func Render(output string, prefix string, cmdError error) error {
	c := colors[rand.Intn(len(colors))]

	color.Set(c)
	defer color.Unset()
	_, err := fmt.Fprintf(os.Stdout, "%v:\n", prefix)
	if err != nil {
		return err
	}
	color.Set(color.FgHiWhite)
	_, err = fmt.Fprintf(os.Stdout, "%v", output)
	if err != nil {
		return err
	}
	if cmdError != nil {
		color.Set(color.FgRed)
		_, err = fmt.Fprintf(os.Stdout, "%v\n", cmdError)
	}
	return err
}
