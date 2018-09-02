package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/urfave/cli"
)

const (
	EOSString = "EOS"
	Splitter  = "\t"
)

func main() {
	app := cli.NewApp()
	app.Name = "marcj"
	app.Usage = "Morphological Analysis Result Converter for Japanese"
	app.Version = "0.1.0"

	// flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "tool, t",
			Value: "sudachi",
			Usage: "Analyzer tool you used",
		},
	}

	// action
	app.Action = func(c *cli.Context) error {
		switch c.String("tool") {
		case "sudachi":
			runSudachi(bufio.NewReader(os.Stdin))
		default:
			runDefault() // raise error
		}
		return nil
	}

	app.Run(os.Args)
}

func runSudachi(reader *bufio.Reader) {
	sent := []string{}
	for {
		input, err := reader.ReadString('\n')

		if err != nil && err == io.EOF {
			break
		}

		input = strings.TrimRight(input, "\n")

		if input == EOSString {
			fmt.Println(strings.Join(sent, " "))
			sent = []string{}
		} else {
			sent = append(sent, strings.Split(input, Splitter)[0])
		}
	}

}

func runDefault() {
	fmt.Println("Not Available")
	os.Exit(1)
}
