package main

import (
	"bufio"
	"fmt"
	cliColor "github.com/gookit/color"
	"github.com/lucasb-eyer/go-colorful"
	"io"
	"os"
	"regexp"
)

type colorWord struct {
	colorful.Color
	//This space is intentionally left blank for future things like self-cleanup
}

var (
	c map[string]colorWord

	palCounter int
	pal1       []colorful.Color
)

const numColors = 50

func main() {
	reader := bufio.NewReader(os.Stdin)

	c = make(map[string]colorWord)

	pal1, _ = colorful.HappyPalette(numColors)
	palCounter = 0

	for {
		inLine, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		// TODO: add command line switch(es) for different regexp
		re := regexp.MustCompile(`\w?[\w-:.\/]*\w`)
		outLine := re.ReplaceAllStringFunc(inLine, doColorWord)

		fmt.Print(outLine)
	}

	// TODO: add command line switch to print stats
	fmt.Printf("Stats: %v words indexed.\n", len(c))
}

func doColorWord(in string) string {
	//skip short words:
	if len(in) <= 0 {
		return in
	}

	if _, found := c[in]; !found {
		c[in] = colorWord{pal1[palCounter]}
		palCounter = (palCounter + 1) % numColors
	}

	r, g, b := c[in].RGB255()

	x := cliColor.NewRGBStyle(cliColor.RGB(r, g, b))
	return x.Sprint(in)
}
